package gormdb
//
//import (
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/postgres"
//	"github.com/wubin/gin-test/model"
//)
//
//func InitPGSql() {
//	sqldb, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
//	if err != nil {
//		panic(err)
//	}
//	DB = db{sqlDb}
//	//defer sqlDb.Close()
//	sqlDb.AutoMigrate(&model.User{})
//}