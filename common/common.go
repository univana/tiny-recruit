package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// SessionName session 名称
const SessionName = "__tiny-recruit_session__"

/*//正则表达式
const RegexpEmail = `^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`

// 默认PageSize
const PageSize = 20
const RollPage = 4*/

const WorkingDirectory = "./"

//Role 返回用户权限名
func Role(role int) string {

	if role == 0 {
		return "管理员"
	} else if role == 1 {
		return "求职者"
	} else {
		return "企业用户"
	}
}

// AppKey app_key
func AppKey() string {
	return beego.AppConfig.DefaultString("app_key", "tiny-recruit")
}

// DefaultAvatar 默认头像
func DefaultAvatar() string {
	return beego.AppConfig.DefaultString("avatar", "/static/images/avatar.jpg")
}

type Province struct {
	ProvinceID   string `orm:"column(pro_id)" json:"pro_id"`
	ProvinceName string `orm:"column(pro_name)" json:"pro_name"`
}

// GetAllProvinces 获取所有的省
func GetAllProvinces() ([]Province, error) {
	o := orm.NewOrm()
	var provinces []Province
	sql := "select * from t_province"
	_, err := o.Raw(sql).QueryRows(&provinces)
	if err != nil {
		return nil, err
	}
	return provinces, err
}
