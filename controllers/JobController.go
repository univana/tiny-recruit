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
		DeliverTime time.Time     `json:"deliver_time"` //投递时间
		ModifyTime  time.Time     `json:"modify_time"`  //投递修改时间
		Status      string        `json:"status"`       //投递状态
		Resume      models.Resume `json:"resume"`       //投递对应的简历
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
		type Data struct {
			DeliverTime time.Time
			ModifyTime  time.Time
			Status      string
		}
		var res Data
		sql := "select deliver_time,modify_time,status from t_deliverance where member_id = ? and job_id = ?"
		o := orm.NewOrm()
		err := o.Raw(sql, memberID, jobID).QueryRow(&res)
		if err != nil {
			logs.Error("Error JobController query t_deliverance: ", err)
			c.JsonResult(1, "获取投递数据错误")
		}
		deliverList = append(deliverList, Deliver{DeliverTime: res.DeliverTime, ModifyTime: res.ModifyTime, Status: res.Status, Resume: resume})
	}
	c.JsonResult(0, "ok", deliverList)

}
