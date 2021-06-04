package controllers

import (
	"compress/gzip"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io"
	"myApp/common"
	"myApp/models"
	"strconv"
	"strings"
	"time"

	"myApp/utils"

	"github.com/astaxie/beego"
)

//BaseController :基本控制器类
type BaseController struct {
	beego.Controller
	Member *models.Member    //用户
	Option map[string]string //全局设置
}

//CookieRemember :Cookie记忆器
type CookieRemember struct {
	MemberId int
	Account  string
	Time     time.Time
}

//Prepare :每个子类Controller公用方法调用前，都执行一下Prepare方法
func (c *BaseController) Prepare() {
	c.Member = models.NewMember() //初始化
	//从session中获取用户信息
	if member, ok := c.GetSession(common.SessionName).(models.Member); ok && member.MemberId > 0 {
		c.Member = &member
	} else {
		//如果Cookie中存在登录信息，从cookie中获取用户信息
		if cookie, ok := c.GetSecureCookie(common.AppKey(), "login"); ok {
			var remember CookieRemember
			err := utils.Decode(cookie, &remember)
			if err == nil {
				member, err := models.NewMember().Find(remember.MemberId)
				if err == nil {
					c.SetMember(*member)
					c.Member = member
				}
			}
		}
	}
	c.Data["Member"] = c.Member

	notifies, err := models.GetAllNotifiesByMemberID(c.Member.MemberId)
	if err != nil {
		logs.Error("Error get notify:", err)
	}
	c.Data["Notifies"] = notifies
	c.Data["HasUnread"] = models.HasUnread(c.Member.MemberId)
	c.Data["BaseUrl"] = c.BaseUrl()
	c.Data["SITE_NAME"] = "tiny-recruit"
	//设置全局配置
	c.Option = make(map[string]string)
	c.Option["ENABLED_CAPTCHA"] = "false"
}

func (c *BaseController) BaseUrl() string {
	host := beego.AppConfig.String("sitemap_host")
	if len(host) > 0 {
		if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
			return host
		}
		return c.Ctx.Input.Scheme() + "://" + host
	}
	return c.Ctx.Input.Scheme() + "://" + c.Ctx.Request.Host
}

// 设置登录用户信息
func (c *BaseController) SetMember(member models.Member) {
	if member.MemberId <= 0 {
		c.DelSession(common.SessionName)
		c.DelSession("uid")
		c.DestroySession()
	} else {
		c.SetSession(common.SessionName, member)
		c.SetSession("uid", member.MemberId)
	}
}

// Ajax接口返回Json
func (c *BaseController) JsonResult(errCode int, errMsg string, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["errcode"] = errCode
	jsonData["message"] = errMsg

	if len(data) > 0 && data[0] != nil {
		jsonData["data"] = data[0]
	}
	returnJSON, err := json.Marshal(jsonData)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	//启用gzip压缩
	if strings.Contains(strings.ToLower(c.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
		c.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		w := gzip.NewWriter(c.Ctx.ResponseWriter)
		defer w.Close()
		w.Write(returnJSON)
		w.Flush()
	} else {
		io.WriteString(c.Ctx.ResponseWriter, string(returnJSON))
	}
	c.StopRun()
}

// GetCities 获取省份对应的所有城市信息
func (c *BaseController) GetCities() {
	provinceName := c.GetString("province_name")
	sql := "select city_id,city_name from t_city join t_province tp on t_city.pro_id = tp.pro_id where pro_name= ?"
	o := orm.NewOrm()
	type City struct {
		CityID   string `orm:"column(city_id)" json:"city_id"`
		CityName string `orm:"column(city_name)" json:"city_name"`
	}
	var cities []City
	_, err := o.Raw(sql, provinceName).QueryRows(&cities)
	if err != nil {
		logs.Error("Error BaseController GetCities: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok", cities)
	}
}

// GetSkillTags 获取所有的技能标签
func (c *BaseController) GetSkillTags() {
	tags, err := models.GetAllSkillTags()
	if err != nil {
		logs.Error("Error BaseController GetSkillTags: ", err)
	} else {
		c.JsonResult(0, "ok", tags)
	}
}

func (c *BaseController) DeleteSkillTag() {
	tagID, _ := strconv.Atoi(c.GetString("tag_id"))
	var tag = models.SkillTag{
		TagID:   tagID,
		Deleted: 1,
	}
	err := tag.InsertOrUpdate("deleted")
	if err != nil {
		logs.Error("Error Delete Tag: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}

// AddSkillTag 添加技能标签
func (c *BaseController) AddSkillTag() {
	var tag = models.SkillTag{
		TagID:   0,
		Name:    c.GetString("tag_name"),
		Type:    c.GetString("tag_type"),
		Deleted: 0,
	}
	err := tag.InsertOrUpdate()
	if err != nil {
		logs.Error("Error BaseController AddSkillTag: ", err)
		c.JsonResult(1, err.Error())
	} else {
		c.JsonResult(0, "ok")
	}
}
