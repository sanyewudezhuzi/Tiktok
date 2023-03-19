package srvSocial

import (
	"sync"

	"github.com/sanyewudezhuzi/tiktok/dao/daoSocial"
	"github.com/sanyewudezhuzi/tiktok/model"
	"github.com/sanyewudezhuzi/tiktok/pkg/e"
	"github.com/sanyewudezhuzi/tiktok/serializer"
)

type Follow struct {
	UID         uint
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

// FollowList 关注列表服务
func (s *Follow) FollowList(uid uint) serializer.Response {
	// 验证 claims 的 id
	if uid != s.UID {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "用户身份信息错误",
		}
	}

	// 获取关注信息列表
	followInfo, err := daoSocial.GetFollowInfoListByUID(uid)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取关注信息失败",
			Data:       err,
		}
	}

	// 获取关注列表
	wg := sync.WaitGroup{}
	userList := make([]model.User, len(followInfo))
	for k := range followInfo {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			userList[i], _ = daoSocial.GetUserByFollowInfo(followInfo[i])
		}(k)
	}
	wg.Wait()

	// 封装数据
	followList := make([]serializer.User, len(userList))
	for k := range userList {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			followList[i] = serializer.SerializerUser(userList[i], uid)
		}(k)
	}
	wg.Wait()

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "关注列表服务成功",
		Data:       followList,
	}
}

// FollowerList 粉丝列表服务
func (s *Follow) FollowerList(uid uint) serializer.Response {
	// 验证 claims 的 id
	if uid != s.UID {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "用户身份信息错误",
		}
	}

	// 获取粉丝信息列表
	followerInfo, err := daoSocial.GetFollowerInfoListByUID(uid)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取粉丝信息失败",
			Data:       err,
		}
	}

	// 获取粉丝列表
	wg := sync.WaitGroup{}
	userList := make([]model.User, len(followerInfo))
	for k := range followerInfo {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			userList[i], _ = daoSocial.GetUserByFollowerInfo(followerInfo[i])
		}(k)
	}
	wg.Wait()

	// 封装数据
	followList := make([]serializer.User, len(userList))
	for k := range userList {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			followList[i] = serializer.SerializerUser(userList[i], uid)
		}(k)
	}
	wg.Wait()

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "粉丝列表服务成功",
		Data:       followList,
	}
}
