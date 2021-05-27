package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/captcha"
	"golang.org/x/crypto/bcrypt"
	"myApp/common"
	"myApp/models"
	"myApp/utils"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type AccountController struct {
	BaseController
}

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
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

		//检查验证码是否正确
		res := cpt.VerifyReq(c.Ctx.Request)
		if res == false {
			c.JsonResult(1, "验证码错误！")
		} else {
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
	sql := "select deliverance_id,member_id,tj.title,tj.job_id,t_deliverance.status,deliver_time,t_deliverance.modify_time from t_deliverance join t_job tj on t_deliverance.job_id = tj.job_id where t_deliverance.status <> '已撤销' and member_id=?"
	_, err := o.Raw(sql, c.Member.MemberId).QueryRows(&delivers)
	if err != nil {
		logs.Error("Error AccountController GetDelivers: ", err)
		c.JsonResult(1, "获取职位信息错误！")
	}
	c.JsonResult(0, "ok", delivers)

}

// GetCollections 获取用户的收藏信息
func (c *AccountController) GetCollections() {
	type Data struct {
		CollectionID int        `json:"collection_id"`
		Job          models.Job `json:"job"`
		CreateTime   time.Time  `json:"create_time"`
		Status       int        `json:"status"`
	}
	var res []Data
	collections, err := models.GetAllCollectionsByMemberID(c.Member.MemberId)
	if err != nil {
		logs.Error("Error AccountController GetCollections: ", err)
		c.JsonResult(1, err.Error())
	} else {
		for _, collection := range collections {
			job := models.GetJobByID(collection.JobID)
			res = append(res, Data{
				CollectionID: collection.CollectionID,
				Job:          job,
				CreateTime:   collection.CreateTime,
				Status:       collection.Status,
			})
		}

		c.JsonResult(0, "ok", res)
	}
}

// Edit 编辑用户基本信息
func (c *AccountController) Edit() {
	nickname := c.GetString("nickname")
	huntStatus := c.GetString("hunt_status")
	password := c.GetString("password")
	fmt.Println(nickname, huntStatus, password)

	member := models.NewMember()
	member.MemberId = c.Member.MemberId
	member.Nickname = nickname
	member.HuntStatus = huntStatus
	if password != "" {
		/*对密码进行加密*/
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			logs.Error("Error hash the password:", err)
		}
		encodePassword := string(hash)
		member.Password = encodePassword
		err = member.Update("nickname", "hunt_status", "password")
		if err != nil {
			logs.Error("Error AccountController Edit: ", err)
			c.JsonResult(1, err.Error())
		} else {
			m, _ := models.NewMember().Find(c.Member.MemberId)
			c.SetMember(*m)
			c.JsonResult(0, "ok")
		}

	} else {
		err := member.Update("nickname", "hunt_status")
		if err != nil {
			logs.Error("Error AccountController Edit: ", err)
			c.JsonResult(1, err.Error())
		} else {
			m, _ := models.NewMember().Find(c.Member.MemberId)
			c.SetMember(*m)
			c.JsonResult(0, "ok")
		}
	}

}

// GetMembers 获取所有用户信息
func (c *AccountController) GetMembers() {
	members, err := models.GetAllMembers()
	if err != nil {
		logs.Error("Error AccountController GetMembers: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok", members)
	}
}

// SetMemberStatus 设置用户状态
func (c *AccountController) SetMemberStatus() {
	memberID, _ := strconv.Atoi(c.GetString("member_id"))
	status, _ := strconv.Atoi(c.GetString("status"))

	var member = models.Member{
		MemberId: memberID,
		Status:   status,
	}
	err := member.Update("status")
	if err != nil {
		logs.Error("Error AccountController SetMemberStatus: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}

// SetMemberRole 设置用户权限
func (c *AccountController) SetMemberRole() {
	memberID, _ := strconv.Atoi(c.GetString("member_id"))
	role, _ := strconv.Atoi(c.GetString("role"))

	var member = models.Member{
		MemberId: memberID,
		Role:     role,
	}
	err := member.Update("role")
	if err != nil {
		logs.Error("Error AccountController SetMemberRole: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}
