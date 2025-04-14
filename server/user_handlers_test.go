package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/libaishwarya/myapp/catservice/mockcatfact"
	"github.com/libaishwarya/myapp/store"
	"github.com/libaishwarya/myapp/store/inmemory"
	"github.com/libaishwarya/myapp/userservice"
	mockthirdparty "github.com/libaishwarya/myapp/userservice/mock_third_party"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userStore := inmemory.NewInMemoryUserStore()
	r := gin.Default()
	handler := NewUserHandler(userStore, nil, nil)
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
	handler := NewUserHandler(userStore, nil, nil)
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
	handler := NewUserHandler(userStore, nil, nil)
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
	handler := NewUserHandler(userStore, nil, nil)
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

// func TestThirdParty_Success(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	thirdparty := &mockthirdparty.MockThirdParty{}
// 	userStore := inmemory.NewInMemoryUserStore()

// 	r := gin.Default()
// 	handler := NewUserHandler(userStore, thirdparty)
// 	AttachUserRoutes(handler, r)

// 	req, _ := http.NewRequest("POST", "/fetch", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// }

// func TestThirdParty_Failure(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	thirdparty := &mockthirdparty.MockThirdParty{
// 		Fail: true,
// 	}
// 	userStore := inmemory.NewInMemoryUserStore()

// 	r := gin.Default()
// 	handler := NewUserHandler(userStore, thirdparty)
// 	AttachUserRoutes(handler, r)

// 	req, _ := http.NewRequest("POST", "/fetch", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusBadGateway, w.Code)

// 	var body map[string]interface{}
// 	err := json.NewDecoder(w.Body).Decode(&body)
// 	assert.NoError(t, err)

// 	assert.Equal(t, "failed to get users", body["message"])
// }

// func TestThirdParty_Failure_MoreUsers(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	users := []userservice.ThirdPartyUser{}
// 	for i := 0; i < 25; i++ {
// 		users = append(users, userservice.ThirdPartyUser{
// 			Name:  "test",
// 			Email: "test",
// 		})
// 	}

// 	thirdparty := &mockthirdparty.MockThirdParty{
// 		Users: users,
// 	}
// 	userStore := inmemory.NewInMemoryUserStore()

// 	r := gin.Default()
// 	handler := NewUserHandler(userStore, thirdparty)
// 	AttachUserRoutes(handler, r)

// 	req, _ := http.NewRequest("POST", "/fetch", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusConflict, w.Code)

// 	var body map[string]interface{}
// 	err := json.NewDecoder(w.Body).Decode(&body)
// 	assert.NoError(t, err)

// 	assert.Equal(t, "more users found", body["message"])
// }

// func TestThirdParty_Failure_TwoUsers(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	users := []userservice.ThirdPartyUser{}
// 	for i := 0; i < 2; i++ {
// 		users = append(users, userservice.ThirdPartyUser{
// 			Name:  "test",
// 			Email: "test",
// 		})
// 	}

// 	thirdparty := &mockthirdparty.MockThirdParty{
// 		Users: users,
// 	}
// 	userStore := inmemory.NewInMemoryUserStore()

// 	r := gin.Default()
// 	handler := NewUserHandler(userStore, thirdparty)
// 	AttachUserRoutes(handler, r)

// 	req, _ := http.NewRequest("POST", "/fetch", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusConflict, w.Code)

// 	var body map[string]interface{}
// 	err := json.NewDecoder(w.Body).Decode(&body)
// 	assert.NoError(t, err)

// 	assert.Equal(t, "only two users found", body["message"])
// }

func TestThirdParty_New(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name            string
		Fail            bool
		MockUsersNumber int
		ExpectedStatus  int
		ExpectedMessage string
	}{
		{
			"success",
			false,
			1,
			http.StatusOK,
			"Fetched user data is stored successfully",
		},
		{
			"two users",
			false,
			2,
			http.StatusConflict,
			"only two users found",
		},
		{
			"twenty plus users",
			false,
			21,
			http.StatusConflict,
			"more users found",
		},
		{
			"get users failure",
			true,
			1,
			http.StatusBadGateway,
			"failed to get users",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			users := []userservice.ThirdPartyUser{}
			for i := 0; i < tt.MockUsersNumber; i++ {
				users = append(users, userservice.ThirdPartyUser{
					Name:  "test",
					Email: "test",
				})
			}

			thirdparty := &mockthirdparty.MockThirdParty{
				Users: users,
				Fail:  tt.Fail,
			}
			userStore := inmemory.NewInMemoryUserStore()

			r := gin.Default()
			handler := NewUserHandler(userStore, thirdparty, nil)
			AttachUserRoutes(handler, r)

			req, _ := http.NewRequest("POST", "/fetch", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			var body map[string]interface{}
			err := json.NewDecoder(w.Body).Decode(&body)
			assert.NoError(t, err)

			assert.Equal(t, tt.ExpectedStatus, w.Code)
			assert.Equal(t, tt.ExpectedMessage, body["message"])

		})

	}
}

func TestFetchCatFact(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		Name            string
		Fail            bool
		ExpectedStatus  int
		ExpectedMessage string
	}{
		{
			"success",
			false,
			http.StatusOK,
			"Cat fact stored successfully",
		},
		{
			"failure",
			true,
			http.StatusInternalServerError,
			"Could not fetch cat fact",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockCatFactService := &mockcatfact.MockCatFact{Fail: tt.Fail}

			userStore := inmemory.NewInMemoryUserStore()

			handler := NewUserHandler(userStore, nil, mockCatFactService)
			router := gin.Default()
			AttachUserRoutes(handler, router)

			req, _ := http.NewRequest("GET", "/fetchCat", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.ExpectedStatus, resp.Code)

			var body map[string]interface{}
			err := json.NewDecoder(resp.Body).Decode(&body)
			assert.NoError(t, err)

			assert.Equal(t, tt.ExpectedMessage, body["message"])

		})
	}
}
