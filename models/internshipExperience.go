package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type InternshipExperience struct {
	IntExpID    int       `orm:"pk;auto;column(int_exp_id)" json:"int_exp_id"` //实习经历ID
	CompanyName string    `json:"company_name"`                                //实习公司名称
	Department  string    `json:"department"`                                  //实习所在部门
	Position    string    `json:"position"`                                    //实习担任职位
	StartTime   time.Time `orm:"type(datetime)" json:"start_time"`             //实习开始时间                                 //开始年份
	EndTime     time.Time `orm:"type(datetime)" json:"end_time"`               //实习结束时间                        //结束年份
	WorkContent string    `orm:"type(text)" json:"work_content"`               //实习经历内容
	Deleted     int       `orm:"default(0)" json:"deleted"`                    //删除标记
	Resume      *Resume   `orm:"rel(fk)"`                                      //简历与实习经历的一对多关系
}

func (e *InternshipExperience) TableName() string {
	return TNInternshipExperience()
}

// GetInternshipExperiencesByResumeID 根据简历ID获取所有实习经历
func GetInternshipExperiencesByResumeID(resumeID int) ([]*InternshipExperience, error) {
	o := orm.NewOrm()
	var internshipExperiences []*InternshipExperience
	_, err := o.QueryTable(TNInternshipExperience()).Filter("resume_id", resumeID).Filter("deleted", 0).All(&internshipExperiences)
	return internshipExperiences, err
}

// InsertOrUpdate 添加或更新实习经历
func (e *InternshipExperience) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var internship InternshipExperience
	err := o.QueryTable(TNInternshipExperience()).Filter("int_exp_id", e.IntExpID).One(&internship)
	if internship.IntExpID > 0 {
		_, err = o.Update(e, fields...)
	} else {
		_, err = o.Insert(e)
	}
	return err
}
