package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listUsers(c *gin.Context)  { c.JSON(http.StatusOK, gin.H{"action": "list_users"}) }
func createUser(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"action": "create_user"}) }
func getUser(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"action": "get_user"}) }
func updateUser(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"action": "update_user"}) }
func deleteUser(c *gin.Context) { c.Status(http.StatusNoContent) }

// 注册user相关的路由 在main里调用
func RegisterUserRoutes(router *gin.Engine) {
	users := router.Group("/api/users")

	users.POST("/create", createUser)
	users.DELETE("/delete/:id", deleteUser)
	users.PUT("/update", updateUser)
	users.GET("/:id", getUser)

}
