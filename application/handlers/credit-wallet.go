package handlers

import (
	"net/http"
	"wallet-engine/domain/entity"
	"wallet-engine/domain/service"
	"wallet-engine/infrastructure/helpers"
	"wallet-engine/infrastructure/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreditWallet credits the wallet with certain
func (a *App) CreditWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		detail := &struct {
			Amount int64 `json:"amount"`
		}{}
		userI, ok := c.Get("user")
		if !ok {
			response.JSON(c, "unauthorized user", http.StatusUnauthorized, nil)
			return
		}
		err := helpers.Decode(c, &detail)
		if err != nil {
			response.JSON(c, "unable to get amount", http.StatusInternalServerError, nil)
			return
		}

		//converts the amount to Kobo
		detail.Amount = detail.Amount * 100
		if detail.Amount == 0 {
			response.JSON(c, "insert amount", http.StatusBadRequest, nil)
			return
		}
		user := userI.(*entity.User)

		// returns the wallet of particular user
		wallet, err := a.DB.FindWallet(user.Phone)

		//check if account is active
		if !wallet.Active {
			response.JSON(c, "wallet not active", http.StatusInternalServerError, nil)
			return
		}
		if err != nil {
			response.JSON(c, "no wallet found", http.StatusInternalServerError, nil)
			return
		}
		id := uuid.New().String()
		transaction := service.NewTransaction(id, user.Phone, "", user.ID, "credit", detail.Amount)
		wallet.Balance += detail.Amount

		//Create a new transaction for each deposit
		err = a.DB.CreateTransaction(transaction)
		if err != nil {
			response.JSON(c, "unable to create transaction", http.StatusInternalServerError, nil)
			return
		}
		//update the previous of the wallet
		err = a.DB.UpdateWallet(*wallet)
		if err != nil {
			response.JSON(c, "unable to perform transaction", http.StatusInternalServerError, nil)
			return
		}

		response.JSON(c, "Desposit successfull", http.StatusOK, gin.H{
			"transaction": transaction,
			"balance":     wallet.Balance / 100,
		})
	}

}
