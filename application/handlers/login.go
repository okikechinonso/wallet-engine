package handlers

import (
	"log"
	"net/http"
	"os"
	"wallet-engine/infrastructure/helpers"
	"wallet-engine/infrastructure/response"
	"wallet-engine/infrastructure/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (a *App) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginDetail := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{}
		err := helpers.Decode(c, &loginDetail)
		if err != nil {
			log.Println(err)
			response.JSON(c, "all fields are required", http.StatusBadRequest, loginDetail)
			return
		}

		user, err := a.DB.FindUserByEmail(loginDetail.Email)
		if err != nil {
			log.Println(err)
			response.JSON(c, "invalid email or password", http.StatusBadRequest, nil)
			return
		}
		log.Println(user)
		err = helpers.CompareHashAndPassword([]byte(loginDetail.Password), user.HashedPassword)
		if err != nil {
			response.JSON(c, "invalid password or email", http.StatusUnauthorized, nil)
			return
		}
		secret := os.Getenv("JWT_SECRET")
		accessClaims, refresheClaims := token.GenerateClaims(user.Email)

		accessToken, err := token.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
		if err != nil {
			response.JSON(c, "Unable to generate accesstoken", http.StatusUnauthorized, nil,)
			return
		}
		refreshToken, err := token.GenerateToken(jwt.SigningMethodHS256, refresheClaims, &secret)
		if err != nil {
			response.JSON(c, "Unable to generate refreshtoken", http.StatusUnauthorized, nil )
			return
		}
		user.HashedPassword = ""
		response.JSON(c, "login successful", http.StatusOK, gin.H{
			"user":          user,
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}
