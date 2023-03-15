package model

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/sanyewudezhuzi/tiktok/conf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account         string `gorm:"unique"` // 账号
	Password        string `gorm:"unique"` // 密码
	Name            string // 昵称
	FollowCount     int64  // 关注总数
	FollowerCount   int64  // 粉丝总数
	IsFollow        bool   // true-已关注，false-未关注
	Avatar          string // 用户头像
	BackgroundImage string // 用户个人页顶部大图
	Signature       string // 个人简介
	TotalFavorited  int64  // 获赞数量
	WorkCount       int64  // 作品数
	FavoriteCount   int64  // 喜欢数
}

// DefaultUser 用户初始化
func (u *User) DefaultUser() {
	rand.Seed(time.Now().Unix())
	u.Name = "用户" + strconv.Itoa(rand.Intn(900000000)+100000000)
	u.FollowCount = 0
	u.FollowerCount = 0
	u.IsFollow = true
	u.Avatar = conf.Host + conf.Port + conf.ImgPath + "default/" + "DefaultAvatar" + strconv.Itoa(rand.Intn(4)) + ".jpg"
	u.BackgroundImage = conf.Host + conf.Port + conf.ImgPath + "default/" + "DefaultBackgroundImage.jpg"
	u.Signature = "这个人很懒，什么都没有写。"
	u.TotalFavorited = 0
	u.WorkCount = 0
	u.FavoriteCount = 0
}

// Bcrypt 密码加密
func (u *User) Bcrypt(passowrd string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(passowrd), 10)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// Check 检验密码是否正确
func (u *User) Check(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
