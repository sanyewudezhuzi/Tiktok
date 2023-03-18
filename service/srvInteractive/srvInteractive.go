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

type Comment struct {
	VID         uint
	ActionType  uint   // 1-发布评论，2-删除评论
	CommentText string // 用户填写的评论内容，在action_type=1的时候使用
	CommentID   uint   // 要删除的评论id，在action_type=2的时候使用
}

// CommentAction 评论操作服务
func (s *Comment) CommentAction(uid uint) serializer.Response {
	// 判断 action_type
	var comment bool
	if s.ActionType == 1 {
		comment = true
	}

	// 判断是否已点赞或已取消赞
	text, err := daoInteractive.GetCommentByUIDANDVID(uid, s.VID)
	if text.ID == 0 && !comment {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "删除评论失败",
			Data:       err,
		}
	}

	tx := model.DB.Begin()

	// 更新视频的评论总数
	if err := daoInteractive.UpdateCommentCountByVID(s.VID, comment); err != nil {
		tx.Rollback()
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "更新视频的评论总数失败",
			Data:       err,
		}
	}

	// 获取用户信息
	user, err := daoBasic.GetUserByUID(uid)
	if err != nil {
		tx.Rollback()
		return serializer.Response{}
	}

	// 更新用户评论
	// 返回响应
	if comment {
		text = model.Comment{
			VID:     s.VID,
			UID:     uid,
			Comment: s.CommentText,
		}
		err, text = daoInteractive.CreateComment(text)
		if err != nil {
			tx.Rollback()
			return serializer.Response{
				StatusCode: e.StatusCodeError,
				StatusMsg:  "创建评论失败",
				Data:       err,
			}
		}
		tx.Commit()
		return serializer.Response{
			StatusCode: e.StatusCodeSuccess,
			StatusMsg:  "评论操作服务成功",
			Data:       serializer.SerializerComment(text, user, user.ID),
		}
	} else {
		if err := daoInteractive.DeleteComment(s.CommentID, text); err != nil {
			tx.Rollback()
			return serializer.Response{
				StatusCode: e.StatusCodeError,
				StatusMsg:  "删除评论失败",
				Data:       err,
			}
		}
		tx.Commit()
		return serializer.Response{
			StatusCode: e.StatusCodeSuccess,
			StatusMsg:  "评论操作服务成功",
			Data:       serializer.SerializerComment(text, user, user.ID),
		}
	}
}

// CommentList 评论列表服务
func (s *Comment) CommentList(uid uint) serializer.Response {
	// 获取评论列表
	commentList, err := daoInteractive.GetCommentListByVID(s.VID)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取评论列表失败",
			Data:       err,
		}
	}

	// 封装数据
	list := make([]serializer.Comment, len(commentList))
	wg := sync.WaitGroup{}
	for k := range commentList {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			user, _ := daoBasic.GetUserByUID(commentList[i].UID)
			list[i] = serializer.SerializerComment(commentList[i], user, uid)
		}(k)
	}
	wg.Wait()

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "评论列表服务成功",
		Data:       list,
	}
}
