package handlers

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/interfaces/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		userI, exists := c.Get("user")
		if !exists {
			response.JSON(c, "", http.StatusInternalServerError, nil, "no user found")
			return
		}
		user := userI.(*entity.User)
		_, err := app.DB.FindUserByEmail(user.Email)
		if err != nil {
			response.JSON(c, "", http.StatusUnauthorized, nil, "user not authorized")
			return
		}
		// log.Println(c.PostForm("title"),c)

		blog := entity.Blog{
			Title: c.PostForm("title"),
			Body:  c.PostForm("body"),
			UserID: user.ID,
			Author: user.FirstName+" "+user.LastName,
		}

		err = app.DB.UpdatePost(blog)
		if err != nil {
			response.JSON(c, "", http.StatusInternalServerError, nil, err.Error())
			return
		}

	}
}
