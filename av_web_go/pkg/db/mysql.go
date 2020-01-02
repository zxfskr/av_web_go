package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// mysql驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB mysql连接池
var DB *gorm.DB

// ModuleInit 初始化mysql数据库连接
func ModuleInit() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", "root", "123456", "0.0.0.0", 3306, "av_web", "10s")
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// defer DB.Close()
	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(18)
	DB.LogMode(false)
	//不要在表名后默认加s
	DB.SingularTable(true)
}
