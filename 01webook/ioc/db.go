package ioc

import (
	"fmt"
	"github.com/spf13/viper"
	"go20240218/01webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//dsn := viper.GetString("db.mysql.dsn")
	//fmt.Println(dsn, " tmh写的")

	type Config struct {
		DSN string `yaml:"dsn"`
	}
	var c Config
	err := viper.UnmarshalKey("db.mysql", &c)
	if err != nil {
		panic(fmt.Errorf("初始化配置失败"))
	}

	//初始化数据库
	//db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	//db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	db, err1 := gorm.Open(mysql.Open(c.DSN))
	if err1 != nil {
		fmt.Println("tmh: 数据库 连接 失败")
		panic(err)
	}

	err2 := dao.InitTable(db)
	if err2 != nil {
		fmt.Println("tmh: 数据库 建表 失败")
		panic(err2)
	}
	return db
}
