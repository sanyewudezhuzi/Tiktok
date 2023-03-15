package srvBasic

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/sanyewudezhuzi/tiktok/conf"
)

// CheckParameter 验证 username 和 password 是否合法
func checkParameter(username, password string) error {
	if len(username) > 32 || len(username) == 0 {
		return errors.New("用户名格式错误")
	}
	if len(password) > 32 || len(password) == 0 {
		return errors.New("密码格式错误")
	}
	return nil
}

// uploadVideo 上载视频到 static/video 目录下，并返回文件 path
func uploadVideo(file multipart.File, account, title string) (string, error) {
	// dirPath := conf.Host + conf.Port + conf.VideoPath + account + "/"
	// if err := createDir(dirPath); err != nil {
	// 	return "", err
	// }
	// path := dirPath + title + ".mp4"
	// content, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	return "", err
	// }
	// if err := ioutil.WriteFile(path, content, 0666); err != nil {
	// 	return "", err
	// }
	// return path, nil

	path := "." + conf.VideoPath + account + "/" + title + strconv.Itoa(int(time.Now().Unix())) + ".mp4"
	save, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer save.Close()
	io.Copy(save, file)
	return conf.Host + conf.Port + path[1:], nil
}

// createDir 创建文件夹
func createDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return os.MkdirAll(path, 755)
	}
	return nil
}
