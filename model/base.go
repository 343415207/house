package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	conn *gorm.DB
)

func DB() *gorm.DB {
	if conn != nil && conn.DB().Ping() == nil {
		return conn
	}

	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/house?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("连接数据库失败:%v", err))
	}

	conn = db
	db.LogMode(true)
	conn.DB().SetConnMaxLifetime(10)
	conn.DB().SetMaxIdleConns(2)
	conn.DB().SetMaxOpenConns(2)

	return conn
}

//GetInfo 获取信息
func GetInfo(source interface{}, where map[string]interface{}) error {
	return DB().Find(source, where).Order(" id desc").Error
}

func SaveOrUpdate(source interface{}, values map[string]interface{}, where map[string]interface{}) error {
	count := 0

	DB().Model(source).Count(&count).Where(where)

	var e error

	if count > 0 {
		e = DB().Model(source).Updates(values).Where(where).Error
	} else {
		e = DB().Create(source).Error
	}

	if e != nil {
		return e
	}

	if e = GetInfo(source, where); e != nil {
		return e
	}

	return nil

}
