package ioc

import (
	"github.com/spf13/viper"
	"go20240218/02webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	type Config struct {
		DSN string `yaml:"dsn"`

		// 有些人的做法
		// localhost:13316
		//Addr string
		//// localhost
		//Domain string
		//// 13316
		//Port string
		//Protocol string
		//// root
		//Username string
		//// root
		//Password string
		//// webook
		//DBName string
	}
	var cfg = Config{
		DSN: "root:root@tcp(localhost:13316)/webook_default",
	}
	// 看起来，remote 不支持 key 的切割
	err := viper.UnmarshalKey("db", &cfg)
	//dsn := viper.GetString("db.mysql")
	//println(dsn)
	//if err != nil {
	//	panic(err)
	//}
	db, err := gorm.Open(mysql.Open(cfg.DSN))
	if err != nil {
		// 我只会在初始化过程中 panic
		// panic 相当于整个 goroutine 结束
		// 一旦初始化过程出错，应用就不要启动了
		panic(err)
	}

	//dao.NewUserDAOV1(func() *gorm.DB {
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//oldDB := db
	//db, err = gorm.Open(mysql.Open())
	//pt := unsafe.Pointer(&db)
	//atomic.StorePointer(&pt, unsafe.Pointer(&db))
	//oldDB.Close()
	//})
	// 要用原子操作
	//return db
	//})

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
