package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var func1,func2,func3,func4,func5 func(c *gin.Context)
func main(){
	r:=gin.Default()

	r.GET("/hello",func(c *gin.Context){
		c.String(http.StatusOK,"ok")
	})

	shopGroup := r.Group("/shop",func1,func2)
	shopGroup.Use(func3)
	{
		shopGroup.GET("/index",func4,func5)
	}

	r.Run(":9999")
}
