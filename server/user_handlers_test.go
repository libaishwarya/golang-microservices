package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"myapp/store"
	"myapp/store/inmemory"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userStore := inmemory.NewInMemoryUserStore()
	r := gin.Default()
	handler := NewUserHandler(userStore)
	AttachUserRoutes(handler, r)

	// Prepare request
	reqBody, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "password",
	})
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userStore := inmemory.NewInMemoryUserStore()
	user := store.User{
		Email:    "test@example.com",
		Password: "password",
	}
	userStore.CreateUser(&user)

	r := gin.Default()
	handler := NewUserHandler(userStore)
	AttachUserRoutes(handler, r)

	// Prepare request
	reqBody, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "password",
	})
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)
}
