package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicate = errors.New("邮箱冲突")
	ErrUserNotFound  = gorm.ErrRecordNotFound
)

type UserDAO interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	FindById(ctx context.Context, id int64) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
	Insert(ctx context.Context, u User) error
	FindByWechat(ctx context.Context, openID string) (User, error)
}

type DBProvider func() *gorm.DB

type GORMUserDAO struct {
	db *gorm.DB

	p DBProvider
}

func NewUserDAOV1(p DBProvider) UserDAO {
	return &GORMUserDAO{
		p: p,
	}
}

func NewUserDAO(db *gorm.DB) UserDAO {
	res := &GORMUserDAO{
		db: db,
	}
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	db, err := gorm.Open(mysql.Open())
	//	pt := unsafe.Pointer(&res.db)
	//	atomic.StorePointer(&pt, unsafe.Pointer(&db))
	//})
	return res
}

func (dao *GORMUserDAO) FindByWechat(ctx context.Context, openID string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("wechat_open_id = ?", openID).First(&u).Error
	//err := dao.p().WithContext(ctx).Where("wechat_open_id = ?", openID).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?", email).Error
	return u, err
}

func (dao *GORMUserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?", email).Error
	return u, err
}

func (dao *GORMUserDAO) FindByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?", email).Error
	return u, err
}

func (dao *GORMUserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("`id` = ?", id).First(&u).Error
	return u, err
}

func (dao *GORMUserDAO) Insert(ctx context.Context, u User) error {
	// 存毫秒数
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突 or 手机号码冲突
			return ErrUserDuplicate
		}
	}
	return err
}

// User 直接对应数据库表结构
// 有些人叫做 entity，有些人叫做 model，有些人叫做 PO(persistent object)
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 全部用户唯一
	Email    sql.NullString `gorm:"unique"`
	Password string

	// 唯一索引允许有多个空值
	// 但是不能有多个 ""
	Phone sql.NullString `gorm:"unique"`
	// 最大问题就是，你要解引用
	// 你要判空
	//Phone *string

	// 往这面加

	// 索引的最左匹配原则：
	// 假如索引在 <A, B, C> 建好了
	// A, AB, ABC 都能用
	// WHERE A =?
	// WHERE A = ? AND B =?    WHERE B = ? AND A =?
	// WHERE A = ? AND B = ? AND C = ?  ABC 的顺序随便换
	// WHERE 里面带了 ABC，可以用
	// WHERE 里面，没有 A，就不能用

	// 如果要创建联合索引，<unionid, openid>，用 openid 查询的时候不会走索引
	// <openid, unionid> 用 unionid 查询的时候，不会走索引
	// 微信的字段
	WechatUnionID sql.NullString
	WechatOpenID  sql.NullString `gorm:"unique"`

	// 创建时间，毫秒数
	Ctime int64
	// 更新时间，毫秒数
	Utime int64
}
