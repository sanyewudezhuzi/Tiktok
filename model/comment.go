package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	VID     uint   `gorm:"not null"` // 评论视频 ID
	UID     uint   `gorm:"not null"` // 评论用户 ID
	Comment string `gorm:"not null"` // 评论内容
}
