package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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
	c.TplName = "job/job-detail.html"
}

func (c *JobController) GetDeliverance() {
	jobID, _ := strconv.Atoi(c.GetString("job_id"))
	//获取该职位的所有投递信息
	type Deliver struct {
		DeliveranceID int           `json:"deliverance_id"`
		DeliverTime   time.Time     `json:"deliver_time"` //投递时间
		ModifyTime    time.Time     `json:"modify_time"`  //投递修改时间
		Status        string        `json:"status"`       //投递状态
		Resume        models.Resume `json:"resume"`       //投递对应的简历
	}
	var deliverList []Deliver
	//获取简历
	resumes, err := models.GetAllResumesByJobID(jobID)
	if err != nil {
		logs.Error("Error JobController get resumes: ", err)
	}

	for _, resume := range resumes {
		//获取投递时间、投递修改时间和投递状态
		memberID := resume.Member.MemberId
		var deliverance models.Deliverance
		o := orm.NewOrm()
		err := o.QueryTable(models.TNDeliverance()).Filter("member_id", memberID).Filter("job_id", jobID).One(&deliverance)
		if err != nil {
			logs.Error("Error JobController query t_deliverance: ", err)
			c.JsonResult(1, "获取投递数据错误")
		}
		deliverList = append(deliverList, Deliver{DeliveranceID: deliverance.DeliveranceID, DeliverTime: deliverance.DeliverTime, ModifyTime: deliverance.ModifyTime, Status: deliverance.Status, Resume: resume})
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
		CreateTime:        time.Now().In(cstZone),
		ModifyTime:        time.Now().In(cstZone),
		Enterprise:        &models.Enterprise{EnterpriseID: enterpriseID},
	}
	err := job.InsertOrUpdate()
	if err != nil {
		logs.Error("Error JobController NewJob: ", err)
		c.JsonResult(1, "添加职位失败！")
	}
	c.JsonResult(0, "ok")
}
