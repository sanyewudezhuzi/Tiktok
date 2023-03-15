package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	VID uint `gorm:"unique"` // 点赞视频 ID
	UID uint `gorm:"unique"` // 点赞用户 ID
}
