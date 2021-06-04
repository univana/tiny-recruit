package controllers

import (
	"github.com/astaxie/beego/logs"
	"myApp/models"
	"strconv"
)

type NotifyController struct {
	BaseController
}

func (c *NotifyController) Send() {
	receiverID, _ := strconv.Atoi(c.GetString("receiver_id"))
	var notify = models.Notify{
		NotifyID:   0,
		ReceiverID: receiverID,
		Type:       c.GetString("type"),
		Content:    c.GetString("content"),
		Title:      c.GetString("title"),
		Read:       0,
		Deleted:    0,
	}
	err := notify.Insert()
	if err != nil {
		logs.Error("Error Send notify: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "发送通知成功！")
	}
}

// GetAllNotifiesByMemberID 根据接受者ID获取所有通知消息
func GetAllNotifiesByMemberID(id int) ([]models.Notify, error) {
	notifies, err := GetAllNotifiesByMemberID(id)
	return notifies, err
}

// Read 将消息通知设为已读
func (c *NotifyController) Read() {
	notifyID, _ := strconv.Atoi(c.GetString("notify_id"))
	var notify = models.Notify{
		NotifyID: notifyID,
		Read:     1,
	}
	err := notify.Update("read")
	if err != nil {
		logs.Error("Error update notify:", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "设置已读成功！")
	}
}

// Refresh 刷新消息通知
func (c *NotifyController) Refresh() {
	notifies, err := models.GetAllNotifiesByMemberID(c.Member.MemberId)
	if err != nil {
		logs.Error("Error refresh notifies:", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok", notifies)
	}
}

func (c *NotifyController) Delete() {
	notifyID, _ := strconv.Atoi(c.GetString("notify_id"))
	var notify = models.Notify{
		NotifyID: notifyID,
		Deleted:  1,
	}
	err := notify.Update("deleted")
	if err != nil {
		logs.Error("Error delete notify:", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "刪除消息通知成功！")
	}
}
