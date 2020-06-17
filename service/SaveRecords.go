package service

import (
	"HummingBird/dao"
	"HummingBird/model"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func SaveExcelFormDir(dir string) {
	//创建表结构
	//dao.CreateTable(model.Job{})
	dao.CreateTable(model.Plasticine{})
	//提取目录中的文件名
	excelFilesName := traverseDir(dir)

	for _, ef := range excelFilesName {
		go saveFileFromExcel("exl/" + ef.Name())
	}
}

func saveFileFromExcel(filename string) {
	plasticines := readexcl(filename)

	log.Printf("已读取文件%s", filename)
	//错误写法 plasticine在整个循环过程中只存在一个副本
	//for _, plasticine := range plasticines {
	//	log.Printf("写入来自%s 的记录 %s", filename, plasticine.A)
	//	go dao.AddRecord(&plasticine)
	//}
	for i := 0; i < len(plasticines); i++ {
		log.Printf("写入来自%s 的记录 %s", filename, plasticines[i].A)
		go dao.AddRecord(&plasticines[i])
	}
}

//读传入的excel文件并返回数组切片
func readexcl(filename string) []model.Plasticine {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	plasticines := make([]model.Plasticine, 0, 1024)

	rows, err := f.GetRows("Sheet1")
	//拿到excel表格的每一列的列名
	model.ColumnNameMaping = rows[0]
	//第一行忽略
	rows = rows[1:]

	for _, row := range rows {
		var plasticine model.Plasticine
		plasticine.Id = GenerateID()
		val := reflect.ValueOf(&plasticine)
		for i, v := range row {
			val.Elem().FieldByName(string(65 + i)).SetString(v)
		}
		plasticines = append(plasticines, plasticine)
	}

	return plasticines
}

//读取dir目录下的文件并返回数组切片
func traverseDir(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic("readDir wrong!")
	}

	files_return := make([]os.FileInfo, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			files_return = append(files_return, file)
		}
	}

	return files_return
}
