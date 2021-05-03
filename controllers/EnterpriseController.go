package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
	"strconv"
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

// ShowEnterprise 展示企业详情页
func (c *EnterpriseController) ShowEnterprise() {
	enterpriseID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	enterprise, err := models.GetEnterpriseByID(enterpriseID)
	if err != nil {
		logs.Error("Error EnterpriseController ShowEnterprise: ", err)
	}
	enterprise.LoadJobs()
	c.Data["Enterprise"] = enterprise
	c.TplName = "enterprise/enterprise-detail.html"
}
