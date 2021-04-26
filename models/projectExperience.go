package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ProjectExperience struct {
	ProExpID    int       `orm:"pk;auto;column(pro_exp_id)" json:"pro_exp_id"` //项目经历ID
	Name        string    `json:"name"`                                        //项目名称
	Role        string    `json:"角色"`                                          //项目中承担的角色
	StartTime   time.Time `orm:"type(datetime)" json:"start_time"`             //项目开始时间
	EndTime     time.Time `orm:"type(datetime)" json:"end_time"`               //项目结束时间
	Description string    `orm:"type(text)" json:"description"`                //项目描述

	Resume *Resume `orm:"rel(fk)"` //简历与教育经历的一对多关系
}

func (e *ProjectExperience) TableName() string {
	return TNProjectExperience()
}

// GetProjectExperiencesByResumeID 根据简历ID获取所有项目经历
func GetProjectExperiencesByResumeID(resumeID int) ([]*ProjectExperience, error) {
	o := orm.NewOrm()
	var projectExperiences []*ProjectExperience
	_, err := o.QueryTable(TNProjectExperience()).Filter("resume_id", resumeID).All(&projectExperiences)
	return projectExperiences, err
}
