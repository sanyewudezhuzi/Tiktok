package model

import (
	"github.com/sanyewudezhuzi/tiktok/conf"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UID           uint   `gorm:"not null"` // 视频作者 ID
	PlayUrl       string `gorm:"unique"`   // 视频播放地址
	CoverUrl      string `gorm:"not null"` // 视频封面地址
	FavoriteCount int64  // 视频的点赞总数
	CommentCount  int64  // 视频的评论总数
	IsFavorite    bool   // true-已点赞，false-未点赞
	Title         string // 视频标题
}

// DefaultVideo 初始化 video
func (v *Video) DefaultVideo() {
	v.CoverUrl = conf.Host + conf.Port + conf.VideoPath + "default/" + "DefaultCoverUrl.jpg" // 此处将视频封面设置为默认值
	v.FavoriteCount = 0
	v.CommentCount = 0
	v.IsFavorite = false
}
