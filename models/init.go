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
	return "md_members"
}

//TNResume return-用户表名
func TNResume() string {
	return "md_resume"
}
