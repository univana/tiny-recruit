package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
)

type NavigationController struct {
	BaseController
}

func (c *NavigationController) Home() {
	c.TplName = "navigation/home.html"
}

func (c *NavigationController) Job() {
	//获取所有的职位信息
	jobs, err := models.GetAllJobs()
	if err != nil {
		logs.Error("Error get jobs:", err)
	}
	c.Data["Jobs"] = jobs

	c.TplName = "navigation/job.html"
}

func (c *NavigationController) Enterprise() {

	//获取全部企业信息
	enterprises, err := models.GetAllEnterprises()
	if err != nil {
		logs.Error("Error get enterprises: ", err)
	}
	c.Data["Enterprises"] = enterprises

	c.TplName = "navigation/enterprise.html"
}

// UserCenter 用户中心导航
func (c *NavigationController) UserCenter() {
	c.TplName = "navigation/userCenter.html"
	c.Data["Member"] = c.Member
}

// EnterpriseHome 企业中心导航
func (c *NavigationController) EnterpriseHome() {

	enterprise := models.GetEnterpriseByMemberID(c.Member.MemberId)
	//为企业加载对应的职位信息
	enterprise.LoadJobs()
	c.Data["Enterprise"] = enterprise
	c.Data["Jobs"] = enterprise.Jobs
	c.Data["Member"] = c.Member
	c.TplName = "navigation/enterprise-home.html"
}
