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
		user := entity.User{}
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
			err = a.DB.NewUser(user)
			if err != nil {
				response.JSON(c, err.Error(), http.StatusInternalServerError, nil)
				return
			}
			if _, err = a.DB.FindWallet(user.Phone); err != nil {
				wallet, err := a.DB.NewWallet(user.Phone, user.ID)
				if err != nil {
					response.JSON(c, err.Error(), http.StatusInternalServerError, nil)
					return
				}
				user.UserWallet = *wallet
			}
			log.Println(user,"here")
			response.JSON(c, "wallet created successfully", http.StatusOK, nil)
			return
		}
		response.JSON(c, err.Error(), http.StatusBadRequest, nil)
	}
}
