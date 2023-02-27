package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/bluebell/controller"
	"go-web/bluebell/logger"
	"go-web/bluebell/middlewares"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/login", controller.LoginHandler)
	//登陆
	v1.POST("/signup", controller.SignUpHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) //应用JWT认证中间件

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
