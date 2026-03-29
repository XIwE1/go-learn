package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
