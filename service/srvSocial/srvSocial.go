package srvSocial

import (
	"github.com/sanyewudezhuzi/tiktok/dao/daoSocial"
	"github.com/sanyewudezhuzi/tiktok/model"
	"github.com/sanyewudezhuzi/tiktok/pkg/e"
	"github.com/sanyewudezhuzi/tiktok/serializer"
)

type Follow struct {
	FollowedUID uint
	ActionType  uint // 1-关注，2-取消关注
}

// FollowAction 关注操作服务
func (s *Follow) FollowAction(uid uint) serializer.Response {
	// 判断 action_type
	var follow bool
	if s.ActionType == 1 {
		follow = true
	}

	// 判断是否已关注或已取消关注
	followInfo, err := daoSocial.GetFollowInfoByUIDANDFollowedUID(uid, s.FollowedUID)
	if followInfo.ID == 0 && !follow {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "取消关注失败",
			Data:       err,
		}
	} else if followInfo.ID != 0 && follow {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "关注失败",
			Data:       err,
		}
	}

	tx := model.DB.Begin()

	// 更新关注者的关注总数
	if err := daoSocial.UpdateFollowCountByFollowUID(uid, follow); err != nil {
		tx.Rollback()
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "更新关注者的关注总数失败",
			Data:       err,
		}
	}

	// 更新被关注者的粉丝总数
	if err := daoSocial.UpdateFollowerCountByFollowedUID(s.FollowedUID, follow); err != nil {
		tx.Rollback()
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "更新被关注者的粉丝总数失败",
			Data:       err,
		}
	}

	// 更新关注者关注被关注者
	if err := daoSocial.UpdateFollow(uid, s.FollowedUID, follow); err != nil {
		tx.Rollback()
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "更新关注者关注被关注者",
			Data:       err,
		}
	}

	// 返回响应
	tx.Commit()
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "关注操作服务成功",
	}
}
