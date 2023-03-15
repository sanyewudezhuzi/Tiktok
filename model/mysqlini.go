package model

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanyewudezhuzi/tiktok/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// Mysqlini 连接 mysql 数据库
func Mysqlini() {
	dsn := strings.Join([]string{conf.DbUser, ":", conf.DbPassword, "@tcp(", conf.DbHost, ":", conf.DbPort, ")/", conf.DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，mysql5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，mysql5.7 之前的数据库不支持
		DontSupportRenameColumn:   true,  // 用`change`重命名列，mysql8 之前的数据库不支持
		SkipInitializeWithVersion: false, // 根据当前 mysql 版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名单数化
		},
	})
	if err != nil {
		panic("Failed to connect mysql.")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间

	DB = db
}

// AutomigrateMySQL 数据库表自动迁移
func AutomigrateMySQL() {
	err := DB.AutoMigrate(
		&Comment{},
		&Favorite{},
		&Follow{},
		&User{},
		&Video{},
	)
	if err != nil {
		panic("Failed to automigrate.")
	}
}
