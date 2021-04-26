package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(
		new(Member),
		new(Resume),
		new(EducationExperience),
		new(ProjectExperience),
		new(InternshipExperience),
		new(Enterprise),
	)
}

//TNMembers return-用户表名
func TNMembers() string {
	return "t_member"
}

//TNResume return-用户表名
func TNResume() string {
	return "t_resume"
}

//TNEducationExperience return-教育经历表
func TNEducationExperience() string {
	return "t_education_experience"
}

func TNProjectExperience() string {
	return "t_project_experience"
}

func TNInternshipExperience() string {
	return "t_internship_experience"
}

func TNEnterprise() string {
	return "t_enterprise"
}
