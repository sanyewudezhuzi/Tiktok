package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/tiktok/pkg/e"
	"github.com/sanyewudezhuzi/tiktok/pkg/u"
	"github.com/sanyewudezhuzi/tiktok/serializer"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstr := ctx.Query("token")
		if tokenstr == "" {
			return
		} else {
			claims, err := u.ParseToken(tokenstr)
			if err == nil && time.Now().Unix() < claims.ExpiresAt {
				ctx.Set("claims", claims)
			} else {
				ctx.JSON(http.StatusNotFound, serializer.Response{
					StatusCode: e.StatusCodeError,
					StatusMsg:  "获取 token 失败",
				})
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
