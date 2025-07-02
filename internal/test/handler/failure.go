package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Failure(c *gin.Context) {
	c.String(500, "failure")
}
