package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Enterprise struct {
	EnterpriseID   int       `orm:"pk;auto;column(enterprise_id)" json:"enterprise_id"`            //企业ID
	Name           string    `json:"name"`                                                         //企业名称
	Description    string    `orm:"type(text)" json:"description"`                                 //企业简介
	Location       string    `json:"location"`                                                     //企业所在地址
	Status         int       `orm:"default(0)" json:"status"`                                      //企业状态 0:正常 1:禁用 2:删除
	Cover          string    `orm:"default(/static/images/covers/default-cover.jpg)" json:"cover"` //企业封面
	Licence        string    `json:"licence"`                                                      //企业营业执照
	Type           string    `json:"type"`                                                         //行业类型
	Verified       int       `orm:"default(0)" json:"verified"`                                    //是否认证
	FinancingStage string    `json:"financing_stage"`                                              //融资阶段
	Scale          string    `json:"scale"`                                                        //企业规模
	CreateTime     time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`                //创建时间
	LastLoginTime  time.Time `orm:"type(datetime);null" json:"last_login_time"`                    //最后登录时间

	Jobs   []*Job  `orm:"reverse(many)" json:"jobs"` //企业和职位的一对多关系
	Member *Member `orm:"rel(one)"`                  //企业所属的用户
}

func (e *Enterprise) TableName() string {
	return TNEnterprise()
}

// GetAllEnterprises 获取所有的企业数据
func GetAllEnterprises() ([]*Enterprise, error) {
	o := orm.NewOrm()
	var enterprises []*Enterprise
	_, err := o.QueryTable(TNEnterprise()).Exclude("status", 1).All(&enterprises)
	return enterprises, err
}

// GetEnterpriseByID 根据ID查询企业信息
func GetEnterpriseByID(id int) (Enterprise, error) {
	o := orm.NewOrm()
	enterprise := Enterprise{EnterpriseID: id}

	err := o.Read(&enterprise)

	return enterprise, err
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
	_, err = qs.SetCond(cond).Exclude("status", 1).All(&enterprises)

	return enterprises, err
}

// GetEnterpriseByMemberID 根据用户ID或取所属的企业
func GetEnterpriseByMemberID(memberID int) Enterprise {
	o := orm.NewOrm()
	var enterprise Enterprise
	err := o.QueryTable(TNEnterprise()).Filter("member_id", memberID).One(&enterprise)
	if err != nil {
		logs.Error("Error get enterprise: ", err)
		return Enterprise{}
	}
	return enterprise
}

// LoadJobs 加载企业对应的所有职位信息
func (e *Enterprise) LoadJobs() {
	o := orm.NewOrm()
	_, err := o.LoadRelated(e, "Jobs", "status", 0)
	if err != nil {
		logs.Error("Error load jobs: ", err)
	}
}

// InsertOrUpdate 添加或更新企业信息
func (e *Enterprise) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var enterprise Enterprise
	err := o.QueryTable(TNEnterprise()).Filter("enterprise_id", e.EnterpriseID).One(&enterprise)
	if enterprise.EnterpriseID > 0 {
		_, err = o.Update(e, fields...)
	} else {
		_, err = o.Insert(e)
	}
	return err
}

// GetLicencePath 获取营业执照地址
func GetLicencePath(id int) (string, error) {
	o := orm.NewOrm()
	sql := "select licence from t_enterprise where enterprise_id = ?"
	var path string
	err := o.Raw(sql, id).QueryRow(&path)
	if err != nil {
		return "", err
	}
	return path, nil
}
