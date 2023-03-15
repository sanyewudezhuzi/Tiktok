package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowUID   uint `gorm:"unique"` // 关注用户 ID
	FollowedUID uint `gorm:"unique"` // 被关注用户 ID
}
