package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorizaion")
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userId", userId)
}
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userID")
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "users id not found")
		return 0, errors.New("uer not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "users id not found")
		return 0, errors.New("uer not found")
	}
	return idInt, nil
}
