package server

import (
	"fmt"
	"kitchenmaniaapi/application/handlers"
	"kitchenmaniaapi/application/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	App handlers.App
}

func (s *Server) defineRoute(router *gin.Engine) {
	apirouter := router.Group("/api/v1")
	apirouter.POST("/signup", s.App.SignUp())
	apirouter.POST("/login", s.App.Login())

	authorized := apirouter.Group("/")
	authorized.Use(middleware.Authorize(s.App.DB.FindUserByEmail, s.App.DB.TokenInBlacklist))
	authorized.POST("/create", s.App.CreateBlog())
	authorized.POST("/updateblog",s.App.UpdatePost())
	authorized.GET("/posts",s.App.GetAllPost())
	authorized.DELETE("/delete",s.App.DeletePost())
}

func (s *Server) setupRoute() *gin.Engine {
	r := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	// setup cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	s.defineRoute(r)
	return r
}

func (s *Server) Start() {
	r := s.setupRoute()
	port := os.Getenv("PORT")
	
	server := &http.Server{
		Addr:    ":"+port,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("unable to start serve", err)
	}
}


