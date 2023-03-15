package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	VID     uint   `gorm:"unique"` // 评论视频 ID
	UID     uint   `gorm:"unique"` // 评论用户 ID
	Comment string `gorm:"unique"` // 评论内容
}
