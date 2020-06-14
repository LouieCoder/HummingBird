package model

type Job struct {
	Department_name                    string
	Job_title                          string
	Department_id                      int64
	Job_id                             int64
	Subject_category                   string
	Area_belong                        string
	Exam_area                          string
	Enrollment                         int64
	Enrollment_target                  string
	Education                          string
	Other_requirements                 string `gorm:"type:varchar(500)"`
	Gender_requirement                 string
	Polic_requirement                  string
	Minimum_service_requirements       string
	Basic_work_experience_requirements string
	Fitness_test                       string
	Job_categary                       string
	Department_categary1               string
	Department_attribute1              string
	Department_grade                   string
	Department_attribute2              string
	Department_categary2               string
	Phone                              string
	Tips                               string
}

/*
更改gorm的默认表名为job
*/
func (j Job) TableName() string {
	return "jobs"
}
