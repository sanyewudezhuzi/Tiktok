package srvBasic

import (
	"fmt"
	"mime/multipart"

	"github.com/sanyewudezhuzi/tiktok/dao/daoBasic"
	"github.com/sanyewudezhuzi/tiktok/model"
	e "github.com/sanyewudezhuzi/tiktok/pkg/e"
	"github.com/sanyewudezhuzi/tiktok/pkg/u"
	"github.com/sanyewudezhuzi/tiktok/serializer"
)

type Basic struct {
	Username string
	Password string
	UserID   uint
}

// Register 用户注册服务
func (s *Basic) Register() serializer.Response {
	// 验证 username 和 password 是否合法
	if err := checkParameter(s.Username, s.Password); err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  err.Error(),
		}
	}
	user, exist := daoBasic.ExistOrNotByAccount(s.Username)
	if exist {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "用户名已注册",
		}
	}

	// 封装数据
	user.Account = s.Username
	user.DefaultUser()
	if err := user.Bcrypt(s.Password); err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "注册用户失败",
			Data:       err,
		}
	}

	// 数据持久化
	if err := daoBasic.CreateUser(&user); err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "注册用户失败",
			Data:       err,
		}
	}

	// 签发 token
	tokenstr, err := u.GenerateToken(user.ID, user.Account)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "签发 token 失败",
			Data:       err,
		}
	}

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "用户注册服务成功",
		Data: map[string]interface{}{
			"user_id": user.ID,
			"token":   tokenstr,
		},
	}
}

// Login 用户登录服务
func (s *Basic) Login() serializer.Response {
	// 验证 username 和 password 是否合法
	if err := checkParameter(s.Username, s.Password); err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  err.Error(),
		}
	}
	user, exist := daoBasic.ExistOrNotByAccount(s.Username)
	if !exist {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "用户名未注册",
		}
	}
	if !user.Check(s.Password) {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "密码错误",
		}
	}

	// 签发 token
	tokenstr, err := u.GenerateToken(user.ID, user.Account)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "签发 token 失败",
			Data:       err,
		}
	}

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "用户登录服务成功",
		Data: map[string]interface{}{
			"user_id": user.ID,
			"token":   tokenstr,
		},
	}
}

// UserInfo 用户信息服务
func (s *Basic) UserInfo(uid uint) serializer.Response {
	// 验证 claims 的 id
	if uid != s.UserID {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "用户身份信息错误",
		}
	}

	// 获取 user 信息
	user, err := daoBasic.GetUserByUID(uid)
	fmt.Println("user:", user)
	fmt.Println("err:", err)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取用户信息失败",
			Data:       err,
		}
	}

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "用户信息服务成功",
		Data: map[string]interface{}{
			"user": serializer.SerializerUser(user),
		},
	}
}

type Video struct {
	UserID   uint
	Tokenstr string `form:"token"`
	Title    string `form:"title"`
}

// Publish 视频投稿服务
func (s *Video) Publish(uid uint, file *multipart.FileHeader) serializer.Response {
	// 获取 user
	user, err := daoBasic.GetUserByUID(uid)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取用户信息失败",
			Data:       err,
		}
	}

	// 上传文件到 static
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取视频文件失败",
			Data:       err,
		}
	}
	path, err := uploadVideo(f, user.Account, s.Title)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "上传视频文件失败",
			Data:       err,
		}
	}

	// 封装数据
	video := model.Video{
		UID:     uid,
		PlayUrl: path,
		// CoverUrl: conf.Host + conf.Port + conf.VideoPath + "xxx",
		Title: s.Title,
	}
	video.DefaultVideo()

	// 数据持久化
	if err := daoBasic.CreateVideo(&video); err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "视频投稿失败",
			Data:       err,
		}
	}

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "视频投稿服务成功",
		Data:       serializer.SerializerVideo(video, user),
	}
}

// PublishList 发布列表服务
func (s *Video) PublishList(uid uint) serializer.Response {
	// 验证 claims 的 id
	if uid != s.UserID {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "用户身份信息错误",
		}
	}

	// 获取 user
	user, err := daoBasic.GetUserByUID(uid)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取用户信息失败",
			Data:       err,
		}
	}

	// 获取 list
	list, err := daoBasic.GetListByUID(uid)
	if err != nil {
		return serializer.Response{
			StatusCode: e.StatusCodeError,
			StatusMsg:  "获取发布列表失败",
			Data:       err,
		}
	}

	// 返回响应
	return serializer.Response{
		StatusCode: e.StatusCodeSuccess,
		StatusMsg:  "发布列表服务成功",
		Data:       serializer.SerializerList(list, user),
	}
}
