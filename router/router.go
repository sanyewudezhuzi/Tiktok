package router

import (
	"net/http"

	apiBasic "github.com/sanyewudezhuzi/tiktok/controller/apiBasic"
	apiInteractive "github.com/sanyewudezhuzi/tiktok/controller/apiInteractive"
	apiSocial "github.com/sanyewudezhuzi/tiktok/controller/apiSocial"
	"github.com/sanyewudezhuzi/tiktok/middleware"

	"github.com/gin-gonic/gin"
)

// 路由分组
func Router() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"ping": "pong"}) })
	d := r.Group("douyin").Use(middleware.JWT())
	{
		// 基础接口
		d.GET("feed/", apiBasic.Feed)
		d.POST("user/register", apiBasic.UserRegister)
		d.POST("user/login", apiBasic.UserLogin)
		d.GET("user", apiBasic.User)
		d.POST("publish/action", apiBasic.PublishAction)
		d.GET("publish/list", apiBasic.PublishList)

		// 互动接口
		d.POST("favorite/action", apiInteractive.FavoriteAction)
		d.GET("favorite/list", apiInteractive.FavoriteList)
		d.POST("comment/action", apiInteractive.CommentAction)
		d.GET("comment/list", apiInteractive.CommentList)

		// 社交接口
		d.POST("relation/action", apiSocial.RelationAction)
		d.GET("relation/follow/list", apiSocial.RelationFollowList)
		d.GET("relation/follower/list", apiSocial.RelationFollowerList)
		d.GET("relation/friend/list", apiSocial.RelationFriendList)
		d.POST("message/action", apiSocial.MessageAction)
		d.GET("message/chat", apiSocial.MessageChat)
	}

	return r
}
