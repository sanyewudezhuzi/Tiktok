package serializer

import "github.com/sanyewudezhuzi/tiktok/model"

type Video struct {
	ID            uint   `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

// SerializerVideo 序列化 video
func SerializerVideo(v model.Video, u model.User) Video {
	return Video{
		ID:            v.ID,
		Author:        SerializerUser(u),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

// SerializerVideos 序列化 list
func SerializerList(l []model.Video, u model.User) []Video {
	v := make([]Video, len(l))
	for k := range l {
		v[k] = SerializerVideo(l[k], u)
	}
	return v
}
