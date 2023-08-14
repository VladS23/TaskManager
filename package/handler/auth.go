package handler

import (
	"github.com/gin-gonic/gin"
	todo "myTaskManager"
	"net/http"
)

func (h *Handler) sighUp(c *gin.Context) {
	var input todo.User
	if err := c.Bind(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) sighIn(c *gin.Context) {

}
