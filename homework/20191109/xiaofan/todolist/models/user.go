package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"todolist/utils"
)

// user结构
type User struct {
	Id         int        `json:"id"`
	Name       string     `orm:"type(varchar);size(32);default();" json:"name"`
	Password   string     `orm:"type(varchar);size(1024);default();" json:"-"`
	Birthday   *time.Time `orm:"type(date);null;" json:"birthday"`
	Sex        bool       `orm:"default(false);" json:"sex"`
	Tel        string     `orm:"type(varchar);size(16);default();" json:"tel"`
	Addr       string     `orm:"type(varchar);size(512);default();" json:"addr"`
	Desc       string     `orm:"type(text);default();" json:"desc"`
	IsSuper    bool       `orm:"default(false);" json:"is_super"`
	CreateTime *time.Time `orm:"type(datetime);auto_now_add;" json:"create_time"`
}

func (u *User) SetPassword(plain string) {
	u.Password = utils.Md5String(plain)
}

func (u *User) ValidatePassword(password string) bool {
	return utils.Md5String(password) == u.Password
}

/*
func AddUser(user *User) error {
	o := orm.NewOrm()
	_, err := o.Insert(user)
	return err
}
*/

// 获取用户
func GetUser(user *User, cols ...string) error {
	o := orm.NewOrm()
	return o.Read(user, cols...)
}

// 初始化数据
func init() {
	orm.RegisterModel(new(User))
}
