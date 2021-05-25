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

func (c *EnterpriseController) Add() {
	//获取logo
	_, info, err := c.GetFile("file")

	//获取文件后缀
	//保存图片
	ext := strings.ToLower(path.Ext(info.Filename))
	v4, _ := uuid.NewV4()

	logoName := fmt.Sprintf("%s%s", v4, ext)
	logoPath := fmt.Sprintf("static/images/covers/%s", logoName)
	err = c.SaveToFile("file", logoPath)
	if err != nil {
		logs.Error("Error EnterpriseController Add: ", err)
	}

	//获取licence
	_, info, err = c.GetFile("file-licence")

	//获取文件后缀
	//保存图片
	ext = strings.ToLower(path.Ext(info.Filename))
	v4, _ = uuid.NewV4()

	licenceName := fmt.Sprintf("%s%s", v4, ext)
	licencePath := fmt.Sprintf("static/images/licences/%s", licenceName)
	err = c.SaveToFile("file", licencePath)
	if err != nil {
		logs.Error("Error EnterpriseController Add: ", err)
	}

	//更新企业信息
	enterpriseID, _ := strconv.Atoi(c.GetString("enterprise_id"))
	var enterprise = models.Enterprise{
		Member:         &models.Member{MemberId: c.Member.MemberId},
		EnterpriseID:   enterpriseID,
		Name:           c.GetString("name"),
		Type:           c.GetString("type"),
		FinancingStage: c.GetString("financing_stage"),
		Scale:          c.GetString("scale"),
		Description:    c.GetString("description"),
		Location:       c.GetString("location"),
		Cover:          "/" + logoPath,
		Licence:        "/" + licencePath,
	}
	err = enterprise.InsertOrUpdate()
	if err != nil {
		logs.Error("Error EnterpriseController Add: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}

}

// GetEnterprises 获取所有企业数据
func (c *EnterpriseController) GetEnterprises() {
	enterprises, err := models.GetAllEnterprises()
	if err != nil {
		logs.Error("Error EnterpriseController GetEnterprises: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok", enterprises)
	}
}

// SetEnterpriseStatus 设置企业的状态
func (c *EnterpriseController) SetEnterpriseStatus() {
	enterpriseID, _ := strconv.Atoi(c.GetString("enterprise_id"))
	status, _ := strconv.Atoi(c.GetString("status"))
	var enterprise = models.Enterprise{
		EnterpriseID: enterpriseID,
		Status:       status,
	}
	err := enterprise.InsertOrUpdate("status")
	if err != nil {
		logs.Error("Error EnterpriseController SetEnterpriseStatus:", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}

// GetLicence 获取营业执照
func (c *EnterpriseController) GetLicence() {
	enterpriseID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	licencePath, err := models.GetLicencePath(enterpriseID)
	if err != nil {
		logs.Error("Error EnterpriseController GetLicence: ", err)
	} else {
		licencePath = licencePath[1:]
		c.Ctx.Output.Download(licencePath)
	}
}
