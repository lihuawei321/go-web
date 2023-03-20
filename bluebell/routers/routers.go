package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go-web/bluebell/controller"
	_ "go-web/bluebell/docs"
	"go-web/bluebell/logger"
	"go-web/bluebell/middlewares"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/login", controller.LoginHandler)
	//登陆
	v1.POST("/signup", controller.SignUpHandler)

	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/post/:id", controller.PostDetailHandler)
	v1.GET("/posts/", controller.GetPostListHandler)
	//根据时间或分数获取帖子列表
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.Use(middlewares.JWTAuthMiddleware()) //应用JWT认证中间件

	{
		v1.POST("/post", controller.CreatePostHandler)
		//投票
		v1.POST("/vote", controller.PostVoteController)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
