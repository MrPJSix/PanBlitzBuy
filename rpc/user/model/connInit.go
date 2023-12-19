package model

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"pan-blitz-buy/rpc/user/internal/config"
	"time"
)

// ModelInit 每次启动初始化数据库，便于更改数据库结构
func InitMySQL(c *config.Config) *gorm.DB {
	log.Println("User服务初始化MySQL连接开始...")

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Schema,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("连接MySQL数据库失败！", err)
		panic("连接MySQL数据库失败！")
	}

	err = db.AutoMigrate(
		&User{},
		&Admin{},
	)

	if err != nil {
		log.Println("MySQL数据库迁移出错！", err)
		panic("MySQL数据库迁移出错！")
	}

	log.Println("User服务初始化MySQL连接结束...")

	return db
}

func InitCache(c *config.Config) *cache.Cache {
	codeCache := cache.New(time.Duration(c.CodeCache.ExpireTime)*time.Second, time.Duration(c.CodeCache.CleanTime)*time.Second)
	return codeCache
}
