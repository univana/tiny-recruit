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
	beego.Router("/enterprise/home", &controllers.NavigationController{}, "*:EnterpriseHome")
	beego.Router("/panel", &controllers.NavigationController{}, "*:Panel")
	beego.Router("/userCenter/getResume", &controllers.ResumeController{}, "*:GetResumeByMemberID")

	//注册
	beego.Router("/regist", &controllers.AccountController{}, "*:Regist")
	beego.Router("/doregist", &controllers.AccountController{}, "post:DoRegist")

	//登录
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")
	beego.Router("/login", &controllers.AccountController{}, "*:Login")

	//企业过滤
	beego.Router("/enterprise/filter/filter", &controllers.EnterpriseController{}, "*:Filter")

	//职位内容展示
	beego.Router("/job/detail/?:id", &controllers.JobController{}, "*:ShowJob")

	//企业内容展示
	beego.Router("/enterprise/detail/?:id", &controllers.EnterpriseController{}, "*:ShowEnterprise")

	//投递简历
	beego.Router("/job/detail/deliver", &controllers.DeliveranceController{}, "*:Deliver")

	//查询职位对应的所有投递数据
	beego.Router("/enterprise/home/queryDeliverance", &controllers.JobController{}, "*:GetDeliverance")

	//修改投递状态
	beego.Router("/enterprise/home/changeStatus", &controllers.DeliveranceController{}, "*:ChangeStatus")

	//职位过滤器
	beego.Router("/job/filter", &controllers.JobController{}, "*:Filter")

	//新建职位
	beego.Router("/enterprise/home/newJob", &controllers.JobController{}, "*:NewJob")

	//编辑企业基本信息
	beego.Router("/enterprise/home/edit", &controllers.EnterpriseController{}, "*:Edit")

	//新建企业基本信息
	beego.Router("/enterprise/home/add", &controllers.EnterpriseController{}, "*:Add")

	//获取职位
	beego.Router("/enterprise/home/getJob", &controllers.JobController{}, "*:GetJob")

	//编辑职位
	beego.Router("/enterprise/home/editJob", &controllers.JobController{}, "*:EditJob")

	//删除职位
	beego.Router("/enterprise/home/deleteJob", &controllers.JobController{}, "*:DeleteJob")

	//编辑用户基本信息
	beego.Router("/userCenter/editAccount", &controllers.AccountController{}, "*:Edit")

	//获取用户的投递数据
	beego.Router("/userCenter/getDelivers", &controllers.AccountController{}, "*:GetDelivers")

	//获取用户收藏信息
	beego.Router("/userCenter/getCollections", &controllers.AccountController{}, "*:GetCollections")

	//取消收藏
	beego.Router("/userCenter/cancelCollection", &controllers.CollectionController{}, "*:Cancel")

	//编辑简历
	beego.Router("/userCenter/editResume", &controllers.ResumeController{}, "*:EditResume")

	//添加教育经历
	beego.Router("/userCenter/addEdu", &controllers.ResumeController{}, "*:AddEdu")

	//添加项目经历
	beego.Router("/userCenter/addProject", &controllers.ResumeController{}, "*:AddProject")

	//添加实习经历
	beego.Router("/userCenter/addInternship", &controllers.ResumeController{}, "*:AddInternship")

	//收藏职位
	beego.Router("/job/detail/collect", &controllers.CollectionController{}, "*:Collect")

	//获取所有用户信息
	beego.Router("/panel/getMembers", &controllers.AccountController{}, "*:GetMembers")

	//设置用户状态
	beego.Router("/panel/setMemberStatus", &controllers.AccountController{}, "*:SetMemberStatus")

	//设置用户权限
	beego.Router("/panel/setMemberRole", &controllers.AccountController{}, "*:SetMemberRole")

	//获取所有职位
	beego.Router("/panel/getJobs", &controllers.JobController{}, "*:GetJobs")

	//设置职位状态
	beego.Router("/panel/setJobStatus", &controllers.JobController{}, "*:SetJobStatus")

	//获取省份对应的城市信息
	beego.Router("/common/getCities", &controllers.BaseController{}, "*:GetCities")

}
