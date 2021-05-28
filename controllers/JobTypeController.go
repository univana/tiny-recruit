package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
	"strconv"
)

type JobTypeController struct {
	BaseController
}

// AddChildType 添加子职位类型
func (c *JobTypeController) AddChildType() {
	name := c.GetString("name")
	parentID, _ := strconv.Atoi(c.GetString("parent_id"))
	parentLevel := c.GetString("parent_level")
	var level string
	if parentLevel == "一级" {
		level = "二级"
	} else if parentLevel == "二级" {
		level = "三级"
	}
	var jobType = models.JobType{
		TypeID:   0,
		Name:     name,
		ParentID: parentID,
		Level:    level,
		Deleted:  0,
	}
	err := jobType.InsertOrUpdate()
	if err != nil {
		logs.Error("Error JobTypeController AddChildType: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}
