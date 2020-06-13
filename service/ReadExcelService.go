package service

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"log"
	"os"
	"HummingBird/model"
	"strconv"
	"time"
)

//读传入的excel文件并返回数组切片
func readexcl(filename string) []model.Job {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	jobs := make([]model.Job, 0, 1024)

	rows, err := f.GetRows("Sheet1")
	rows = rows[1:]
	for _, row := range rows {
		job := new(model.Job)
		job.CreatedAt = time.Now()
		job.UpdatedAt = time.Now()

		job.Department_name = row[0]
		job.Job_title = row[1]
		job.Subject_category = row[4]
		job.Area_belong = row[5]
		if job.Enrollment, err = strconv.ParseInt(row[7], 10, 64); err != nil {
			log.Fatalln(err)
		}
		job.Enrollment_target = row[8]
		job.Education = row[9]
		job.Other_requirements = row[10]
		job.Gender_requirement = row[11]
		job.Polic_requirement = row[12]
		job.Minimum_service_requirements = row[13]
		job.Basic_work_experience_requirements = row[14]
		job.Fitness_test = row[15]
		job.Department_attribute1 = row[17]
		job.Department_grade = row[19]
		job.Phone = row[22]

		jobs = append(jobs, *job)
	}

	return jobs
}


//读取dir目录下的文件并返回数组切片
func traverseDir(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic("readDir wrong!")
	}

	files_return := make([]os.FileInfo,0)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			files_return = append(files_return,file)
		}
	}

	return files_return
}
