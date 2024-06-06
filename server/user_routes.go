package server

import (
	"myapp/store"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userStore store.UserStore) *gin.Engine {
	r := gin.Default()

	h := NewUserHandler(userStore)

	AttachUserRoutes(h, r)

	return r
}

func AttachUserRoutes(h *UserHandler, r *gin.Engine) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}
