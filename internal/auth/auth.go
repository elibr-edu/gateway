package auth

import (
	"github.com/elibr-edu/gateway/internal/auth/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	handl := NewHandler()

	r.POST("/login", handl.Login)
	r.POST("/refresh", handl.Refresh)
	r.GET("/ping", handl.Ping)
}

func NewHandler() *handler.Handler {
	return handler.NewHandler(nil)
}
