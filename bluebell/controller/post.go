package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/bluebell/logic"
	"go-web/bluebell/models"
	"go.uber.org/zap"
)

// PostHandler 创建帖子功能
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数校验

	// c.ShouldBindJSON()
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	// 从 c 取到当前发请求的用户的ID值
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
