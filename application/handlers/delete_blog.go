package handlers

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/interfaces/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		userI, exist := c.Get("user")
		if !exist {
			response.JSON(c, "can't get user from context", http.StatusInternalServerError, nil, "no user found")
			return
		}
		user := userI.(*entity.User)
		blogID := c.Query("blog_id")
		log.Println(blogID)
		log.Println(user.ID)
		err := app.DB.DeletePost(blogID, user.ID)
		if err != nil {
			log.Println(err)
			response.JSON(c, "can't get user from context", http.StatusInternalServerError, nil, "no user found")
			return
		}
		response.JSON(c, "Successfully deleted", http.StatusAccepted, nil, "")
	}
}
