package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
)

type ResumeController struct {
	BaseController
}

func (c *ResumeController) GetResumeByMemberID() {
	c.TplName = "navigation/userCenter.html"

	memberID := c.Member.MemberId
	resume, err := models.GetResumeByMemberID(memberID)
	if err != nil {
		logs.Error("Error get resume:", err)
	}
	//查找简历对应的教育经历
	educationExperiences, err := models.GetEducationExperiencesByResumeID(resume.ResumeID)
	if err != nil {
		logs.Error("Error get education experiences: ", err)
	}
	resume.EducationExperiences = educationExperiences

	//查找简历对应的项目经历
	projectExperiences, err := models.GetProjectExperiencesByResumeID(resume.ResumeID)
	if err != nil {
		logs.Error("Error get project experiences: ", err)
	}
	resume.ProjectExperiences = projectExperiences

	//查找简历对应的实习经历
	internshipExperiences, err := models.GetInternshipExperiencesByResumeID(resume.ResumeID)
	if err != nil {
		logs.Error("Error get internship experiences: ", err)
	}
	resume.InternshipExperiences = internshipExperiences

	c.JsonResult(0, "ok", resume)
}
