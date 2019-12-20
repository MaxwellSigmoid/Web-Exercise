package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"
)

// ORM gorm引擎的实例
var DBInstance *gorm.DB

func Init() {
	var err error

	//连接数据库，为方便教学，数据库使用sqlite
	db, err := gorm.Open("sqlite3", "./database/test.db")
	if err != nil {
		panic("连接数据库失败")
	}

	DBInstance = db
	// 打印sql详情
	db.LogMode(true)
}
