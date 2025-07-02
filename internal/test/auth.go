package test

import (
	"github.com/elibr-edu/gateway/internal/test/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	handle := NewHandler()

	r.GET("/ping", handle.Ping)
	r.GET("/failure", handle.Failure)
	r.GET("/rate", handle.Rate)

}

func NewHandler() *handler.Handler {
	return handler.NewHandler(nil)
}
