package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init() {
	DSN := "root:mysql@(localhost)/test?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open("mysql", DSN)
	if err != nil {
		log.Fatalln(err)
		return
	}
	//DB.DB().SetMaxIdleConns(5)
	//DB.DB().SetMaxOpenConns(10)
}

func GetConn() *gorm.DB {
	return DB
}
/*
传入结构体,按照该结构创建表
 */
func CreateTable(i interface{}) {
	if !DB.HasTable(i) {
		DB.CreateTable(i)
	}
}
/*
创建一行记录
 */
func AddRecord(i interface{}){
	DB.Create(i)
}
func DestoryConn()  {
	DB.Close()
}