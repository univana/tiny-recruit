package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
)

type EnterpriseController struct {
	BaseController
}

func (c *EnterpriseController) Filter() {
	city := c.GetString("city")
	stage := c.GetString("financing_stage")
	scale := c.GetString("scale")
	enterpriseType := c.GetString("type")

	enterprises, err := models.FilterEnterprises(city, stage, scale, enterpriseType)
	if err != nil {
		logs.Error("Error filter enterprises: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok", enterprises)
	}
	c.TplName = "navigation/enterprise.html"
}
