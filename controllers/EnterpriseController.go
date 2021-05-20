package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	uuid "github.com/iris-contrib/go.uuid"
	"myApp/models"
	"path"
	"strconv"
	"strings"
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

func (c *EnterpriseController) Edit() {
	c.TplName = "navigation/enterprise-home.html"

	_, info, err := c.GetFile("file")
	if info != nil {
		//获取文件后缀
		//保存图片
		ext := strings.ToLower(path.Ext(info.Filename))
		v4, _ := uuid.NewV4()

		fileName := fmt.Sprintf("%s%s", v4, ext)
		filePath := fmt.Sprintf("static/images/covers/%s", fileName)
		err = c.SaveToFile("file", filePath)
		if err != nil {
			logs.Error("Error EnterpriseController Edit: ", err)
		}

		//更新企业信息
		enterpriseID, _ := strconv.Atoi(c.GetString("enterprise_id"))
		var enterprise = models.Enterprise{
			EnterpriseID:   enterpriseID,
			Name:           c.GetString("name"),
			Type:           c.GetString("type"),
			FinancingStage: c.GetString("financing_stage"),
			Scale:          c.GetString("scale"),
			Description:    c.GetString("description"),
			Location:       c.GetString("location"),
			Cover:          "/" + filePath,
		}
		err := enterprise.InsertOrUpdate("location", "cover", "name", "type", "financing_stage", "scale", "description")
		if err != nil {
			logs.Error("Error EnterpriseController Edit: ", err)
			c.JsonResult(1, err.Error())
		} else {
			c.JsonResult(0, "ok")
		}

	} else {
		//更新企业信息
		enterpriseID, _ := strconv.Atoi(c.GetString("enterprise_id"))
		var enterprise = models.Enterprise{
			EnterpriseID:   enterpriseID,
			Name:           c.GetString("name"),
			Type:           c.GetString("type"),
			FinancingStage: c.GetString("financing_stage"),
			Scale:          c.GetString("scale"),
			Description:    c.GetString("description"),
			Location:       c.GetString("location"),
		}
		err := enterprise.InsertOrUpdate("location", "name", "type", "financing_stage", "scale", "description")
		if err != nil {
			logs.Error("Error EnterpriseController Edit: ", err)
			c.JsonResult(1, err.Error())
		} else {
			c.JsonResult(0, "ok")
		}
	}

}
