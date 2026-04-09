package middleware

import (
	"myproject/common/apperr"
	"myproject/common/httpx"
	"net/http"

	"github.com/gin-gonic/gin"
)

var expectedHost = "localhost:8080"

// 中间件（拦截器），功能：预处理，登录授权、验证、分页、耗时统计...
func AuthRequiredMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 通过自定义中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
		// ... check token, session, etc.
		ctx.Set("usersesion", "userid-1")
		ctx.Next() // 放行
		// ctx.Abort() // 阻止
	}
}

// 中间件 测试路由级中间件 校验id是否为有效的格式
func ValidateIdMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if len(id) > 5 {
			httpx.Fail(ctx, apperr.ErrBadRequest.Status, 102, "不合法，id不能大于5位数")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// 中间件 设置安全头
func HeaderMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Host != expectedHost {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		// 通过禁止页面在 <iframe> 中加载来防止点击劫持
		ctx.Header("X-Frame-Options", "DENY")
		// 控制浏览器允许加载哪些资源（脚本、样式、图片、字体等）以及从哪些来源
		ctx.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		ctx.Header("X-XSS-Protection", "1; mode=block")
		// 强制浏览器在指定的 max-age 期间对所有未来请求使用 HTTPS
		ctx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// 控制传出请求中发送多少引用者信息
		ctx.Header("Referrer-Policy", "strict-origin")
		// 防止 MIME 类型嗅探攻击，否则攻击者可能会将有害脚本伪装成无害文件
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		ctx.Next()
	}
}
