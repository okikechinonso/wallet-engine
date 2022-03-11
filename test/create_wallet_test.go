package test

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"testing"
	"wallet-engine/application/handlers"
	"wallet-engine/application/server"
	"wallet-engine/mock"
)

func TestCreateWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedDB := mock.NewMockPersistor(ctrl)

	s := &server.Server{
		App: handlers.App{mockedDB},
	}
	route := gin.Default()
	s.DefineRoute(route)

}
