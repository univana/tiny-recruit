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
