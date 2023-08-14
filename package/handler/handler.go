package handler

import (
	"github.com/gin-gonic/gin"
	"myTaskManager/package/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sigh-up", h.sighUp)
		auth.POST("/sigh-in", h.sighIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/items_id", h.getItemById)
				items.PUT("/items_id", h.updateItem)
				items.DELETE("/items_id", h.deleteItem)
			}
		}
	}
	return router
}
