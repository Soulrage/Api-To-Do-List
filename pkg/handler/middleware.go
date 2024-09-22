package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"To-Do/internal/service"
)

const (
	authHeader = "Authorization"
	userCtx = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
	}

	userID, err := service.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
	}
	c.Set(userCtx, userID)
}
