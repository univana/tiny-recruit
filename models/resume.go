package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Resume 简历模型
type Resume struct {
	ResumeId   int       `orm:"pk;auto" json:"resume_id"`                       //简历ID
	MemberId   int       ` json:"member_id"`                                    //用户ID
	Name       string    `orm:"size(30)" json:"name"`                           //姓名
	Gender     int       `json:"gender"`                                        //性别
	Birthday   time.Time `orm:"type(datetime)" json:"birthday"`                 //生日
	Photo      string    `json:"photo"`                                         //照片
	Tel        string    `json:"tel"`                                           //手机号
	Email      string    `json:"email"`                                         //邮箱
	Advantage  string    `orm:"type(text)" json:"advantage"`                    //个人优势
	HopeJob    string    `json:"hope_job"`                                      //期望职位
	HopeSalary string    `json:"hope_salary"`                                   //期望薪资
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	ModifyTime time.Time `orm:"type(datetime);auto_now_add" json:"modify_time"` //修改时间
}

//TableName :return-笔记本表名
func (m *Resume) TableName() string {
	return TNResume()
}

//NewDiary :return-日记
func NewDiary() *Resume {
	return &Resume{}
}

// GetResumeByMemberID 根据用户ID获得简历
func GetResumeByMemberID(memberID int) (Resume, error) {
	o := orm.NewOrm()
	var resume Resume
	//查找数据库
	err := o.QueryTable(TNResume()).Filter("member_id", memberID).One(&resume)
	return resume, err
}
