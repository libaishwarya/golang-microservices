package server

import (
	"github.com/libaishwarya/myapp/store"
	"github.com/libaishwarya/myapp/userservice"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userStore store.UserStore, userService userservice.JsonPlaceholder) *gin.Engine {
	r := gin.Default()

	h := NewUserHandler(userStore, userService)

	AttachUserRoutes(h, r)

	return r
}

func AttachUserRoutes(h *UserHandler, r *gin.Engine) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/fetch", h.Store)
	r.GET("/fetchCat", h.StoreCatFact)

}
