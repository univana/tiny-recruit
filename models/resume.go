package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

// Resume 简历模型
type Resume struct {
	ResumeID   int       `orm:"pk;auto;column(resume_id)" json:"resume_id"`     //简历ID
	Name       string    `orm:"size(30)" json:"name"`                           //姓名
	Gender     int       `json:"gender"`                                        //性别
	Birthday   time.Time `orm:"type(datetime)" json:"birthday"`                 //生日
	Photo      string    `json:"photo"`                                         //照片
	SkillTags  string    `json:"skill_tags"`                                    //技能标签
	Tel        string    `json:"tel"`                                           //手机号
	Email      string    `json:"email"`                                         //邮箱
	Advantage  string    `orm:"type(text)" json:"advantage"`                    //个人优势
	HopeJob    string    `json:"hope_job"`                                      //期望职位
	HopeSalary string    `json:"hope_salary"`                                   //期望薪资
	City       string    `json:"city"`                                          //城市
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	ModifyTime time.Time `orm:"type(datetime)" json:"modify_time"`              //修改时间

	Member                *Member                 `orm:"rel(one)"`                                    //简历所属用户 一对一关系
	EducationExperiences  []*EducationExperience  `orm:"reverse(many)" json:"education_experiences"`  //简历和教育经历的一对多关系
	ProjectExperiences    []*ProjectExperience    `orm:"reverse(many)" json:"project_experiences"`    //简历和项目经历的一对多关系
	InternshipExperiences []*InternshipExperience `orm:"reverse(many)" json:"internship_experiences"` //简历和实习经历的一对多关系
}

//TableName :return-笔记本表名
func (r *Resume) TableName() string {
	return TNResume()
}

// GetResumeByMemberID 根据用户ID获得简历
func GetResumeByMemberID(memberID int) (Resume, error) {
	o := orm.NewOrm()
	var resume Resume
	//查找数据库
	err := o.QueryTable(TNResume()).Filter("member_id", memberID).One(&resume)
	return resume, err
}

// LoadExperiences 为简历加载经历
func (r *Resume) LoadExperiences() {
	o := orm.NewOrm()
	//加载教育经历
	_, err := o.LoadRelated(r, "EducationExperiences")
	if err != nil {
		logs.Error("Error load experiences: ", err)
	}
	//加载实习经历
	_, err = o.LoadRelated(r, "InternshipExperiences")
	if err != nil {
		logs.Error("Error load experiences: ", err)
	}
	//加载项目经历
	_, err = o.LoadRelated(r, "ProjectExperiences")
	if err != nil {
		logs.Error("Error load experiences: ", err)
	}
}

// InsertOrUpdate 插入或更新简历
func (r *Resume) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var resume Resume
	err := o.QueryTable(TNResume()).Filter("resume_id", r.ResumeID).One(&resume)
	if resume.ResumeID > 0 {
		_, err = o.Update(r, fields...)
	} else {
		_, err = o.Insert(r)
	}
	return err
}
