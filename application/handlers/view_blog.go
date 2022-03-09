package handlers

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/interfaces/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) GetAllPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		userI, exists := c.Get("user")
		if !exists {
			response.JSON(c, "", http.StatusInternalServerError, nil, "no user found")
			return
		}
		user := userI.(*entity.User)
		posts, err := app.DB.GetAllPosts(*user)
		if err != nil {
			response.JSON(c, "", http.StatusInternalServerError, nil, "no user found")
			return
		}
		response.JSON(c, "successful", http.StatusAccepted, posts, "")

	}
}
