package service

import (
	"HummingBird/dao"
	"HummingBird/model"
)

func saveExcel(filename string)  {
	createTable(model.Job{})

	jobs := readexcl(filename)
	for _, job := range jobs {
		dao.AddRecord(&job)
	}
}

//根据给定的结构体实例创建表
func createTable(i interface{}){
	//传递初始化为空的结构体实例 dao.CreateTable(model.Job{})
	dao.CreateTable(i)
}