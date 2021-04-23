package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
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
	c.JsonResult(0, "ok", resume)
}
