package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *App) GetComment() gin.HandlerFunc {
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

func (a *App) GetCommentByEmail() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		mp := make(map[string]string)
		
		err := ctx.ShouldBindJSON(&mp)
		log.Print(mp)
		if err != nil{
			
			ctx.JSON(http.StatusBadRequest, "Enter a valid fields")
		}

		if _,ok := mp["id"]; !ok{
			log.Print(mp)
			ctx.JSON(http.StatusInternalServerError, "Must container email field")
		}
		movie, err := a.DB.GetCommentByEmail(mp["id"])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		ctx.JSON(http.StatusAccepted, movie)
	}
}

func (a *App) GetByDateRange() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		mp := make(map[string]time.Time)
		str := ctx.Query("page")

		page, err := strconv.Atoi(str)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Enter a valid page number")
			return
		}
		err = ctx.ShouldBindJSON(&mp)
		log.Print(mp)
		if err != nil{
			
			ctx.JSON(http.StatusBadRequest, "Enter a valid fields")
		}

		if _,ok := mp["id"]; !ok{
			log.Print(mp)
			ctx.JSON(http.StatusInternalServerError, "Must container email field")
		}
		movie, err := a.DB.GetCommentByDate(mp["to"],mp["from"],page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		ctx.JSON(http.StatusAccepted, movie)
	}
}
