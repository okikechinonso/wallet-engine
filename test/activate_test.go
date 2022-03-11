package test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"wallet-engine/application/handlers"
	"wallet-engine/application/server"
	"wallet-engine/domain/entity"
	"wallet-engine/infrastructure/token"
	"wallet-engine/mock"
)

func TestActivateWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedDB := mock.NewMockPersistor(ctrl)

	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}

	s := &server.Server{
		App: handlers.App{mockedDB},
	}
	route := gin.Default()
	s.DefineRoute(route)

	user := &entity.User{
		FirstName: "john",
		LastName:  "Dee",
		Password:  "password123",
		Email:     "jdoe@gmail.com",
		Phone:     "07012345678",
	}
	detail := &struct {
		Amount int64 `json:"amount"`
	}{
		Amount: 10,
	}

	accessClaims, _ := token.GenerateClaims("jdoe@gmail.com")
	secret := os.Getenv("JWT_SECRET")
	accToken, err := token.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	if err != nil {
		t.Fail()
		return
	}

	body, err := json.Marshal(detail)
	if err != nil {
		t.Fail()
		return
	}
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/activate", strings.NewReader(string(body)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))

	t.Run("Activate", func(t *testing.T) {
		wallet := &entity.Wallet{
			WalletAddress: user.Phone,
			Active:        true,
		}

		mockedDB.EXPECT().FindUserByEmail(user.Email).Return(user, nil)
		mockedDB.EXPECT().TokenInBlacklist(accToken).Return(false)
		mockedDB.EXPECT().FindWallet(gomock.Any()).Return(wallet, nil)
		mockedDB.EXPECT().ActiveWallet(gomock.Any()).Return(nil)

		resp := httptest.NewRecorder()

		route.ServeHTTP(resp, req)
		log.Println(resp.Body.String())
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.Contains(t, resp.Body.String(), "false")
	})
	t.Run("Activate", func(t *testing.T) {
		wallet := &entity.Wallet{
			WalletAddress: user.Phone,
			Active:        false,
		}

		mockedDB.EXPECT().FindUserByEmail(user.Email).Return(user, nil)
		mockedDB.EXPECT().TokenInBlacklist(accToken).Return(false)
		mockedDB.EXPECT().FindWallet(gomock.Any()).Return(wallet, nil)
		mockedDB.EXPECT().ActiveWallet(gomock.Any()).Return(nil)

		resp := httptest.NewRecorder()

		route.ServeHTTP(resp, req)
		log.Println(resp.Body.String())
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.Contains(t, resp.Body.String(), "true")
	})

}
