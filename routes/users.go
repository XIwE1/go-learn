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

func UserAuthorizedHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 进行某些判断
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "user接口需要携带token请求"})
		ctx.Abort()
	}
}

// 注册user相关的路由 在main里调用
func RegisterUserRoutes(router *gin.Engine) {
	users := router.Group("/api/users")

	// 分组中间件
	users.Use(UserAuthorizedHandler())

	users.POST("/create", createUser)
	users.DELETE("/delete/:id", deleteUser)
	users.PUT("/update", updateUser)
	users.GET("/:id", getUser)

}
