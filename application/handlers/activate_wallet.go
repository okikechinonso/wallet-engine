package handlers

import (
	"log"
	"net/http"
	"wallet-engine/infrastructure/response"

	"github.com/gin-gonic/gin"
)

// ActivateWallet activate or deactivate wallet
func (a *App) ActivateWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("user")
		if !ok {
			response.JSON(c, "unauthorized user", http.StatusUnauthorized, nil)
			return
		}
		walletAddress := c.Query("wallet_address")
		log.Println("here")

		//gets the wallet of a user by address
		wallet, err := a.DB.FindWallet(walletAddress)
		if err != nil {
			log.Println("here")
			response.JSON(c, "wallet doesn't exist", http.StatusNotFound, nil)
			return
		}

		//check if the wallet is active
		if !wallet.Active {
			wallet.Active = true

			//activate wallet if wallet is not active
			err = a.DB.ActiveWallet(*wallet)
			log.Println("here")
			if err != nil {
				response.JSON(c, "unable to activate wallet", http.StatusInternalServerError, nil)
				return
			}
			response.JSON(c, "account dectivated", http.StatusCreated, wallet)
			return
		}

		//deactivate wallet if wallet is active
		wallet.Active = false
		err = a.DB.ActiveWallet(*wallet)
		if err != nil {
			response.JSON(c, "unable to deactivate wallet", http.StatusInternalServerError, nil)
			return
		}
		response.JSON(c, "account dectivated", http.StatusCreated, wallet)
	}
}
