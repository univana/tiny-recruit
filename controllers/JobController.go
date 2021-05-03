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
