package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
	"strconv"
	"time"
)

type CollectionController struct {
	BaseController
}

// Collect 搜藏职位
func (c *CollectionController) Collect() {
	jobID, _ := strconv.Atoi(c.GetString("job_id"))
	var collection = models.Collection{
		CollectionID: 0,
		MemberID:     c.Member.MemberId,
		JobID:        jobID,
		CreateTime:   time.Now().In(time.Local),
		Status:       0,
	}
	err := collection.InsertOrUpdate()
	if err != nil {
		logs.Error("Error CollectionController Collect: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}

}
