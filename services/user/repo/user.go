package repo

import (
	"errors"
	"log"
	"time"
)

type User struct {
	Id         int64     `xorm:" pk autoincr INT(11)"`
	Phone      int64     `xorm:"not null default 0 unique(phone) comment('手机号码') bigint"`
	Password   string    `xorm:"not null default '' comment('密码') VARCHAR(255)"`
	UpdateTime time.Time `xorm:"updated_at not null  DEFAULT CURRENT_TIMESTAMP  timestamp"`
	CreateTime time.Time `xorm:"created_at not null DEFAULT CURRENT_TIMESTAMP timestamp"`
	Version    uint32    `xorm:"version default 1 int"`
}

func init() {
	user := new(User)
	if isExist, _ := x.IsTableExist(user); !isExist {
		if err := x.Sync2(user); err != nil {
			log.Fatal("sync tables err:%v", err)
		}
	}
}

func (u *User) Create(user User) (bool, error) {
	affected, err := x.Insert(user)
	if err != nil {
		return false, err
	}
	if affected == 0 {
		return false, errors.New("insert user fail")
	}
	return true, nil
}

func (u *User) FindByPhone(phone int64) (User, error) {
	var user User
	rel, err := x.Where("phone = ?", phone).Get(&user)
	if !rel || err != nil {
		log.Printf("user find by phone err:%v", err)
		return user, errors.New("not found")
	}
	return user, nil
}
