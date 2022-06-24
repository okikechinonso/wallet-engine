package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *App) GetMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		str := ctx.Query("page")

		page, err := strconv.Atoi(str)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Enter a valid page number")
			return
		}
		movies, err := a.DB.GetMovies(page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Unable to fetch movies")
		}
		ctx.JSON(http.StatusAccepted, movies)
		log.Println(page)
	}
}
