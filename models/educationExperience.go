package models

import "github.com/astaxie/beego/orm"

type EducationExperience struct {
	EduExpID   int    `orm:"pk;auto;column(edu_exp_id)" json:"edu_exp_id"` //教育经历ID
	School     string `json:"school"`                                      //学校名称
	Education  string `json:"education"`                                   //学历
	Profession string `json:"profession"`                                  //专业
	StartYear  int    `json:"start_year"`                                  //开始年份
	EndYear    int    `json:"end_year"`                                    //结束年份
	Experience string `orm:"type(text)" json:"experience"`                 //教育经历内容

	Resume *Resume `orm:"rel(fk)" json:"resume"` //简历与教育经历的一对多关系
}

func (e *EducationExperience) TableName() string {
	return TNEducationExperience()
}

// GetEducationExperiencesByResumeID 根据简历ID获取所有教育经历
func GetEducationExperiencesByResumeID(resumeID int) ([]*EducationExperience, error) {
	o := orm.NewOrm()
	var educationExperiences []*EducationExperience
	_, err := o.QueryTable(TNEducationExperience()).Filter("resume_id", resumeID).All(&educationExperiences)
	return educationExperiences, err
}

// InsertOrUpdate 添加或更新教育经历
func (e *EducationExperience) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var edu EducationExperience
	err := o.QueryTable(TNEducationExperience()).Filter("edu_exp_id", e.EduExpID).One(&edu)
	if edu.EduExpID > 0 {
		_, err = o.Update(e, fields...)
	} else {
		_, err = o.Insert(e)
	}
	return err
}
