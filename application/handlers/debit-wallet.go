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

func (a *App) DebitWallet() gin.HandlerFunc {
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
		detail.Amount = detail.Amount * 100
		if detail.Amount == 0 {
			response.JSON(c, "insert amount", http.StatusBadRequest, nil)
			return
		}
		user := userI.(*entity.User)
		wallet, err := a.DB.FindWallet(user.Phone)

		if !wallet.Active {
			response.JSON(c, "wallet not active", http.StatusInternalServerError, nil)
			return
		}

		if err != nil {
			response.JSON(c, "no wallet found", http.StatusInternalServerError, nil)
			return
		}
		if wallet.Balance == 0 || wallet.Balance < detail.Amount {
			response.JSON(c, "insufficient fund",http.StatusBadRequest,nil)
			return
		}
		id := uuid.New().String()
		transaction := service.NewTransaction(id, user.Phone, "", user.ID, "credit", detail.Amount)
		wallet.Balance -= detail.Amount

		err = a.DB.CreateTransaction(transaction)
		if err != nil {
			response.JSON(c, "unable to create transaction", http.StatusInternalServerError, nil)
			return
		}
		err = a.DB.UpdateWallet(*wallet)
		if err != nil {
			response.JSON(c, "unable to perform transaction", http.StatusInternalServerError, nil)
			return
		}
		transaction.Amount = transaction.Amount/100
		response.JSON(c, "Desposit successfull", http.StatusOK, gin.H{
			"transaction": transaction,
			"balance":     wallet.Balance / 100,
		})
	}
}
