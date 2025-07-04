package auth

import (
	"github.com/elibr-edu/gateway/internal/auth/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	handle := NewHandler()

	r.POST("/login", handle.Login)
	r.POST("/refresh", handle.Refresh)
	r.GET("/ping", handle.Ping)
}

func NewHandler() *handler.Handler {
	return handler.NewHandler(nil)
}
