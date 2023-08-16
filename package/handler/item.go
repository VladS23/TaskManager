package handler

import (
	"github.com/gin-gonic/gin"
	todo "myTaskManager"
	"net/http"
	"strconv"
)

// @Summary Create todo item
// @Security ApiKeyAuth
// @Tags items
// @Description create todo item
// @ID create-item

// @Accept  json
// @Produce  json

// @Param input body todo.TodoItem true "item info"
// @Param id path int true "List ID"

// @Success 200 {integer} integer 1

// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse

// @Router /api/lists/{id}/items/ [post]

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All todo items
// @Security ApiKeyAuth
// @Tags items
// @Description get all todo items
// @ID get-all-items

// @Accept  json
// @Produce  json

// @Param id path int true "List ID"

// @Success 200 {integer} integer 1

// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse

// @Router /api/lists/{id}/items [get]

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get todo item by id
// @Security ApiKeyAuth
// @Tags items
// @Description get todo item
// @ID get-item-by-id

// @Accept  json
// @Produce  json

// @Param id path int true "Item ID"

// @Success 200 {integer} integer 1

// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse

// @Router /api/items/{id} [get]

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update todo item
// @Security ApiKeyAuth
// @Tags items
// @Description update todo item
// @ID update-item

// @Accept  json
// @Produce  json

// @Param input body todo.UpdateItemInput true "item info"
// @Param id path int true "List ID"

// @Success 200 {integer} integer 1

// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse

// @Router /api/items/{id} [put]

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{Statuc: "ok"})
}

// @Summary Delete todo item by id
// @Security ApiKeyAuth
// @Tags items
// @Description delete todo item
// @ID delete-item-by-id

// @Accept  json
// @Produce  json

// @Param id path int true "Item ID"

// @Success 200 {integer} integer 1

// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse

// @Router /api/items/{id} [delete]

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{Statuc: "ok"})
}
