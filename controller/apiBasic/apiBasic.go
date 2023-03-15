package apiBasic

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/tiktok/pkg/u"
	"github.com/sanyewudezhuzi/tiktok/service/srvBasic"
)

// Feed 视频流接口
func Feed(ctx *gin.Context) {

}

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var userRegisterService srvBasic.User
	userRegisterService.Username = ctx.Query("username")
	userRegisterService.Password = ctx.Query("password")
	if err := ctx.ShouldBind(&userRegisterService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, userRegisterService.Register())
	}
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var userLoginService srvBasic.User
	userLoginService.Username = ctx.Query("username")
	userLoginService.Password = ctx.Query("password")
	if err := ctx.ShouldBind(&userLoginService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, userLoginService.Login())
	}
}

// User 用户信息
func User(ctx *gin.Context) {
	var userInfoService srvBasic.User
	id, _ := strconv.Atoi(ctx.Query("user_id"))
	userInfoService.UserID = uint(id)
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"err": "获取 claims 失败"})
		return
	}
	if err := ctx.ShouldBind(&userInfoService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, userInfoService.UserInfo(claims.(*u.Claims).ID))
	}
}

// PublishAction 视频投稿 ******
func PublishAction(ctx *gin.Context) {
	var publishListService srvBasic.Publish
	file, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	if err := ctx.ShouldBind(&publishListService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		claims, err := u.ParseToken(publishListService.Tokenstr)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"err": "获取 claims 失败"})
			return
		}
		fmt.Println("tokenstr:", publishListService.Tokenstr)
		fmt.Println("title", publishListService.Title)
		fmt.Println("file:", file)
		fmt.Println("claims:", claims)
		ctx.JSON(http.StatusOK, publishListService.Publish(claims.ID, file))
	}
}

// PublishList 发布列表
func PublishList(ctx *gin.Context) {
	var publishListService srvBasic.Publish
	id, _ := strconv.Atoi(ctx.Query("user_id"))
	publishListService.UserID = uint(id)
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"err": "获取 claims 失败"})
		return
	}
	if err := ctx.ShouldBind(&publishListService); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
	} else {
		ctx.JSON(http.StatusOK, publishListService.PublishList(claims.(*u.Claims).ID))
	}
}
