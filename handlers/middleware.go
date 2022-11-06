package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx = "userId"
	passwordCtx = "password"
)


func (h *Handler) userAuthorization(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Auth header is empty")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 && headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusBadRequest, "Invalid auth header")
		return
	}

	userId, password, err := h.services.Authorization.ParseJWTToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if userId == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid auth token")
		return
	}

	c.Set(userCtx, userId)
	c.Set(passwordCtx, password)
} 


