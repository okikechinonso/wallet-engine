package handlers

import (
	"log"
	"net/http"
	"strings"
	"wallet-engine/domain/entity"
	"wallet-engine/infrastructure/helpers"
	"wallet-engine/infrastructure/response"

	"github.com/gin-gonic/gin"
)

func (a *App) CreateWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &entity.User{}
		wallet := &entity.Wallet{}
		err := helpers.Decode(c, &user)
		if err != nil {
			response.JSON(c, err.Error(), http.StatusBadRequest, nil)
			return
		}
		hashedPassword, err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			response.JSON(c, "internal server error", http.StatusInternalServerError, nil)
			return
		}

		user.Email = strings.ToLower(user.Email)
		user.HashedPassword = string(hashedPassword)
		_, err = a.DB.FindUserByEmail(user.Email)
		if err != nil {
			data := make(map[string]interface{})
			_, err = a.DB.FindWallet(user.Phone)
			if err == nil {
				response.JSON(c, "wallet already exist", http.StatusInternalServerError, nil)
				return
			}
			user, err = a.DB.NewUser(*user)
			if err != nil {
				response.JSON(c, "Error creating user", http.StatusInternalServerError, nil)
				return
			}
			data["user"] = user

			wallet, err = a.DB.NewWallet(user.Phone, user.ID)
			
			if err != nil {
				response.JSON(c, "Error creating wallet", http.StatusInternalServerError, nil)
				return
			}
			data["wallet"] = wallet
			log.Println(user, "here")
			response.JSON(c, "wallet created successfully", http.StatusOK, data)
			return
		}
		response.JSON(c, "user already exist", http.StatusBadRequest, nil)
	}
}
