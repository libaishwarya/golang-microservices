package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/libaishwarya/myapp/catservice"
	"github.com/libaishwarya/myapp/store"
	"github.com/libaishwarya/myapp/userservice"
)

type UserHandler struct {
	store      store.UserStore
	thirdparty userservice.JsonPlaceholder
	catFact    catservice.CatFactService
	validate   *validator.Validate
}

func NewUserHandler(store store.UserStore, thirdparty userservice.JsonPlaceholder, catFact catservice.CatFactService) *UserHandler {
	return &UserHandler{store: store, thirdparty: thirdparty, catFact: catFact, validate: validator.New()}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user store.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var user store.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validate.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedUser, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user"})
		return
	}

	if storedUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (h *UserHandler) Store(c *gin.Context) {
	thirdPartyUsers, err := h.thirdparty.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "failed to get users"})
		return
	}

	if len(thirdPartyUsers) == 2 {
		c.JSON(http.StatusConflict, gin.H{"message": "only two user found"})
		return
	}

	if len(thirdPartyUsers) > 20 {
		c.JSON(http.StatusConflict, gin.H{"message": "more users found"})
		return
	}

	var values []store.ExternalUser
	for _, value := range thirdPartyUsers {
		values = append(values, store.ExternalUser{
			ID:    value.ID,
			Name:  value.Name,
			Email: value.Email,
		})

	}

	for _, value := range values {
		if err := h.store.StoreRes(&value); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not store the fetched data"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fetched user data is stored successfully"})

}

func (h *UserHandler) StoreCatFact(c *gin.Context) {
	fact, err := h.catFact.GetCatFact()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch cat fact"})
		return
	}

	storeFact := &store.CatFact{
		Fact:   fact.Fact,
		Length: fact.Length,
	}

	if err := h.store.StoreCatFact(storeFact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store the cat fact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cat fact stored successfully"})
}
