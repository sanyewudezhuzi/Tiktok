package conf

import (
	"github.com/go-ini/ini"
)

var (
	// server
	AppMode string
	Port    string
	// mysql
	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	// redis
	RedisDB     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
	// path
	Host      string
	ImgPath   string
	VideoPath string
	// common
	SecretKey string
)

// LoadEnvironment 加载环境变量
func LoadEnvironment() {
	file := load_ini()
	load(file)
}

// load_ini 读取 .ini 配置文件
func load_ini() *ini.File {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		panic("Profile path error.")
	}
	return file
}

// load 加载全局变量
func load(f *ini.File) {
	loadServer(f)
	loadMySQL(f)
	loadRedis(f)
	loadPath(f)
	loadCommon(f)
}

// loadServer 加载 .ini 文件的 server 项
func loadServer(f *ini.File) {
	AppMode = f.Section("server").Key("AppMode").String()
	Port = f.Section("server").Key("Port").String()
}

// loadMySQL 加载 .ini 文件的 mysql 项
func loadMySQL(f *ini.File) {
	DB = f.Section("mysql").Key("DB").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassword = f.Section("mysql").Key("DbPassword").String()
	DbName = f.Section("mysql").Key("DbName").String()
}

// loadRedis 加载 .ini 文件的 redis 项
func loadRedis(f *ini.File) {
	RedisDB = f.Section("redis").Key("RedisDB").String()
	RedisAddr = f.Section("redis").Key("RedisAddr").String()
	RedisPw = f.Section("redis").Key("RedisPw").String()
	RedisDbName = f.Section("redis").Key("RedisDbName").String()
}

// loadPath 加载 .ini 文件的 path 项
func loadPath(f *ini.File) {
	Host = f.Section("path").Key("Host").String()
	ImgPath = f.Section("path").Key("ImgPath").String()
	VideoPath = f.Section("path").Key("VideoPath").String()
}

// loadCommon 加载 .ini 文件的 common 项
func loadCommon(f *ini.File) {
	SecretKey = f.Section("common").Key("SecretKey").String()
}
