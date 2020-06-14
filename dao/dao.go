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
	DB.DB().SetMaxOpenConns(20)
}

func GetConn() *gorm.DB {
	return DB
}
/*
传入结构体,按照该结构创建表
 */
func CreateTable(model interface{}) {
	if !DB.HasTable(model) {
		DB.CreateTable(model)
	}
}
/*
根据传入参数record(指针类型)创建一行记录
 */
func AddRecord(record interface{}){
	DB.Create(record)
}
func DestoryConn()  {
	DB.Close()
}