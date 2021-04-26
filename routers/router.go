package routers

import (
	"myApp/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//欢迎页面
	beego.Router("/", &controllers.NavigationController{}, "*:Home")
	beego.Router("/home", &controllers.NavigationController{}, "*:Home")
	beego.Router("/job", &controllers.NavigationController{}, "*:Job")
	beego.Router("/enterprise", &controllers.NavigationController{}, "*:Enterprise")
	beego.Router("/userCenter", &controllers.NavigationController{}, "*:UserCenter")

	beego.Router("/userCenter/getResume", &controllers.ResumeController{}, "*:GetResumeByMemberID")
	//注册
	beego.Router("/regist", &controllers.AccountController{}, "*:Regist")
	beego.Router("/doregist", &controllers.AccountController{}, "post:DoRegist")

	//登录
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")
	beego.Router("/login", &controllers.AccountController{}, "*:Login")

	beego.Router("/enterprise/filter/financing-stage", &controllers.EnterpriseController{}, "*:FilterByStage")

}
