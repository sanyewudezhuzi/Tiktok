package srvInteractive

import (
	"fmt"
	"sync"

	"github.com/sanyewudezhuzi/tiktok/dao/daoBasic"
	"github.com/sanyewudezhuzi/tiktok/dao/daoInteractive"
	"github.com/sanyewudezhuzi/tiktok/model"
	"github.com/sanyewudezhuzi/tiktok/pkg/e"
	"github.com/sanyewudezhuzi/tiktok/serializer"
)

type Favorite struct {
	UID        uint
	VID        uint
	ActionType uint // 1-点赞，2-取消点赞
}

// Favorite 赞操作服务
func (s *Favorite) Favorite(uid uint) serializer.Response {
	// 判断 action_type
	var like bool
	if s.ActionType == 1 {
		like = true
	}

	// 判断是否已点赞或已取消赞
	favorite, _ := daoInteractive.GetFavoriteByUIDANDVID(uid, s.VID)
	if favorite.ID == 0 && !like {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "取消点赞失败",
		}
	} else if favorite.ID != 0 && like {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "点赞失败",
		}
	}

	tx := model.DB.Begin()

	// 更新用户喜欢数
	if err := daoInteractive.UpdateFavoriteCountByUID(uid, like); err != nil {
		tx.Rollback()
		fmt.Println("err:", err)
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "更新用户喜欢数失败",
			Data:       err,
		}
	}

	// 更新视频的点赞总数
	// 更新视频用户获赞总数
	if err := daoInteractive.UpdateFavoriteCountAndTotalFavoritedByVID(s.VID, like); err != nil {
		tx.Rollback()
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "更新获赞数失败",
			Data:       err,
		}
	}

	// 更新用户是否点赞视频
	if err := daoInteractive.UpdateFavoriteByUIDANDVID(uid, s.VID, like); err != nil {
		tx.Rollback()
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "更新用户点赞信息失败",
			Data:       err,
		}
	}

	// 返回响应
	tx.Commit()
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "赞操作服务成功",
	}
}

// FavoriteList 喜欢列表服务
func (s *Favorite) FavoriteList(uid uint) serializer.Response {
	// 验证 claims 的 id
	if uid != s.UID {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "用户身份信息错误",
		}
	}

	// 获取 videolist
	videolist, err := daoInteractive.GetVideoListByUID(uid)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取视频失败",
			Data:       err,
		}
	}

	// 封装视频
	wg := sync.WaitGroup{}
	favoritelist := make([]serializer.Video, len(videolist))
	for k := range videolist {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			user, _ := daoBasic.GetUserByUID(videolist[i].UID)
			favoritelist[i], _ = serializer.SerializerVideo(videolist[i], user, uid)
		}(k)
	}
	wg.Wait()

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "喜欢列表服务成功",
		Data:       favoritelist,
	}
}
