package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
	"strconv"
)

type DeliveranceController struct {
	BaseController
}

// Deliver 投递简历的后端逻辑
func (c *DeliveranceController) Deliver() {
	jobID, _ := strconv.Atoi(c.GetString("job_id"))

	deliverance := models.Deliverance{JobID: jobID, MemberID: c.Member.MemberId, Status: "待查看"}
	err := deliverance.InsertOrUpdate()
	if err != nil {
		logs.Error("Error deliver :", err)
		c.JsonResult(1, "投递失败")
	} else {
		c.JsonResult(0, "投递成功")
	}

	c.TplName = "job/job-detail.html"
}
