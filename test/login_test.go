package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wallet-engine/application/handlers"
	"wallet-engine/application/server"
	"wallet-engine/domain/entity"
	"wallet-engine/infrastructure/helpers"
	"wallet-engine/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedDB := mock.NewMockPersistor(ctrl)

	s := &server.Server{
		App: handlers.App{mockedDB},
	}
	route := gin.Default()
	s.DefineRoute(route)

	t.Run("Test_For_Login_Request", func(t *testing.T) {
		loginRequest := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Email:    "jdoe@gmail.com",
			Password: "",
		}
		jsonFile, err := json.Marshal(loginRequest)
		if err != nil {
			t.Error("Failed to marshal file")
		}
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(string(jsonFile)))
		req.Header.Set("Content-Type", "application/json")
		route.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Test_FindUserByEmail", func(t *testing.T) {
		hashedP, _ := helpers.GenerateHashPassword("password")
		user := &entity.User{Email: "jdoe@gmail.com", HashedPassword: string(hashedP)}
		mockedDB.EXPECT().FindUserByEmail(user.Email).Return(user, nil)
		loginRequest := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Email:    "jdoe@gmail.com",
			Password: "password",
		}
		jsonFile, err := json.Marshal(loginRequest)
		if err != nil {
			t.Error("Failed to marshal file")
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", strings.NewReader(string(jsonFile)))

		route.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), loginRequest.Email)
	})
}
