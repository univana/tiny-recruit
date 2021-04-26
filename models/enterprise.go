package models

import "time"

type Enterprise struct {
	EnterpriseID  int       `orm:"pk;auto" json:"enterprise_id"`                   //企业ID
	Name          string    `json:"name"`                                          //企业名称
	Description   string    `orm:"type(text)" json:"description"`                  //企业简介
	Location      string    `json:"location"`                                      //企业所在地址
	Status        int       `orm:"default(0)" json:"status"`                       //企业状态
	CreateTime    time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	LastLoginTime time.Time `orm:"type(datetime);null" json:"last_login_time"`     //最后登录时间
}

func (m *Enterprise) TableName() string {
	return TNEnterprise()
}
