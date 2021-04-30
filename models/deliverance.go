package models

import "github.com/astaxie/beego/orm"

type Deliverance struct {
	DeliveranceID int    `orm:"pk;auto;column(deliverance_id)" json:"deliverance_id"` //投递ID
	MemberID      int    `orm:"column(member_id)" json:"member_id"`                   //用户ID
	JobID         int    `orm:"column(job_id)" json:"job_id"`                         //职位ID
	Status        string `orm:"default(待查看)" json:"status"`                           //投递状态
}

func (d *Deliverance) TableName() string {
	return TNDeliverance()
}

func (d *Deliverance) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var deliverance Deliverance
	err := o.QueryTable(TNDeliverance()).Filter("deliverance_id", d.DeliveranceID).One(&deliverance)
	if deliverance.DeliveranceID > 0 {
		_, err = o.Update(d, fields...)
	} else {
		_, err = o.Insert(d)
	}
	return err
}
