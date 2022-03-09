package middleware

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/domain/service"
	"kitchenmaniaapi/interfaces/response"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize(findUserByEmail func(string) (*entity.User, error), tokenInBlacklist func(*string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("JWT_SECRET")
		accToken := service.GetTokenFromHeader(c)
		accessToken, accessClaims, err := service.AuthorizeToken(&accToken, &secret)
		if err != nil {
			log.Printf("authorize access token error: %s\n", err.Error())
			respondAndAbort(c, "", http.StatusUnauthorized, nil, "unauthorized")
			return
		}

		if tokenInBlacklist(&accessToken.Raw) || isTokenExpired(accessClaims) {
			rt := &struct {
				RefreshToken string `json:"refresh_token,omitempty" binding:"required"`
			}{}

			if err := c.ShouldBindJSON(rt); err != nil {
				log.Printf("no refresh token in request body: %v\n", err)
				respondAndAbort(c, "", http.StatusBadRequest, nil, "unauthorized")
				return
			}

			if tokenInBlacklist(&rt.RefreshToken) {
				log.Printf("refresh token is blacklisted: %v\n", err)
				respondAndAbort(c, "", http.StatusUnauthorized, nil, "refresh token is invalid")
				return
			}

			_, rtClaims, err := service.AuthorizeToken(&rt.RefreshToken, &secret)
			if err != nil {
				log.Printf("authorize refresh token error: %v\n", err)
				respondAndAbort(c, "", http.StatusUnauthorized, nil, "refresh token is invalid")
				return
			}

			if isTokenExpired(rtClaims) {
				log.Printf("refresh token is expired")
				respondAndAbort(c, "", http.StatusUnauthorized, nil, "refresh token is invalid")
				return
			}

			if sub, ok := rtClaims["sub"].(float64); ok && sub != 1 {
				log.Printf("invalid refresh token, the sub claim isn't correct")
				respondAndAbort(c, "", http.StatusUnauthorized, nil, "refresh token is invalid")
				return
			}

			//generate a new access token, and rest its exp time
			accessClaims["exp"] = time.Now().Add(service.AccessTokenValidity).Unix()
			newAccessToken, err := service.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
			if err != nil {
				log.Printf("can't generate new access token: %v\n", err)
				respondAndAbort(c, "", http.StatusUnauthorized, nil, "can't generate new access token")
				return
			}
			respondAndAbort(c, "new access token generated", http.StatusOK, gin.H{"access_token": *newAccessToken}, "access token is invalid")
			return
		}
		var user *entity.User
		if email, ok := accessClaims["user_email"].(string); ok {
			if user, err = findUserByEmail(email); err != nil {
				log.Printf("find user by email error: %v\n", err)
				respondAndAbort(c, "", http.StatusNotFound, nil, "user not found")
				return
			}
		} else {
			log.Printf("user email is not string\n")
			respondAndAbort(c, "", http.StatusInternalServerError, nil, "internal server error")
			return
		}
		c.Set("user", user)
		c.Set("access_token", accessToken.Raw)
		// calling next handler
		c.Next()
	}
}

func respondAndAbort(c *gin.Context, message string, status int, data interface{}, errs string) {
	response.JSON(c, message, status, data, errs)
	c.Abort()
}

func isTokenExpired(claims jwt.MapClaims) bool {
	if exp, ok := claims["exp"].(float64); ok {
		return float64(time.Now().Unix()) > exp
	}
	return true
}
