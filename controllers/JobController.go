package controllers

import (
	"myApp/models"
	"strconv"
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
