package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser     string = "root"
	dbPassword string = "881019"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "testgorm"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

type User struct {
	Id         int    `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(32); not null; default:''"`
	Password   string
	Birthday   time.Time `gorm:"type:date"`
	Sex        bool
	Tel        string `gorm:"column:telephone"`
	Addr       string
	Desciption string `gorm:"type:text"`
}

func (u *User) TableName() string {
	return "user"
}

func main() {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	db.LogMode(true)
	db.AutoMigrate(&User{})

	//查找对象进行更新
	var user User
	if db.First(&user, "name=?", "kk_3").Error == nil {
		db.Delete(&user)
	}

	db.Where("id > ?", 15).Delete(&User{})

	db.Close()
}
