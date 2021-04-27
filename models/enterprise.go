package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Enterprise struct {
	EnterpriseID   int       `orm:"pk;auto;column(enterprise_id)" json:"enterprise_id"`            //企业ID
	Name           string    `json:"name"`                                                         //企业名称
	Description    string    `orm:"type(text)" json:"description"`                                 //企业简介
	Location       string    `json:"location"`                                                     //企业所在地址
	Status         int       `orm:"default(0)" json:"status"`                                      //企业状态
	Cover          string    `orm:"default(/static/images/covers/default-cover.jpg)" json:"cover"` //企业封面
	Type           string    `json:"type"`                                                         //行业类型
	FinancingStage string    `json:"financing_stage"`                                              //融资阶段
	Scale          string    `json:"scale"`                                                        //企业规模
	CreateTime     time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`                //创建时间
	LastLoginTime  time.Time `orm:"type(datetime);null" json:"last_login_time"`                    //最后登录时间
}

func (m *Enterprise) TableName() string {
	return TNEnterprise()
}

// GetAllEnterprises 获取所有的企业数据
func GetAllEnterprises() ([]*Enterprise, error) {
	o := orm.NewOrm()
	var enterprises []*Enterprise
	_, err := o.QueryTable(TNEnterprise()).All(&enterprises)
	return enterprises, err
}

// FilterEnterprises 过滤企业
func FilterEnterprises(city string, stage string, scale string, enterpriseType string) ([]*Enterprise, error) {
	o := orm.NewOrm()
	var enterprises []*Enterprise
	var err error
	qs := o.QueryTable(TNEnterprise())
	cond := orm.NewCondition()

	//按城市过滤
	if city != "不限" {
		cond = cond.And("location", city)
	}
	//按融资状态过滤
	if stage != "不限" {
		cond = cond.And("financing_stage", stage)
	}
	//按企业规模过滤
	if scale != "不限" {
		cond = cond.And("scale", scale)
	}
	//按行业领域过滤
	if enterpriseType != "不限" {
		cond = cond.And("type", enterpriseType)
	}
	_, err = qs.SetCond(cond).All(&enterprises)

	return enterprises, err

}
