package handler

import (
	"github.com/gin-gonic/gin"
	todo "myTaskManager"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getAllLists(c *gin.Context) {

}
func (h *Handler) getListById(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
