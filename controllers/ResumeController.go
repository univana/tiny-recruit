package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	uuid "github.com/iris-contrib/go.uuid"
	"myApp/models"
	"path"
	"strconv"
	"strings"
	"time"
)

type ResumeController struct {
	BaseController
}

func (c *ResumeController) GetResumeByMemberID() {

	memberID, _ := strconv.Atoi(c.GetString("id"))
	resume, err := models.GetResumeByMemberID(memberID)
	if err == orm.ErrNoRows {
		c.JsonResult(0, "ok", models.Resume{ResumeID: 0})
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

	_, info, err := c.GetFile("file")
	if info != nil {
		//获取文件后缀
		//保存图片
		ext := strings.ToLower(path.Ext(info.Filename))
		v4, _ := uuid.NewV4()

		fileName := fmt.Sprintf("%s%s", v4, ext)
		filePath := fmt.Sprintf("static/images/photos/%s", fileName)
		err = c.SaveToFile("file", filePath)
		if err != nil {
			logs.Error("Error ResumeController Edit: ", err)
		}

		resumeID, _ := strconv.Atoi(c.GetString("resume_id"))
		gender, _ := strconv.Atoi(c.GetString("gender"))
		b := c.GetString("birthday")
		y, _ := strconv.Atoi(b[0:4])
		m, _ := strconv.Atoi(b[5:7])
		d, _ := strconv.Atoi(b[8:10])
		birthday := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)

		resume := models.Resume{
			ResumeID:   resumeID,
			Name:       c.GetString("name"),
			Gender:     gender,
			Birthday:   birthday,
			Advantage:  c.GetString("advantage"),
			Tel:        c.GetString("tel"),
			Email:      c.GetString("email"),
			HopeSalary: c.GetString("hope_salary"),
			HopeJob:    c.GetString("hope_job"),
			City:       c.GetString("city"),
			Photo:      "/" + filePath,
			ModifyTime: time.Now(),
			Member:     &models.Member{MemberId: c.Member.MemberId},
			SkillTags:  c.GetString("skill_tags"),
		}
		err := resume.InsertOrUpdate("skill_tags", "photo", "name", "gender", "birthday", "advantage", "tel", "email", "hope_salary", "hope_job", "city")
		if err != nil {
			logs.Error("Error ResumeController EditResume: ", err)
			c.JsonResult(1, "修改简历失败！")
		} else {
			c.JsonResult(0, "ok")
		}
	} else {
		resumeID, _ := strconv.Atoi(c.GetString("resume_id"))
		gender, _ := strconv.Atoi(c.GetString("gender"))
		b := c.GetString("birthday")
		y, _ := strconv.Atoi(b[0:4])
		m, _ := strconv.Atoi(b[5:7])
		d, _ := strconv.Atoi(b[8:10])
		birthday := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)

		resume := models.Resume{
			ResumeID:   resumeID,
			Name:       c.GetString("name"),
			Gender:     gender,
			Birthday:   birthday,
			Advantage:  c.GetString("advantage"),
			Tel:        c.GetString("tel"),
			Email:      c.GetString("email"),
			HopeSalary: c.GetString("hope_salary"),
			HopeJob:    c.GetString("hope_job"),
			City:       c.GetString("city"),
			ModifyTime: time.Now(),
			Member:     &models.Member{MemberId: c.Member.MemberId},
			SkillTags:  c.GetString("skill_tags"),
		}
		err := resume.InsertOrUpdate("skill_tags", "name", "gender", "birthday", "advantage", "tel", "email", "hope_salary", "hope_job", "city")
		if err != nil {
			logs.Error("Error ResumeController EditResume: ", err)
			c.JsonResult(1, "修改简历失败！")
		} else {
			c.JsonResult(0, "ok")
		}
	}
}

// AddEdu 添加教育经历
func (c *ResumeController) AddEdu() {

	startYear, _ := strconv.Atoi(strings.Split(c.GetString("start_year"), " ")[3])
	endYear, _ := strconv.Atoi(strings.Split(c.GetString("end_year"), " ")[3])
	resume, err := models.GetResumeByMemberID(c.Member.MemberId)
	if err != nil {
		logs.Error("Error ResumeController AddEdu: ", err)
	}
	var edu = models.EducationExperience{
		School:     c.GetString("school"),
		Education:  c.GetString("education"),
		Profession: c.GetString("profession"),
		StartYear:  startYear,
		EndYear:    endYear,
		Experience: c.GetString("experience"),
		Resume:     &models.Resume{ResumeID: resume.ResumeID},
	}
	err = edu.InsertOrUpdate()
	if err != nil {
		logs.Error("Error ResumeController AddEdu: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}

// AddProject 添加项目经历
func (c *ResumeController) AddProject() {
	s := strings.Split(c.GetString("start_time"), "-")
	y, _ := strconv.Atoi(s[0])
	m, _ := strconv.Atoi(s[1])
	d, _ := strconv.Atoi(s[2])
	startTime := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)

	e := strings.Split(c.GetString("end_time"), "-")
	y, _ = strconv.Atoi(e[0])
	m, _ = strconv.Atoi(e[1])
	d, _ = strconv.Atoi(e[2])
	endTime := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	resume, err := models.GetResumeByMemberID(c.Member.MemberId)
	if err != nil {
		logs.Error("Error ResumeController AddProject: ", err)
	}

	var pro = models.ProjectExperience{
		ProExpID:    0,
		Name:        c.GetString("name"),
		Role:        c.GetString("role"),
		StartTime:   startTime,
		EndTime:     endTime,
		Description: c.GetString("description"),
		Resume:      &models.Resume{ResumeID: resume.ResumeID},
	}

	err = pro.InsertOrUpdate()
	if err != nil {
		logs.Error("Error ResumeController AddProject: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}

// AddInternship 添加实习经历
func (c *ResumeController) AddInternship() {
	s := strings.Split(c.GetString("start_time"), "-")
	y, _ := strconv.Atoi(s[0])
	m, _ := strconv.Atoi(s[1])
	d, _ := strconv.Atoi(s[2])
	startTime := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)

	e := strings.Split(c.GetString("end_time"), "-")
	y, _ = strconv.Atoi(e[0])
	m, _ = strconv.Atoi(e[1])
	d, _ = strconv.Atoi(e[2])
	endTime := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	resume, err := models.GetResumeByMemberID(c.Member.MemberId)
	if err != nil {
		logs.Error("Error ResumeController AddInternship: ", err)
	}

	var internship = models.InternshipExperience{
		IntExpID:    0,
		CompanyName: c.GetString("company_name"),
		Department:  c.GetString("department"),
		Position:    c.GetString("position"),
		StartTime:   startTime,
		EndTime:     endTime,
		WorkContent: c.GetString("work_content"),
		Resume:      &models.Resume{ResumeID: resume.ResumeID},
	}

	err = internship.InsertOrUpdate()
	if err != nil {
		logs.Error("Error ResumeController AddInternship: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}

// DeleteExperience 删除经历
func (c *ResumeController) DeleteExperience() {
	id, _ := strconv.Atoi(c.GetString("id"))
	expType := c.GetString("type")
	var err error
	if expType == "1" {
		var eduExp = models.EducationExperience{
			EduExpID: id,
			Deleted:  1,
		}
		err = eduExp.InsertOrUpdate("deleted")

	} else if expType == "2" {
		var proExp = models.ProjectExperience{
			ProExpID: id,
			Deleted:  1,
		}
		err = proExp.InsertOrUpdate("deleted")
	} else {
		var intExp = models.InternshipExperience{
			IntExpID: id,
			Deleted:  1,
		}
		err = intExp.InsertOrUpdate("deleted")

	}
	if err != nil {
		logs.Error("Error ResumeController DeleteExperience: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}

}
