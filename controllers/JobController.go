package controllers

import (
	"fmt"
	"strconv"
)

type JobController struct {
	BaseController
}

func (c *JobController) ShowJob() {
	jobID, _ := strconv.Atoi(c.Ctx.Input.Param(":key"))
	fmt.Println(jobID)
	c.TplName = "job/show.html"
}
