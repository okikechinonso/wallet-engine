package handlers

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/interfaces/helpers"
	"kitchenmaniaapi/interfaces/response"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (a *App) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &entity.User{}

		err := helpers.Decode(c, user)
		if err != nil {
			// fmt.Errorf("enter all required field")
			response.JSON(c, "", http.StatusBadRequest, nil, "enter all field")
			log.Printf("Error: %v", err.Error())
			return
		}

		hashedPassword, err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			response.JSON(c, "", http.StatusInternalServerError, nil, "internal server error")
			return
		}
		user.Email = strings.ToLower(user.Email)
		user.HashedPassword = string(hashedPassword)
		_, err = a.DB.FindUserByEmail(user.Email)
		if err != nil {
			err = a.DB.NewUser(user)
			if err != nil {
				response.JSON(c, "", http.StatusInternalServerError, nil, "internal server error")
				return
			}
			response.JSON(c, "user created successful", http.StatusOK, nil, "")
			return
		}
		response.JSON(c, "", http.StatusBadRequest, nil, "user email already exist")
	}
}

func (a *App) SignupWithGoogle() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
