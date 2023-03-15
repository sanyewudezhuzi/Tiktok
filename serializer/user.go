package serializer

import (
	"github.com/sanyewudezhuzi/tiktok/dao/daoBasic"
	"github.com/sanyewudezhuzi/tiktok/model"
)

type User struct {
	Id               uint   `json:"id"`
	Name             string `json:"name"`
	Follow_count     int64  `json:"follow_count"`
	Follower_count   int64  `json:"follower_count"`
	Is_follow        bool   `json:"is_follow"`
	Avatar           string `json:"avatar"`
	Background_image string `json:"background_image"`
	Signature        string `json:"signature"`
	Total_favorited  int64  `json:"total_favorited"`
	Work_count       int64  `json:"work_count"`
	Favorite_count   int64  `json:"favorite_count"`
}

// SerializerUser 序列化 user
// u: user ;
// uid: token 的 uid ;
func SerializerUser(u model.User, uid uint) User {
	isfollow := false
	if u.ID == uid {
		isfollow = true
	} else {
		isfollow = daoBasic.GetIsFollowByUID(u.ID, uid)
	}
	return User{
		Id:               u.ID,
		Name:             u.Name,
		Follow_count:     u.FollowCount,
		Follower_count:   u.FollowerCount,
		Is_follow:        isfollow,
		Avatar:           u.Avatar,
		Background_image: u.BackgroundImage,
		Signature:        u.Signature,
		Total_favorited:  u.TotalFavorited,
		Work_count:       u.WorkCount,
		Favorite_count:   u.FavoriteCount,
	}
}
