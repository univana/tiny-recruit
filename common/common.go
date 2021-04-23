package common

import "github.com/astaxie/beego"

// session
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

// app_key
func AppKey() string {
	return beego.AppConfig.DefaultString("app_key", "tiny-recruit")
}

//默认头像
func DefaultAvatar() string {
	return beego.AppConfig.DefaultString("avatar", "/static/images/avatar.jpg")
}

/*//TODO :fix
//默认封面
func DefaultCover() string {
	return beego.AppConfig.DefaultString("cover", "/static/images/book.png")
}
*/
