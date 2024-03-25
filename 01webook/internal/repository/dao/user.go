package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserNotFound       = gorm.ErrRecordNotFound
	ErrUserDuplicateEmail = errors.New("userDao: 邮箱冲突")
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (d UserDAO) FindByEmail(
	ctx *gin.Context, email string) (DaoisUser, error) {
	var u DaoisUser
	err := d.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
}

func (d UserDAO) Insert(ctx context.Context, u DaoisUser) error {
	now := time.Now().UnixMilli() //毫秒数
	u.CreatTime = now
	u.UpdateTime = now

	err := d.db.WithContext(ctx).Create(&u).Error //gorm
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			fmt.Println("邮箱冲突")
			return ErrUserDuplicateEmail
		}
	}
	return err
	//panic("zuo ye wu ")
}

// DaoisUser 直接对应 数据库表结构
type DaoisUser struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string //`gorm:"password"`

	CreatTime  int64 //`gorm:"creat_time"`
	UpdateTime int64 //`gorm:"update_time"`
}
