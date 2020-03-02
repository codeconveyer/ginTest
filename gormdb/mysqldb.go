package gormdb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/wubin/gin-test/model"
)

type db struct{
	*gorm.DB
}
var DB db
func init() {
	sqlDb, err := gorm.Open("mysql", "root:prisma@tcp(127.0.0.1:3306)/gin-test")
	if err != nil {
		panic(err)
	}
	sqlDb.LogMode(true)
	sqlDb.DB().SetMaxIdleConns(5)  //最多保持5个空闲连接
	sqlDb.DB().SetMaxOpenConns(100) //最大连接数不超过100
	sqlDb.DB().SetConnMaxLifetime(60) //1分钟一次心跳 WAIT_TIMEOUT 建议（5分钟~30分钟之间）
	DB = db{sqlDb}
	//defer sqlDb.Close()
	sqlDb.AutoMigrate(&model.User{})
}
