package handlers

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/domain/service"
	"kitchenmaniaapi/interfaces/helpers"
	"kitchenmaniaapi/interfaces/response"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func (a *App) CreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		userI, exist := c.Get("user")
		if !exist {
			response.JSON(c, "can't get user from context", http.StatusInternalServerError, nil, "no user found")
			return
		}
		user, ok := userI.(*entity.User)
		if !ok {
			response.JSON(c, "", http.StatusInternalServerError, nil, "Unable to decode user")
			return
		}
		post := &entity.Blog{}

		form, err := c.MultipartForm()
		if err != nil {
			response.JSON(c, "", http.StatusInternalServerError, nil, "No file multipart form found")
			return
		}
		images := form.File["images"]
		if len(images) != 0{
			for _, f := range images {
				file, err := f.Open()
				if err != nil {
					response.JSON(c, "", http.StatusInternalServerError, nil, "No file multipart form found")
					return
				}
	
				if _, ok = helpers.CheckSupportedFile(strings.ToLower(f.Filename)); ok {
					log.Println(filepath.Ext(f.Filename))
					response.JSON(c, "", http.StatusInternalServerError, nil, "invalid file extension")
					return
				}
	
				session, tempFileUrl, err := service.PreAWS(filepath.Ext(f.Filename), "blogphoto")
				if err != nil {
					response.JSON(c, "", http.StatusInternalServerError, nil, "unable to create AWS session"+err.Error())
					return
				}
				url, err := service.UploadFileToS3(session, file, tempFileUrl, f.Size)
				if err != nil {
					response.JSON(c, "", http.StatusInternalServerError, nil, "unable to upload to File"+err.Error())
					return
				}
				imgs := entity.Image{ImageUrl: url}
				post.Images = append(post.Images, imgs)
			}
		}
		

		post.UserID = user.ID
		post.ID = bson.NewObjectId().Hex()
		post.Author = strings.Trim(user.FirstName + " " + user.LastName," ")
		post.Title = c.PostForm("title")
		post.Body = c.PostForm("body")
		log.Println(c)

		err = a.DB.CreatePost(*post)
		if err != nil {
			response.JSON(c, "", http.StatusInternalServerError, nil, err.Error())
			return
		}
		response.JSON(c, "blog created successfully", http.StatusOK, post, "")
	}
}
