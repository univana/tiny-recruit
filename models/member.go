package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/astaxie/beego/orm"
)

//Member :用户模型
type Member struct {
	MemberId      int       `orm:"pk;auto" json:"member_id"`                       //用户ID
	Account       string    `orm:"size(30);unique" json:"account"`                 //账户
	Nickname      string    `json:"nickname"`                                      //昵称
	Password      string    `json:"password"`                                      //密码
	Avatar        string    `json:"avatar"`                                        //头像地址
	Role          int       `orm:"default(1)" json:"role"`                         //权限 1：普通用户 0：管理员
	Status        int       `orm:"default(0)" json:"status"`                       //状态 0：正常 1：禁用
	HuntStatus    string    `json:"hunt_status"`                                   //求职状态
	CreateTime    time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	LastLoginTime time.Time `orm:"type(datetime);null" json:"last_login_time"`     //最后登录时间
}

func (m *Member) TableName() string {
	return TNMembers()
}

func NewMember() *Member {
	return &Member{}
}

func (m *Member) Find(id int) (*Member, error) {
	m.MemberId = id
	if err := orm.NewOrm().Read(m); err != nil {
		return m, err
	}
	return m, nil
}

func (m *Member) Add() error {

	cond := orm.NewCondition().Or("nickname", m.Nickname).Or("account", m.Account)
	var one Member
	o := orm.NewOrm()
	if o.QueryTable(m.TableName()).SetCond(cond).One(&one, "member_id", "nickname", "account"); one.MemberId > 0 {
		if one.Nickname == m.Nickname {
			return errors.New("昵称已存在")
		}

		if one.Account == m.Account {
			return errors.New("用户已存在")
		}
	}

	if _, err := o.Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Member) Update(cols ...string) error {

	if _, err := orm.NewOrm().Update(m, cols...); err != nil {
		return err
	}
	return nil
}

//登录
func (m *Member) Login(account string, password string) (*Member, error) {
	//新建用户实例
	member := &Member{}
	//根据账户名查询用户信息
	err := orm.NewOrm().QueryTable(m.TableName()).Filter("account", account).Filter("status", 0).One(member)
	if err != nil {
		return member, errors.New("用户不存在")
	}

	/*验证密码*/
	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(password))
	if err != nil {
		return member, errors.New("密码错误")
	} else {
		//设置用户身份
		return member, nil
	}
}

//根据用户名获取用户信息
func (m *Member) GetByUsername(username string) (member Member, err error) {
	err = orm.NewOrm().QueryTable(TNMembers()).Filter("account", username).One(&member)
	return
}

func (m *Member) IsAdministrator() bool {
	if m == nil || m.MemberId <= 0 {
		return false
	}
	return m.Role == 0 || m.Role == 1
}
