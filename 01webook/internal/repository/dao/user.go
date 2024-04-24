package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserNotFound       = gorm.ErrRecordNotFound
	ErrUserDuplicateEmail = errors.New("userDao: 邮箱冲突")
)

type UserDAO interface {
	Insert(ctx context.Context, u DaoisUser) error
	FindByEmail(ctx context.Context, email string) (DaoisUser, error)
	UpdateById(ctx context.Context, entity DaoisUser) error
	FindById(ctx context.Context, uid int64) (DaoisUser, error)
	FindByPhone(ctx context.Context, phone string) (DaoisUser, error)
}

type userDAO struct {
	db *gorm.DB
}

func (dao userDAO) UpdateById(ctx context.Context, entity DaoisUser) error {
	// 这种写法依赖于 GORM 的零值和主键更新特性
	// Update 非零值 WHERE id = ?
	//return dao.db.WithContext(ctx).Updates(&entity).Error
	return dao.db.WithContext(ctx).Model(&entity).Where("id = ?", entity.Id).
		Updates(map[string]any{
			"utime":    time.Now().UnixMilli(),
			"nickname": entity.Nickname,
			"birthday": entity.Birthday,
			"about_me": entity.AboutMe,
		}).Error
}

func (dao userDAO) FindById(ctx context.Context, uid int64) (DaoisUser, error) {
	var res DaoisUser
	err := dao.db.WithContext(ctx).Where("id = ?", uid).First(&res).Error
	return res, err
}

func (dao userDAO) FindByPhone(ctx context.Context, phone string) (DaoisUser, error) {
	var res DaoisUser
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&res).Error
	return res, err
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &userDAO{db: db}
}

func (dao userDAO) FindByEmail(
	ctx context.Context, email string) (DaoisUser, error) {
	var u DaoisUser
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
}

func (dao userDAO) Insert(ctx context.Context, u DaoisUser) error {
	now := time.Now().UnixMilli() //毫秒数
	u.Ctime = now
	u.Utime = now

	err := dao.db.WithContext(ctx).Create(&u).Error //gorm
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

//type DaoisUser2 struct {
//	Id       int64  `gorm:"primaryKey,autoIncrement"`
//	Email    string `gorm:"unique"`
//	Password string //`gorm:"password"`
//
//	CreatTime  int64 //`gorm:"creat_time"`
//	UpdateTime int64 //`gorm:"update_time"`
//}

// DaoisUser 直接对应 数据库表结构
type DaoisUser struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 代表这是一个可以为 NULL 的列
	//Email    *string
	Email    sql.NullString `gorm:"unique"`
	Password string

	Nickname string `gorm:"type=varchar(128)"`
	// YYYY-MM-DD
	Birthday int64
	AboutMe  string `gorm:"type=varchar(4096)"`

	// 代表这是一个可以为 NULL 的列
	Phone sql.NullString `gorm:"unique"`

	// 时区，UTC 0 的毫秒数
	// 创建时间
	Ctime int64
	// 更新时间
	Utime int64

	// json 存储
	//Addr string
}
