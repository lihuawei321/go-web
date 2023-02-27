package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/bluebell/logic"
	"go.uber.org/zap"
	"strconv"
)

// 社区

// CommunityHandler 社区列表
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id,community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端报错暴漏给外面
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区详情
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id") //获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	communityList, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeSuccess, err.Error())
		return
	}
	ResponseSuccess(c, communityList)
}
