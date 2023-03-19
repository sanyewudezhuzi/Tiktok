package apiSocial

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/tiktok/pkg/u"
	"github.com/sanyewudezhuzi/tiktok/service/srvSocial"
)

// RelationAction 关注操作
func RelationAction(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "获取 claims 失败"})
		return
	}
	var relationActionService srvSocial.Follow
	followedUID, _ := strconv.Atoi(ctx.Query("to_user_id"))
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))
	relationActionService.FollowedUID = uint(followedUID)
	relationActionService.ActionType = uint(actionType)
	if err := ctx.ShouldBind(&relationActionService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, relationActionService.FollowAction(claims.(*u.Claims).ID))
	}
}

// RelationFollowList 关注列表
func RelationFollowList(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "获取 claims 失败"})
		return
	}
	var relationFollowListService srvSocial.Follow
	uid, _ := strconv.Atoi(ctx.Query("user_id"))
	relationFollowListService.UID = uint(uid)
	if err := ctx.ShouldBind(&relationFollowListService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, relationFollowListService.FollowList(claims.(*u.Claims).ID))
	}
}

// RelationFollowerList 粉丝列表
func RelationFollowerList(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "获取 claims 失败"})
		return
	}
	var relationFollowerListService srvSocial.Follow
	uid, _ := strconv.Atoi(ctx.Query("user_id"))
	relationFollowerListService.UID = uint(uid)
	if err := ctx.ShouldBind(&relationFollowerListService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, relationFollowerListService.FollowerList(claims.(*u.Claims).ID))
	}
}

// RelationFriendList 好友列表
func RelationFriendList(ctx *gin.Context) {

}

// MessageAction 发送消息
func MessageAction(ctx *gin.Context) {

}

// MessageChat 聊天记录
func MessageChat(ctx *gin.Context) {

}
