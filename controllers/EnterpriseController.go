package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
)

type EnterpriseController struct {
	BaseController
}

func (c *EnterpriseController) FilterByStage() {
	stage := c.GetString("financing_stage")
	enterprises, err := models.FilterEnterprisesByStage(stage)
	if err != nil {
		logs.Error("Error filter enterprises by stage: ", err)
		c.JsonResult(1, "Error filter by stage")
	}
	c.JsonResult(0, "ok", enterprises)

	c.TplName = "navigation/enterprise.html"
}
