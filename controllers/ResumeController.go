package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
	"strconv"
	"time"
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

func (c *ResumeController) EditResume() {
	c.TplName = "navigation/userCenter.html"
	resumeID, _ := strconv.Atoi(c.GetString("resume_id"))
	gender, _ := strconv.Atoi(c.GetString("gender"))
	b := c.GetString("birthday")
	y, _ := strconv.Atoi(b[0:4])
	m, _ := strconv.Atoi(b[5:7])
	d, _ := strconv.Atoi(b[8:10])
	birthday := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)

	resume := models.Resume{
		ResumeID: resumeID,
		Name:     c.GetString("name"),
		Gender:   gender,
		Birthday: birthday,
	}
	err := resume.InsertOrUpdate("name", "gender", "birthday")
	if err != nil {
		logs.Error("Error ResumeController EditResume: ", err)
		c.JsonResult(1, "修改简历失败！")
	} else {
		c.JsonResult(0, "ok")
	}

}
