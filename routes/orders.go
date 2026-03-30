package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type GMOrder struct {
	Target string `json:"target"`
}

type PlayerOrder struct {
	Id string `json:"id"`
}

func executeOrder(c *gin.Context) {
	var gmOrders GMOrder
	var PlayerOrders PlayerOrder

	// × ShouldBind  会消费 c.Request.Body，导致后续ShouldBindxx失效
	// if errGM := c.ShouldBind(&gmOrders); errGM != nil {
	// √ 使用 ShouldBindBodyWith 它会读取一次请求体并将其存储在上下文中，后续的绑定重用请求体
	if errGM := c.ShouldBindBodyWith(&gmOrders, binding.JSON); errGM != nil {
		// ...
	}
	if errPlayer := c.ShouldBindBodyWith(&PlayerOrders, binding.JSON); errPlayer != nil {
		// ...
	}
}

func listOrders(c *gin.Context)  { c.JSON(http.StatusOK, gin.H{"action": "list_orders"}) }
func createOrder(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"action": "create_order"}) }
func getOrder(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"action": "get_order"}) }

func RegisterOrderRoutes(router *gin.Engine) {
	orders := router.Group("/api/orders")
	{
		orders.GET("/", listOrders)
		orders.POST("/create", createOrder)
		orders.GET("/:id", getOrder)
	}
}
