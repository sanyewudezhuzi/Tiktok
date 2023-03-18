package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowUID   uint `gorm:"not null"` // 关注用户 ID
	FollowedUID uint `gorm:"not null"` // 被关注用户 ID
}
