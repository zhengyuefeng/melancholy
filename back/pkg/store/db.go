package store

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var db *gorm.DB

// Setup 初始化数据库连接
func Setup() {

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&time_zone=%s",
		conf.C.Database.Username,
		conf.C.Database.Password,
		conf.C.Database.Host,
		conf.C.Database.Port,
		conf.C.Database.Name,
		url.QueryEscape("'Asia/Shanghai'"))


	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("Open DB connect failed %s", err.Error())
	}
	log.Info("Init DB successfully")
}

//GetConn 获取db connect
func GetConn() *gorm.DB {
	return db
}
