package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
	"strconv"
	"time"
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

func (c *DeliveranceController) ChangeStatus() {
	//设置时区
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	status := c.GetString("status")
	deliveranceID, _ := strconv.Atoi(c.GetString("deliverance_id"))

	deliverance := models.Deliverance{DeliveranceID: deliveranceID, Status: status, ModifyTime: time.Now().In(cstZone)}
	err := deliverance.InsertOrUpdate("status", "modify_time")
	if err != nil {
		logs.Error("Error DeliveranceController ChangeStatus: ", err)
		c.JsonResult(1, "修改状态错误!")
	} else {
		c.JsonResult(0, "ok")
	}
}
