package dao

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"go-test/webook/internal/constants"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	if u.BirthDay == "" {
		u.BirthDay = time.UnixMilli(time.Now().UnixMilli()).Format("2006-01-02 15:04:05")
	}
	if u.NickName == "" {
		u.NickName = "webook"
	}
	if u.Profile == "" {
		u.Profile = "这是一个新用户"
	}
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == 1062 {
			return constants.ErrDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
}

func (dao *UserDAO) FindByID(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	return u, err
}

func (dao *UserDAO) DeleteByID(ctx context.Context, id int64) error {
	var u User
	err := dao.db.WithContext(ctx).Where("id = ?", id).Delete(&u).Error
	return err
}

func (dao *UserDAO) UpdateByID(ctx context.Context, id int64, u User) error {
	err := dao.db.WithContext(ctx).Model(&u).Where("id = ?", id).Updates(&u).Error
	return err
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	NickName string
	BirthDay string
	Profile  string
	//创建时间，毫秒数
	Ctime int64
	//更新时间，毫秒数
	Utime int64
}
