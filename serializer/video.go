package serializer

import (
	"github.com/sanyewudezhuzi/tiktok/dao/daoBasic"
	"github.com/sanyewudezhuzi/tiktok/model"
)

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

// SerializerVideo 序列化 PublishVideo
func SerializerPublishVideo(v model.Video, u model.User) Video {
	return Video{
		ID:            v.ID,
		Author:        SerializerUser(u, u.ID),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

// SerializerVideo 序列化 video :
// v: video ;
// vu: video 的 author ;
// u: token 的 uid ;
func SerializerVideo(v model.Video, vu model.User, uid uint) (Video, error) {
	isfavorite := daoBasic.GetIsFavoriteByUID(v.ID, uid)
	return Video{
		ID:            v.ID,
		Author:        SerializerUser(vu, uid),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    isfavorite,
		Title:         v.Title,
	}, nil
}

// SerializerVideos 序列化 list :
// v: video ;
// vu: video 的 author ;
// uid: token 的 uid ;
func SerializerList(l []model.Video, vu model.User, uid uint) []Video {
	v := make([]Video, len(l))
	for k := range l {
		v[k], _ = SerializerVideo(l[k], vu, uid)
	}
	return v
}
