package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/bluebell/logic"
	"go-web/bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// PostHandler 创建帖子功能
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数校验

	// c.ShouldBindJSON()
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post handler err")
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	// 从 c 取到当前发请求的用户的ID值
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	p.AuthorId = userID
	// 2. 创建帖子功能
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// PostDetailHandler 帖子详情
func PostDetailHandler(c *gin.Context) {
	// 1.获取参数（从URL中获取帖子的id)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.根据ID取出帖子数据（查数据库）
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPost(postID) failed", zap.Error(err))
	}
	// 3.返回响应
	ResponseSuccess(c, data)
}
