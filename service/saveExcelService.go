package service

import (
	"HummingBird/dao"
	"HummingBird/model"
)

func saveExcel(filename string)  {
	//创建表结构
	dao.CreateTable(model.Job{})

	jobs := readexcl(filename)
	for _, job := range jobs {
		dao.AddRecord(&job)
	}
}
