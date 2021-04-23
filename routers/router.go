package routers

import (
	"myApp/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//欢迎页面
	beego.Router("/", &controllers.HomeController{}, "*:Welcome")
	//注册
	beego.Router("/regist", &controllers.AccountController{}, "*:Regist")
	beego.Router("/doregist", &controllers.AccountController{}, "post:DoRegist")

	//登录
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")
	beego.Router("/login", &controllers.AccountController{}, "*:Login")

	//日记本
	beego.Router("/notebook", &controllers.NotebookController{}, "*:Index")
	beego.Router("/notebook/add", &controllers.NotebookController{}, "post:AddNotebook")
	beego.Router("/notebook/show/:key", &controllers.NotebookController{}, "*:ShowNotebook")
	beego.Router("/notebook/verify", &controllers.NotebookController{}, "*:TokenVerify")
	beego.Router("/notebook/download/:key",&controllers.NotebookController{},"*:DownloadNotebook")

	//日记
	beego.Router("/diary/edit", &controllers.DiaryController{}, "*:Edit")
	beego.Router("/diary/save", &controllers.DiaryController{}, "*:SaveDiary")
	beego.Router("/diary/show/:key", &controllers.DiaryController{}, "*:ShowDiary")
	beego.Router("/diary/download/:key", &controllers.DiaryController{}, "*:DownloadDiary")

	//管理
	beego.Router("/manage", &controllers.ManageController{}, "*:Index")
	beego.Router("/manage/diary/delete", &controllers.DiaryController{}, "*:DeleteDiary")
	beego.Router("/manage/notebook/delete", &controllers.NotebookController{}, "*:DeleteNotebook")

	//模板
	beego.Router("/template/save", &controllers.TemplateController{}, "*:SaveTemplate")
	beego.Router("/manage/template/delete", &controllers.TemplateController{}, "*:DeleteTemplate")

	//检索
	beego.Router("/search", &controllers.SearchController{}, "*:Index")
	beego.Router("/search/title", &controllers.DiaryController{}, "*:SearchDiary")
	beego.Router("/search/tag", &controllers.DiaryController{}, "*:GetDiaryListByTag")

}
