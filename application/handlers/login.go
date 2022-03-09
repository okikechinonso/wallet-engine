package handlers

import (
	"kitchenmaniaapi/domain/service"
	"kitchenmaniaapi/interfaces/helpers"
	"kitchenmaniaapi/interfaces/response"
	"log"
	"net/http"
	"os"

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
			response.JSON(c, "", http.StatusBadRequest, loginDetail, err.Error())
			return
		}

		user, err := a.DB.FindUserByEmail(loginDetail.Email)
		if err != nil {
			log.Println(err)
			response.JSON(c, "invalid email or password", http.StatusBadRequest, nil, err.Error())
			return
		}
		log.Println(user)
		err = helpers.CompareHashAndPassword([]byte(loginDetail.Password), user.HashedPassword)
		if err != nil {
			response.JSON(c, "invalid password or email", http.StatusUnauthorized, nil, err.Error())
			return
		}
		secret := os.Getenv("JWT_SECRET")
		accessClaims, refresheClaims := service.GenerateClaims(user.Email)

		accessToken, err := service.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
		if err != nil {
			response.JSON(c, "Unable to generate accesstoken", http.StatusUnauthorized, nil, err.Error())
			return
		}
		refreshToken, err := service.GenerateToken(jwt.SigningMethodHS256, refresheClaims, &secret)
		if err != nil {
			response.JSON(c, "Unable to generate refreshtoken", http.StatusUnauthorized, nil, err.Error())
			return
		}
		user.HashedPassword = ""
		response.JSON(c, "login successful", http.StatusOK, gin.H{
			"user":          user,
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		}, "")
	}
}
