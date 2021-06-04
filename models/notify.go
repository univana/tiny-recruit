package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Notify struct {
	NotifyID   int       `orm:"pk;auto;column(notify_id)" json:"notify_id"`     //通知ID
	ReceiverID int       `orm:"column(receiver_id)" json:"receiver_id"`         //接收方ID
	Type       string    `json:"type"`                                          //通知类型
	Title      string    `json:"title"`                                         //通知标题
	Content    string    `orm:"type(text)" json:"content"`                      //通知内容
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //通知创建时间
	Read       int       `json:"read"`                                          //是否已读
	Deleted    int       `json:"deleted"`                                       //是否删除
}

func (n *Notify) TableName() string {
	return TNNotify()
}

// Insert 插入通知
func (n *Notify) Insert() error {
	o := orm.NewOrm()
	_, err := o.Insert(n)
	return err
}

// GetAllNotifiesByMemberID 根据接受者ID获取所有的通知信息
func GetAllNotifiesByMemberID(id int) ([]Notify, error) {
	o := orm.NewOrm()
	var notifies []Notify
	_, err := o.QueryTable(TNNotify()).Filter("receiver_id", id).Filter("deleted", 0).All(&notifies)
	if err != nil {
		return nil, err
	} else {
		return notifies, nil
	}
}

// Update 更新消息通知
func (n *Notify) Update(fields ...string) error {
	o := orm.NewOrm()
	_, err := o.Update(n, fields...)
	return err
}

func HasUnread(memberID int) bool {
	o := orm.NewOrm()
	var notifies []Notify
	all, _ := o.QueryTable(TNNotify()).Filter("receiver_id", memberID).Filter("deleted", 0).Filter("read", 1).All(&notifies)
	if all > 0 {
		return true
	} else {
		return false
	}
}
