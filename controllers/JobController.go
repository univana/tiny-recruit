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
	c.TplName = "job/job-detail.html"
}
