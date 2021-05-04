package controllers

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"myApp/common"
	"myApp/models"
	"myApp/utils"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type AccountController struct {
	BaseController
}

//注册
func (c *AccountController) Regist() {
	var (
		nickname string      //昵称
		avatar   string      //头像的http链接地址
		username string      //用户名
		id       interface{} //用户id
	)

	c.Data["Nickname"] = nickname
	c.Data["Avatar"] = avatar
	c.Data["Username"] = username
	c.Data["Id"] = id

	c.TplName = "account/regist.html"
}

//注册

func (c *AccountController) DoRegist() {
	var err error
	account := c.GetString("account")
	nickname := strings.TrimSpace(c.GetString("nickname"))
	password1 := c.GetString("password1")
	password2 := c.GetString("password2")

	member := models.NewMember()

	if password1 != password2 {
		c.JsonResult(1, "登录密码与确认密码不一致！")
	}

	if l := strings.Count(password1, ""); password1 == "" || l > 20 || l < 6 {
		c.JsonResult(1, "密码必须在6-20个字符之间")
	}

	if l := strings.Count(nickname, "") - 1; l < 2 || l > 20 {
		c.JsonResult(1, "用户昵称限制在2-20个字符")
	}

	member.Account = account
	member.Nickname = nickname
	/*对密码进行加密*/
	hash, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
	if err != nil {
		logs.Error("Error hash the password:", err)
	}
	encodePassword := string(hash)
	member.Password = encodePassword
	//TODO 设置注册时的角色类型
	member.Avatar = common.DefaultAvatar()
	member.Status = 0
	if err := member.Add(); err != nil {
		beego.Error(err)
		c.JsonResult(1, err.Error())
	}

	if err = c.login(member.MemberId); err != nil {
		beego.Error(err.Error())
		c.JsonResult(1, err.Error())
	}

	c.JsonResult(0, "注册成功")
}

//封装一个内部调用的函数，login
func (c *AccountController) login(memberId int) (err error) {
	member, err := models.NewMember().Find(memberId)
	if member.MemberId == 0 {
		return errors.New("用户不存在")
	}
	//如果没有数据
	if err != nil {
		return err
	}
	member.LastLoginTime = time.Now()
	member.Update()
	c.SetMember(*member)
	var remember CookieRemember
	remember.MemberId = member.MemberId
	remember.Account = member.Account
	remember.Time = time.Now()
	v, err := utils.Encode(remember)
	if err == nil {
		c.SetSecureCookie(common.AppKey(), "login", v, 60*5)
	}
	return err
}

//Logout :退出登录
func (c *AccountController) Logout() {
	c.SetMember(models.Member{})
	c.SetSecureCookie(common.AppKey(), "login", "", -3600)
	c.Redirect(beego.URLFor("AccountController.Login"), 302)
}

//登录
func (c *AccountController) Login() {
	var remember CookieRemember
	//验证cookie
	if cookie, ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
		if err := utils.Decode(cookie, &remember); err == nil {
			if err = c.login(remember.MemberId); err == nil {
				c.Redirect(beego.URLFor("HomeController.Welcome"), 302)
				return
			}
		}
	}
	c.TplName = "account/login.html"

	if c.Ctx.Input.IsPost() {
		account := c.GetString("account")
		password := c.GetString("password")
		member, err := models.NewMember().Login(account, password)
		if err != nil {
			c.JsonResult(1, "登录失败", nil)
		}
		member.LastLoginTime = time.Now()
		member.Update()
		c.SetMember(*member)
		remember.MemberId = member.MemberId
		remember.Account = member.Account
		remember.Time = time.Now()
		v, err := utils.Encode(remember)
		if err == nil {
			c.SetSecureCookie(common.AppKey(), "login", v, 24*3600*365)
		}
		c.JsonResult(0, "ok")
	}
}

func (c *AccountController) GetDelivers() {
	type Deliver struct {
		DeliveranceID int       `orm:"column(deliverance_id)" json:"deliverance_id"`
		MemberID      int       `orm:"column(member_id)" json:"member_id"`
		JobID         int       `orm:"column(job_id)" json:"job_id"`
		Title         string    `json:"title"`
		Status        string    `json:"status"`
		DeliverTime   time.Time `json:"deliver_time"`
		ModifyTime    time.Time `json:"modify_time"`
	}
	var delivers []Deliver
	o := orm.NewOrm()
	sql := "select deliverance_id,member_id,tj.title,tj.job_id,t_deliverance.status,deliver_time,t_deliverance.modify_time from t_deliverance join t_job tj on t_deliverance.job_id = tj.job_id where member_id=?"
	_, err := o.Raw(sql, c.Member.MemberId).QueryRows(&delivers)
	if err != nil {
		logs.Error("Error AccountController GetDelivers: ", err)
		c.JsonResult(1, "获取职位信息错误！")
	}
	c.JsonResult(0, "ok", delivers)

}
