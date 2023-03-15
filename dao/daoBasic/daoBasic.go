package daoBasic

import (
	"time"

	"github.com/sanyewudezhuzi/tiktok/conf"
	"github.com/sanyewudezhuzi/tiktok/model"
)

// GetFeedList 获取视频流
func GetFeedList(time time.Time) ([]model.Video, error) {
	feedlist := make([]model.Video, 0, conf.FeedCount)
	err := model.DB.Model(&model.Video{}).Where("created_at <= ?", time.Format("2006-01-02 15:04:05")).
		Order("created_at desc").Limit(conf.FeedCount).
		Find(&feedlist).Error
	return feedlist, err
}

// GetIsFavoriteByUID 通过 UID 获取用户是否点赞视频
// vid: 视频作者
// uid: token 的用户
func GetIsFavoriteByUID(vid, uid uint) bool {
	var favorite model.Favorite
	model.DB.Model(&model.Favorite{}).Where("v_id = ? and uid = ?", vid, uid).First(&favorite)
	return favorite.ID != 0
}

// GetIsFollowByUID 通过 UID 获取用户是否关注用户
// fid: 被关注的用户
// uid：token 的用户
func GetIsFollowByUID(fid, uid uint) bool {
	var follow model.Follow
	model.DB.Model(&model.Follow{}).Where("follow_uid = ? and followed_uid = ?", uid, fid).First(&follow)
	return follow.ID != 0
}

// ExistOrNotByAccount 根据用户名判断用户是否注册
func ExistOrNotByAccount(account string) (model.User, bool) {
	var user model.User
	model.DB.Model(&model.User{}).Where("account = ?", account).First(&user)
	if user.ID == 0 {
		return user, false
	}
	return user, true
}

// CreateUser 创建用户
func CreateUser(user *model.User) error {
	return model.DB.Model(&model.User{}).Create(&user).Error
}

// GetUser 通过 UID 获取用户
func GetUserByUID(uid uint) (model.User, error) {
	var user model.User
	err := model.DB.Model(&model.User{}).Where("id = ?", uid).First(&user).Error
	return user, err
}

// CreateVideo 创建视频
func CreateVideo(video *model.Video) error {
	return model.DB.Model(&model.Video{}).Create(&video).Error
}

// GetListByUID 通过 UID 获取发布列表
func GetListByUID(uid uint) ([]model.Video, error) {
	var list []model.Video
	err := model.DB.Model(&model.Video{}).Where("uid = ?", uid).Find(&list).Error
	return list, err
}
