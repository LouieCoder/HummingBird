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

//func SaveExcelFromFile(filename string) {
//	//创建表结构
//	dao.CreateTable(model.Job{})
//
//	jobs := readexcl(filename)
//	for _, job := range jobs {
//		dao.AddRecord(&job)
//	}
//}

func SaveExcelFormDir(dir string) {
	//创建表结构
	//dao.CreateTable(model.Job{})
	dao.CreateTable(model.Plasticine{})
	//提取目录中的文件名
	excelFilesName := traverseDir(dir)

	for _, ef := range excelFilesName {

		//go func(filename string) {
		//	plasticines := readexcl(dir + filename)
		//
		//	log.Printf("已读取文件%s", filename)
		//
		//	for _, plasticine := range plasticines {
		//		log.Printf("写入来自%s 的记录 %s", filename, plasticine.A)
		//		go dao.AddRecord(&plasticine)
		//	}
		//}(ef.Name())
		saveFileFromFile("exl/" + ef.Name())

	}
}

func saveFileFromFile(filename string) {
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

	//jobs := make([]model.Job, 0, 1024)
	plasticines := make([]model.Plasticine, 0, 1024)

	rows, err := f.GetRows("Sheet1")
	//拿到excel表格的每一列的列名
	model.ColumnNameMaping = rows[0]
	//第一行忽略
	rows = rows[1:]

	for _, row := range rows {
		//job := new(model.Job)
		//
		//job.Department_name = row[0]
		//job.Job_title = row[1]
		//job.Subject_category = row[4]
		//job.Area_belong = row[5]
		//if job.Enrollment, err = strconv.ParseInt(row[7], 10, 64); err != nil {
		//	log.Fatalln(err)
		//}
		//job.Enrollment_target = row[8]
		//job.Education = row[9]
		//job.Other_requirements = row[10]
		//job.Gender_requirement = row[11]
		//job.Polic_requirement = row[12]
		//job.Minimum_service_requirements = row[13]
		//job.Basic_work_experience_requirements = row[14]
		//job.Fitness_test = row[15]
		//job.Department_attribute1 = row[17]
		//job.Department_grade = row[19]
		//job.Phone = row[22]
		//
		//jobs = append(jobs, *job)

		var plasticine model.Plasticine
		val := reflect.ValueOf(&plasticine)
		for i, v := range row {
			val.Elem().FieldByName(string(65 + i)).SetString(v)
		}
		plasticines = append(plasticines, plasticine)
	}

	//return jobs
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
