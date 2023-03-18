package daoInteractive

import (
	"github.com/sanyewudezhuzi/tiktok/model"
)

// GetFavoriteByUIDANDVID 通过 UID 和 VID 获取点赞信息
func GetFavoriteByUIDANDVID(uid, vid uint) (model.Favorite, error) {
	var favorite model.Favorite
	err := model.DB.Model(&model.Favorite{}).Where("uid = ? and v_id = ?", uid, vid).First(&favorite).Error
	return favorite, err
}

// UpdateFavoriteCountByUID 通过 UID 更新用户喜欢数
func UpdateFavoriteCountByUID(uid uint, like bool) error {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("id = ?", uid).First(&user).Error; err != nil {
		return err
	}
	if like {
		user.FavoriteCount++
	} else {
		user.FavoriteCount--
	}
	return model.DB.Model(&model.User{}).Where("id = ?", uid).Update("favorite_count", user.FavoriteCount).Error
}

// UpdateTotalFavoritedByVID 通过 VID 更新视频用户获赞总数
func UpdateFavoriteCountAndTotalFavoritedByVID(vid uint, like bool) error {
	var video model.Video
	var user model.User
	if err := model.DB.Model(&model.Video{}).Where("id = ?", vid).First(&video).Error; err != nil {
		return err
	}
	if err := model.DB.Model(&model.User{}).Where("id = ?", video.UID).First(&user).Error; err != nil {
		return err
	}
	if like {
		video.FavoriteCount++
		user.TotalFavorited++
	} else {
		video.FavoriteCount--
		user.TotalFavorited--
	}
	if err := model.DB.Model(&model.Video{}).Where("id = ?", vid).Update("favorite_count", video.FavoriteCount).Error; err != nil {
		return err
	}
	return model.DB.Model(&model.User{}).Where("id = ?", user.ID).Update("total_favorited", user.TotalFavorited).Error
}

// UpdateFavoriteByUIDANDVID 通过 UID 和 VID 更新点赞信息
func UpdateFavoriteByUIDANDVID(uid, vid uint, like bool) error {
	var favorite model.Favorite = model.Favorite{UID: uid, VID: vid}
	if like {
		return model.DB.Model(&model.Favorite{}).Create(&favorite).Error
	} else {
		return model.DB.Model(&model.Favorite{}).Where("uid = ? and v_id = ?", uid, vid).Delete(&favorite).Error
	}
}

// GetVideoListByUID 通过 UID 获取喜欢视频列表
func GetVideoListByUID(uid uint) ([]model.Video, error) {
	var favoritelist []model.Favorite
	if err := model.DB.Model(&model.Favorite{}).Where("uid = ?", uid).Order("updated_at desc").Find(&favoritelist).Error; err != nil {
		return nil, err
	}
	videolist := make([]model.Video, len(favoritelist))
	for k := range favoritelist {
		videolist[k], _ = GetVideoByVID(favoritelist[k].VID)
	}
	return videolist, nil
}

// GetVideoByVID 通过 VID 获取视频
func GetVideoByVID(vid uint) (model.Video, error) {
	var video model.Video
	err := model.DB.Model(&model.Video{}).Where("id = ?", vid).First(&video).Error
	return video, err
}

// GetCommentByUIDANDVID 通过 UID 和 VID 获取评论信息
func GetCommentByUIDANDVID(uid, vid uint) (model.Comment, error) {
	var comment model.Comment
	err := model.DB.Model(&model.Comment{}).Where("uid = ? and v_id = ?", uid, vid).First(&comment).Error
	return comment, err
}

// UpdateCommentCountByVID 通过 VID 更新视频的评论总数
func UpdateCommentCountByVID(vid uint, comment bool) error {
	var video model.Video
	if err := model.DB.Model(&model.Video{}).Where("id = ?", vid).First(&video).Error; err != nil {
		return err
	}
	if comment {
		video.CommentCount++
	} else {
		video.CommentCount--
	}
	return model.DB.Model(&model.Video{}).Where("id = ?", vid).Update("comment_count", video.CommentCount).Error
}

// CreateComment 创建评论
func CreateComment(text model.Comment) (error, model.Comment) {
	return model.DB.Model(&model.Comment{}).Create(&text).Error, text
}

// DeleteComment 删除评论
func DeleteComment(commentID uint, text model.Comment) error {
	return model.DB.Model(&model.Comment{}).Where("id = ?", commentID).Delete(&text).Error
}

// GetCommentListByVID 根据 VID 获取评论列表
func GetCommentListByVID(vid uint) ([]model.Comment, error) {
	var commentList []model.Comment
	err := model.DB.Model(&model.Comment{}).Order("created_at desc").Where("v_id = ?", vid).Find(&commentList).Error
	return commentList, err
}
