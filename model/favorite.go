package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	VID uint `gorm:"not null"` // 点赞视频 ID
	UID uint `gorm:"not null"` // 点赞用户 ID
}
