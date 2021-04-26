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
	var enterprises []*models.Enterprise
	var err error

	if stage == "不限" {
		enterprises, err = models.GetAllEnterprises()
	} else {
		enterprises, err = models.FilterEnterprisesByStage(stage)
	}

	if err != nil {
		logs.Error("Error filter enterprises by stage: ", err)
		c.JsonResult(1, "Error filter by stage")
	}
	c.JsonResult(0, "ok", enterprises)

	c.TplName = "navigation/enterprise.html"
}

func (c *EnterpriseController) FilterByCity() {
	city := c.GetString("city")
	var enterprises []*models.Enterprise
	var err error

	if city == "不限" {
		enterprises, err = models.GetAllEnterprises()
	} else {
		enterprises, err = models.FilterEnterprisesByCity(city)
	}

	if err != nil {
		logs.Error("Error filter enterprises by city: ", err)
		c.JsonResult(1, "Error filter by city")
	}
	c.JsonResult(0, "ok", enterprises)

	c.TplName = "navigation/enterprise.html"
}

func (c *EnterpriseController) FilterByScale() {
	scale := c.GetString("scale")
	var enterprises []*models.Enterprise
	var err error

	if scale == "不限" {
		enterprises, err = models.GetAllEnterprises()
	} else {
		enterprises, err = models.FilterEnterprisesByScale(scale)
	}

	if err != nil {
		logs.Error("Error filter enterprises by scale: ", err)
		c.JsonResult(1, "Error filter by scale")
	}
	c.JsonResult(0, "ok", enterprises)

	c.TplName = "navigation/enterprise.html"
}
