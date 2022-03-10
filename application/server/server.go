package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"wallet-engine/application/handlers"
	"wallet-engine/application/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	App handlers.App
}

func (s *Server) defineRoute(router *gin.Engine) {
	apirouter := router.Group("/api/v1")
	apirouter.POST("/create", s.App.CreateWallet())
	apirouter.POST("/login",s.App.Login())

	authorized := apirouter.Group("/")
	authorized.Use(middleware.Authorize(s.App.DB.FindUserByEmail, s.App.DB.TokenInBlacklist))
	authorized.POST("/credit",s.App.CreditWallet())
	authorized.POST("/debit",s.App.DebitWallet())
	authorized.PUT("/activate",s.App.ActiveWallet())
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
		MaxAge: 12 * time.Hour,
	}))

	s.defineRoute(r)
	return r
}

func (s *Server) Start() {
	r := s.setupRoute()
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("unable to start serve", err)
	}
}
