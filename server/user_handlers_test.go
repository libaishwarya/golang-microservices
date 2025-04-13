package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/libaishwarya/myapp/store"
	"github.com/libaishwarya/myapp/store/inmemory"
	thirdparty "github.com/libaishwarya/myapp/userservice/third_party"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userStore := inmemory.NewInMemoryUserStore()
	r := gin.Default()
	handler := NewUserHandler(userStore, nil)
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

func TestRegister_Validate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userStore := inmemory.NewInMemoryUserStore()
	r := gin.Default()
	handler := NewUserHandler(userStore, nil)
	AttachUserRoutes(handler, r)

	tests := []struct {
		body           map[string]string
		expectedStatus int
	}{
		{map[string]string{"email": "test@example.com", "password": "password"}, http.StatusOK},
		{map[string]string{"email": "invalid-email", "password": "password"}, http.StatusBadRequest},
		{map[string]string{"email": "test@example.com", "password": "short"}, http.StatusBadRequest},
	}

	for _, test := range tests {
		reqBody, _ := json.Marshal(test.body)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, test.expectedStatus, w.Code)
	}
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
	handler := NewUserHandler(userStore, nil)
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

}

func TestLogin_Validate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userStore := inmemory.NewInMemoryUserStore()
	user := store.User{
		Email:    "test@example.com",
		Password: "password",
	}
	userStore.CreateUser(&user)
	r := gin.Default()
	handler := NewUserHandler(userStore, nil)
	AttachUserRoutes(handler, r)

	tests := []struct {
		body           map[string]string
		expectedStatus int
	}{
		{map[string]string{"email": "test@example.com", "password": "password"}, http.StatusOK},
		{map[string]string{"email": "invalid-email", "password": "password"}, http.StatusBadRequest},
		{map[string]string{"email": "test@example.com", "password": "wrongpassword"}, http.StatusUnauthorized},
	}

	for _, test := range tests {
		reqBody, _ := json.Marshal(test.body)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, test.expectedStatus, w.Code)
	}
}

func TestThirdParty(t *testing.T) {
	gin.SetMode(gin.TestMode)
	thirdparty := &thirdparty.ThirdParty{}
	userStore := inmemory.NewInMemoryUserStore()

	r := gin.Default()
	handler := NewUserHandler(userStore, thirdparty)
	AttachUserRoutes(handler, r)

	req, _ := http.NewRequest("POST", "/fetch", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}
