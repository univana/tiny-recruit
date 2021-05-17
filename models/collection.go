package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Collection struct {
	CollectionID int       `orm:"pk;auto;column(collection_id)" json:"collection_id"` //收藏ID
	MemberID     int       `orm:"column(member_id)" json:"member_id"`                 //用户ID
	JobID        int       `orm:"column(job_id)" json:"job_id"`                       //职位ID
	CreateTime   time.Time `json:"create_time"`                                       //创建时间
	Status       int       `orm:"default(0)" json:"status"`                           //收藏状态
}

func (c *Collection) TableName() string {
	return TNCollection()
}

// InsertOrUpdate 添加或更新收藏信息
func (c *Collection) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var collection Collection
	err := o.QueryTable(TNCollection()).Filter("collection_id", c.CollectionID).One(&collection)
	if collection.CollectionID > 0 {
		_, err = o.Update(c, fields...)
	} else {
		_, err = o.Insert(c)
	}
	return err

}
