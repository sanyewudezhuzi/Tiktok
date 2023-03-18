package apiInteractive

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/tiktok/pkg/u"
	"github.com/sanyewudezhuzi/tiktok/service/srvInteractive"
)

// FavoriteAction 赞操作
func FavoriteAction(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "获取 claims 失败"})
		return
	}
	var favoriteActionService srvInteractive.Favorite
	vid, _ := strconv.Atoi(ctx.Query("video_id"))
	actionType, _ := strconv.Atoi(ctx.Query("action_type"))
	favoriteActionService.VID = uint(vid)
	favoriteActionService.ActionType = uint(actionType)
	if err := ctx.ShouldBind(&favoriteActionService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, favoriteActionService.Favorite(claims.(*u.Claims).ID))
	}
}

// FavoriteList 喜欢列表
func FavoriteList(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "获取 claims 失败"})
		return
	}
	var favoriteListService srvInteractive.Favorite
	uid, _ := strconv.Atoi(ctx.Query("user_id"))
	favoriteListService.UID = uint(uid)
	if err := ctx.ShouldBind(&favoriteListService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, favoriteListService.FavoriteList(claims.(*u.Claims).ID))
	}
}

// CommentAction 评论操作
func CommentAction(ctx *gin.Context) {

}

// CommentList 评论列表
func CommentList(ctx *gin.Context) {

}
