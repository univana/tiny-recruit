package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Job struct {
	JobID             int       `orm:"pk;auto;column(job_id)" json:"job_id"`           //职位ID
	Title             string    `json:"title"`                                         //职位标题
	Description       string    `orm:"type(text)" json:"description"`                  //职位描述
	Location          string    `json:"location"`                                      //职位工作地址
	MinMonthlySalary  int       `json:"min_monthly_salary"`                            //最低月薪
	MaxMonthlySalary  int       `json:"max_monthly_salary"`                            //最高月薪
	PayTimes          int       `json:"pay_times"`                                     //每年发薪次数
	RequireEducation  string    `json:"require_education"`                             //学历要求
	RequireExperience string    `json:"require_experience"`                            //经验要求
	Type              string    `json:"type"`                                          //职位类型
	Nature            string    `json:"nature"`                                        //职位性质
	Status            int       `orm:"default(0)" json:"status"`                       //职位状态
	CreateTime        time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //职位创建时间
	ModifyTime        time.Time `orm:"type(datetime);null" json:"modify_time"`         //最后修改时间

	Enterprise *Enterprise `orm:"rel(fk)" json:"enterprise"` //企业和职位的一对多关系

}

func (m *Job) TableName() string {
	return TNJob()
}

// GetAllJobs 获取所有职位信息
func GetAllJobs() ([]Job, error) {
	var jobs []Job
	var err error
	o := orm.NewOrm()
	_, err = o.QueryTable(TNJob()).All(&jobs)

	//同步对应企业数据
	for i := 0; i < len(jobs); i++ {
		_, err = o.LoadRelated(&jobs[i], "Enterprise")
	}

	return jobs, err
}
