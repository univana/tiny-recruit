package models

import "time"

type Collection struct {
	CollectionID int       `orm:"pk;auto;column(collection_id)" json:"collection_id"` //收藏ID
	MemberID     int       `json:"member_id"`                                         //用户ID
	JobID        int       `json:"job_id"`                                            //职位ID
	CreateTime   time.Time `json:"create_time"`                                       //创建时间
	Status       int       `orm:"default(0)" json:"status"`                           //收藏状态
}

func (c *Collection) TableName() string {
	return TNCollection()
}
