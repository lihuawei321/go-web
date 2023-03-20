package logic

import (
	"go-web/bluebell/dao/redis"
	"go-web/bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 投票功能
// 1.用户投票的数据
// 2.

// 投一票就加432分 86400/200 -> 需要200张赞成票可以给你的帖子续一天 -> 《redis实战》
/* 投票的几种情况
direction=1时，有2种情况：
	1. 之前没有投过票，现在投赞成票  --> 更新分数和投票记录
	2. 之前投反对票，现在改投赞成票
direction=0时，有2种情况：
	1. 之前投赞成票，现在要取消投票
	2. 之前投反对票，现在要取消投票
direction=-1时，有2种情况：
	1. 之前没有投过票，现在投反对票
	2. 之前投赞成票，现在改投反对票

投票的限制：
	每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
	1. 到期之后将redis中保存的赞成票数和反对票数存储到mysql表中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/
// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.String("postID", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
	// 1. 判断投票的限制
	// 2. 更新帖子的分数
	// 3. 记录用户为该帖子投票的数据
}
