package ioc

import (
	"fmt"
	"go20240218/01webook/config"
	"go20240218/01webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//初始化数据库
	//db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
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
