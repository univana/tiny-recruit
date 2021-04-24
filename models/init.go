package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(
		new(Member),
		new(Resume),
	)
}

//TNMembers return-用户表名
func TNMembers() string {
	return "t_member"
}

//TNResume return-用户表名
func TNResume() string {
	return "t_resume"
}
