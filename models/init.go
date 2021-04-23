package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(
		new(Member),
		new(Notebook),
		new(Diary),
		new(Template),
	)
}

//TNMembers :return-用户表名
func TNMembers() string {
	return "md_members"
}

//TNNotebooks :return-笔记本表名
func TNNotebooks() string {
	return "md_notebooks"
}

//TNDiaries :return-日记表名
func TNDiaries() string {
	return "md_diaries"
}

//TNTemplates :return-模板表名
func TNTemplates() string {
	return "md_templates"
}
