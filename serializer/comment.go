package serializer

import (
	"github.com/sanyewudezhuzi/tiktok/model"
)

type Comment struct {
	ID         uint   `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

// SerializerComment 序列化 comment
// u: user ;
// uid: token 的 uid ;
func SerializerComment(c model.Comment, u model.User, uid uint) Comment {
	return Comment{
		ID:         c.ID,
		User:       SerializerUser(u, uid),
		Content:    c.Comment,
		CreateDate: c.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
