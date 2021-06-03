package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"myApp/common"
	"myApp/models"
	"strconv"
	"time"
)

type JobController struct {
	BaseController
}

func (c *JobController) ShowJob() {
	jobID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	//获取职位信息
	job := models.GetJobByID(jobID)
	c.Data["Job"] = job

	//确定职位是否被投递
	c.Data["Delivered"] = models.IsDelivered(job.JobID, c.Member.MemberId)
	c.Data["Collected"] = models.IsCollected(job.JobID, c.Member.MemberId)
	c.TplName = "job/job-detail.html"
}

func (c *JobController) GetDeliverance() {
	jobID, _ := strconv.Atoi(c.GetString("job_id"))
	job := models.GetJobByID(jobID)
	//获取该职位的所有投递信息
	type Deliver struct {
		DeliveranceID int           `json:"deliverance_id"`
		DeliverTime   time.Time     `json:"deliver_time"` //投递时间
		ModifyTime    time.Time     `json:"modify_time"`  //投递修改时间
		Status        string        `json:"status"`       //投递状态
		Resume        models.Resume `json:"resume"`       //投递对应的简历
		Match         float64       `json:"match"`        //简历-职位匹配度
	}
	var deliverList []Deliver
	//获取简历
	resumes, err := models.GetAllResumesByJobID(jobID)
	if err != nil {
		logs.Error("Error JobController get resumes: ", err)
	}

	for _, resume := range resumes {
		//获取投递时间、投递修改时间和投递状态
		resume.LoadExperiences()

		//匹配度计算
		matchingDegree := common.GetMatchingDegree(resume, job, 0.3, 0.2, 0.3, 0.2)

		memberID := resume.Member.MemberId
		var deliverance models.Deliverance
		o := orm.NewOrm()
		err := o.QueryTable(models.TNDeliverance()).Filter("member_id", memberID).Filter("job_id", jobID).Exclude("status", "已撤销").One(&deliverance)
		if err != nil {
			logs.Error("Error JobController query t_deliverance: ", err)
			c.JsonResult(1, "获取投递数据错误")
		}
		deliverList = append(deliverList, Deliver{DeliveranceID: deliverance.DeliveranceID, DeliverTime: deliverance.DeliverTime, ModifyTime: deliverance.ModifyTime, Status: deliverance.Status, Resume: resume, Match: matchingDegree})
	}
	c.JsonResult(0, "ok", deliverList)
}

// Filter 职位过滤器后端逻辑
func (c *JobController) Filter() {
	searchContent := c.GetString("search_content")
	city := c.GetString("city")
	requireExp := c.GetString("require_experience")
	requireEdu := c.GetString("require_education")
	scale := c.GetString("scale")
	financingStage := c.GetString("financing_stage")
	enterpriseType := c.GetString("type")
	jobs, err := models.FilterJobs(searchContent, city, requireExp, requireEdu, financingStage, scale, enterpriseType)
	if err != nil {
		logs.Error("Error JobController Filter: ", err)
		c.JsonResult(1, "职位过滤错误！")
	} else {
		if len(jobs) == 0 {
			//如果没有职位
			c.JsonResult(1, "没有满足筛选条件的职位！")
		}
		c.JsonResult(0, "ok", jobs)
	}
}

func (c *JobController) NewJob() {
	var cstZone = time.FixedZone("CST", 8*3600)
	minS, _ := strconv.Atoi(c.GetString("min_monthly_salary"))
	maxS, _ := strconv.Atoi(c.GetString("max_monthly_salary"))
	pt, _ := strconv.Atoi(c.GetString("pay_times"))
	enterpriseID, _ := strconv.Atoi(c.GetString("enterprise_id"))
	job := models.Job{
		JobID:             0,
		Title:             c.GetString("title"),
		Description:       c.GetString("description"),
		Location:          c.GetString("location"),
		MinMonthlySalary:  minS,
		MaxMonthlySalary:  maxS,
		PayTimes:          pt,
		RequireEducation:  c.GetString("require_education"),
		RequireExperience: c.GetString("require_experience"),
		Type:              c.GetString("type"),
		Nature:            c.GetString("nature"),
		Status:            0,
		Department:        c.GetString("department"),
		CreateTime:        time.Now().In(cstZone),
		ModifyTime:        time.Now().In(cstZone),
		Enterprise:        &models.Enterprise{EnterpriseID: enterpriseID},
		SkillTags:         c.GetString("skill_tags"),
	}
	err := job.InsertOrUpdate()
	if err != nil {
		logs.Error("Error JobController NewJob: ", err)
		c.JsonResult(1, "添加职位失败！")
	}
	c.JsonResult(0, "ok")
}

// GetJob 根据职位ID返回给前端职位信息
func (c *JobController) GetJob() {
	jobID, _ := strconv.Atoi(c.GetString("job_id"))
	job := models.GetJobByID(jobID)
	c.JsonResult(0, "ok", job)
}

// EditJob 编辑职位信息
func (c *JobController) EditJob() {
	var cstZone = time.FixedZone("CST", 8*3600)
	minS, _ := strconv.Atoi(c.GetString("min_monthly_salary"))
	maxS, _ := strconv.Atoi(c.GetString("max_monthly_salary"))
	pt, _ := strconv.Atoi(c.GetString("pay_times"))
	enterpriseID, _ := strconv.Atoi(c.GetString("enterprise_id"))
	jobID, _ := strconv.Atoi(c.GetString("job_id"))
	job := models.Job{
		JobID:             jobID,
		Title:             c.GetString("title"),
		Description:       c.GetString("description"),
		Location:          c.GetString("location"),
		MinMonthlySalary:  minS,
		MaxMonthlySalary:  maxS,
		PayTimes:          pt,
		Department:        c.GetString("department"),
		RequireEducation:  c.GetString("require_education"),
		RequireExperience: c.GetString("require_experience"),
		Type:              c.GetString("type"),
		Nature:            c.GetString("nature"),
		Status:            0,
		ModifyTime:        time.Now().In(cstZone),
		Enterprise:        &models.Enterprise{EnterpriseID: enterpriseID},
		SkillTags:         c.GetString("skill_tags"),
	}
	err := job.InsertOrUpdate("skill_tags", "department", "title", "description", "location", "min_monthly_salary",
		"max_monthly_salary", "pay_times", "require_education", "require_experience", "type", "nature", "modify_time")
	if err != nil {
		logs.Error("Error JobController EditJob: ", err)
		c.JsonResult(1, "修改职位失败！")
	}
	c.JsonResult(0, "ok")
}

// DeleteJob 删除职位
func (c *JobController) DeleteJob() {
	jobID, _ := strconv.Atoi(c.GetString("job_id"))
	job := models.Job{JobID: jobID, Status: 1}
	err := job.InsertOrUpdate("Status")
	if err != nil {
		logs.Error("Error JobController DeleteJob: ", err)
		c.JsonResult(1, "删除职位失败")
	}
	c.JsonResult(0, "ok")
}

// GetJobs 获取所有职位信息
func (c *JobController) GetJobs() {
	jobs, err := models.GetAllJobs()
	if err != nil {
		logs.Error("Error JobController GetJobs: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok", jobs)
	}
}

// SetJobStatus 设置职位状态
func (c *JobController) SetJobStatus() {
	jobID, _ := strconv.Atoi(c.GetString("job_id"))
	status, _ := strconv.Atoi(c.GetString("status"))
	var job = models.Job{
		JobID:  jobID,
		Status: status,
	}
	err := job.InsertOrUpdate("status")
	if err != nil {
		logs.Error("Error JobController SetJobStatus:", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}
