package daoSocial

import (
	"github.com/sanyewudezhuzi/tiktok/model"
)

// GetFollowInfoByUIDANDFollowedUID 通过关注者与被关注者的 UID 获取用户是否关注对方
func GetFollowInfoByUIDANDFollowedUID(uid, followedUID uint) (model.Follow, error) {
	var favoriteInfo model.Follow
	err := model.DB.Model(&model.Follow{}).Where("follow_uid = ? and followed_uid = ?", uid, followedUID).First(&favoriteInfo).Error
	return favoriteInfo, err
}

// UpdateFollowCountByFollowUID 更新关注者的关注总数
func UpdateFollowCountByFollowUID(uid uint, follow bool) error {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("id = ?", uid).First(&user).Error; err != nil {
		return err
	}
	if follow {
		user.FollowCount++
	} else {
		user.FollowCount--
	}
	return model.DB.Model(&model.User{}).Where("id = ?", uid).Update("follow_count", user.FollowCount).Error
}

// UpdateFollowerCountByFollowedUID 更新被关注者的粉丝总数
func UpdateFollowerCountByFollowedUID(followedUID uint, follow bool) error {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("id = ?", followedUID).First(&user).Error; err != nil {
		return err
	}
	if follow {
		user.FollowerCount++
	} else {
		user.FollowerCount--
	}
	return model.DB.Model(&model.User{}).Where("id = ?", followedUID).Update("follower_count", user.FollowerCount).Error
}

// UpdateFollow 更新关注者关注被关注者
func UpdateFollow(followUID, followedUID uint, follow bool) error {
	followInfo := model.Follow{
		FollowUID:   followUID,
		FollowedUID: followedUID,
	}
	if follow {
		return model.DB.Model(&model.Follow{}).Create(&followInfo).Error
	}
	return model.DB.Model(&model.Follow{}).Where("follow_uid = ? and followed_uid = ?", followUID, followedUID).Delete(&followInfo).Error
}
