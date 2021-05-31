package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/common"
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

	//获取所有省份信息
	provinces, err := common.GetAllProvinces()
	if err != nil {
		logs.Error("Error get provinces: ", err)
	}
	c.Data["Provinces"] = provinces

	c.TplName = "navigation/job.html"
}

func (c *NavigationController) Enterprise() {

	//获取全部企业信息
	enterprises, err := models.GetAllEnterprises()
	if err != nil {
		logs.Error("Error get enterprises: ", err)
	}
	c.Data["Enterprises"] = enterprises

	//获取所有省份信息
	provinces, err := common.GetAllProvinces()
	if err != nil {
		logs.Error("Error get provinces: ", err)
	}
	c.Data["Provinces"] = provinces

	c.TplName = "navigation/enterprise.html"
}

// UserCenter 用户中心导航
func (c *NavigationController) UserCenter() {
	c.TplName = "navigation/userCenter.html"
	c.Data["Member"] = c.Member
	//获取所有职位信息
	// JobTypeOption 用以返回职位类型
	type JobTypeOption struct {
		Value    string          `json:"value"`
		Label    string          `json:"label"`
		Children []JobTypeOption `json:"children"`
	}
	var jobTypesToReturn []JobTypeOption
	firstLevelTypes, _ := models.GetFirstLevelJobTypes()
	for _, f := range firstLevelTypes {
		var tmp JobTypeOption //一级分类的临时变量
		tmp.Label = f.Name
		tmp.Value = f.Name

		//为一级分类添加二级分类
		secondLevelTypes, err := models.GetSecondLevelJobTypesByParentID(f.TypeID)
		if err != nil {
			logs.Error("Error NavigationController EnterpriseHome: ", err)
		}
		for _, s := range secondLevelTypes {
			var sTmp JobTypeOption
			sTmp.Label = s.Name
			sTmp.Value = s.Name
			/*为二级分类添加三级子分类*/
			thirdLevelTypes, err := models.GetThirdLevelJobTypesByParentID(s.TypeID)
			if err != nil {
				logs.Error("Error NavigationController EnterpriseHome: ", err)
			}
			for _, t := range thirdLevelTypes {
				var tTmp JobTypeOption
				tTmp.Label = t.Name
				tTmp.Value = t.Name
				sTmp.Children = append(sTmp.Children, tTmp)
			}

			tmp.Children = append(tmp.Children, sTmp)
		}
		jobTypesToReturn = append(jobTypesToReturn, tmp)

	}
	c.Data["JobTypes"] = jobTypesToReturn
}

// EnterpriseHome 企业中心导航
func (c *NavigationController) EnterpriseHome() {

	enterprise := models.GetEnterpriseByMemberID(c.Member.MemberId)
	//为企业加载对应的职位信息
	enterprise.LoadJobs()
	c.Data["Enterprise"] = enterprise
	c.Data["Jobs"] = enterprise.Jobs
	c.Data["Member"] = c.Member

	//获取所有省份信息
	provinces, err := common.GetAllProvinces()
	if err != nil {
		logs.Error("Error get provinces: ", err)
	}
	c.Data["Provinces"] = provinces

	//获取所有职位信息
	// JobTypeOption 用以返回职位类型
	type JobTypeOption struct {
		Value    string          `json:"value"`
		Label    string          `json:"label"`
		Children []JobTypeOption `json:"children"`
	}
	var jobTypesToReturn []JobTypeOption
	firstLevelTypes, err := models.GetFirstLevelJobTypes()
	for _, f := range firstLevelTypes {
		var tmp JobTypeOption //一级分类的临时变量
		tmp.Label = f.Name
		tmp.Value = f.Name

		//为一级分类添加二级分类
		secondLevelTypes, err := models.GetSecondLevelJobTypesByParentID(f.TypeID)
		if err != nil {
			logs.Error("Error NavigationController EnterpriseHome: ", err)
		}
		for _, s := range secondLevelTypes {
			var sTmp JobTypeOption
			sTmp.Label = s.Name
			sTmp.Value = s.Name
			/*为二级分类添加三级子分类*/
			thirdLevelTypes, err := models.GetThirdLevelJobTypesByParentID(s.TypeID)
			if err != nil {
				logs.Error("Error NavigationController EnterpriseHome: ", err)
			}
			for _, t := range thirdLevelTypes {
				var tTmp JobTypeOption
				tTmp.Label = t.Name
				tTmp.Value = t.Name
				sTmp.Children = append(sTmp.Children, tTmp)
			}

			tmp.Children = append(tmp.Children, sTmp)
		}
		jobTypesToReturn = append(jobTypesToReturn, tmp)

	}
	c.Data["JobTypes"] = jobTypesToReturn
	c.TplName = "navigation/enterprise-home.html"
}

// Panel 控制面板
func (c *NavigationController) Panel() {
	c.TplName = "navigation/panel.html"

	type Res struct {
		TypeID     int    `orm:"column(type_id)" json:"type_id"`
		Name       string `json:"name"`
		ParentName string `json:"parent_name"`
		ParentID   int    `orm:"column(parent_id)" json:"parent_id"`
		Level      string `json:"level"`
	}
	var res []Res
	//获取所有职位类型
	types, err := models.GetAllTypes()
	if err != nil {
		logs.Error("Error NavigationController Panel: ", err)
	}
	//获取所有类型的父类型名
	for _, t := range types {
		parentName, err := t.GetParentName()
		if err != nil {
			logs.Error("Error get parent name: ", err)
		}
		res = append(res, Res{
			TypeID:     t.TypeID,
			Name:       t.Name,
			ParentID:   t.ParentID,
			Level:      t.Level,
			ParentName: parentName,
		})
	}
	c.Data["JobTypes"] = res

}
