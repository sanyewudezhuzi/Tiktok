package daoBasic

import "github.com/sanyewudezhuzi/tiktok/model"

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
