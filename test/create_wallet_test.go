package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wallet-engine/application/handlers"
	"wallet-engine/application/server"
	"wallet-engine/domain/entity"
	"wallet-engine/infrastructure/helpers"
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

	t.Run("login_successful_login", func(t *testing.T) {
		user := &entity.User{
			FirstName: "john",
			LastName:  "Dee",
			Password:  "password123",
			Email:     "jdoe@gmail.com",
			Phone:     "07012345678",
		}
		wallet := &entity.Wallet{
			WalletAddress: user.Phone,
			UserID:        user.ID,
			Balance:       0,
			Active:        false,
		}
		hash, _ := helpers.GenerateHashPassword("password123")
		user.HashedPassword = string(hash)
		err := errors.New("no user found")

		mockedDB.EXPECT().FindUserByEmail(user.Email).Return(user, err)
		mockedDB.EXPECT().FindWallet(user.Phone).Return(wallet, err)
		mockedDB.EXPECT().NewUser(gomock.Any()).Return(user, nil)
		mockedDB.EXPECT().NewWallet(user.Phone, user.ID).Return(wallet, nil)

		body, err := json.Marshal(user)
		if err != nil {
			t.Fail()
			return
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/create", strings.NewReader(string(body)))
		route.ServeHTTP(w, req)

		bodyString := w.Body.String()
		log.Println(w.Body.String())
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, bodyString, fmt.Sprintf("wallet created successfully"))

	})
	t.Run("login_unsuccessful_login", func(t *testing.T) {
		user := &entity.User{
			FirstName: "john",
			LastName:  "Dee",
			Password:  "password123",
			Email:     "jdoe@gmail.com",
			Phone:     "07012345678",
		}

		hash, _ := helpers.GenerateHashPassword("password123")
		user.HashedPassword = string(hash)
		err := errors.New("no user found")

		mockedDB.EXPECT().FindUserByEmail(user.Email).Return(user, nil)

		body, err := json.Marshal(user)
		if err != nil {
			t.Fail()
			return
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/create", strings.NewReader(string(body)))
		route.ServeHTTP(w, req)

		bodyString := w.Body.String()
		log.Println(w.Body.String())
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, bodyString, fmt.Sprintf("user already exist"))

	})
}
