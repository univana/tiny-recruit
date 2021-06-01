package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Job struct {
	JobID             int       `orm:"pk;auto;column(job_id)" json:"job_id"`           //职位ID
	Title             string    `json:"title"`                                         //职位标题
	Description       string    `orm:"type(text)" json:"description"`                  //职位描述
	Location          string    `json:"location"`                                      //职位工作地址
	SkillTags         string    `json:"skill_tags"`                                    //职位的技能标签
	MinMonthlySalary  int       `json:"min_monthly_salary"`                            //最低月薪
	MaxMonthlySalary  int       `json:"max_monthly_salary"`                            //最高月薪
	PayTimes          int       `json:"pay_times"`                                     //每年发薪次数
	RequireEducation  string    `json:"require_education"`                             //学历要求
	RequireExperience string    `json:"require_experience"`                            //经验要求
	Department        string    `json:"department"`                                    //职位所属部门
	Type              string    `json:"type"`                                          //职位类型
	Nature            string    `json:"nature"`                                        //职位性质
	Status            int       `orm:"default(0)" json:"status"`                       //职位状态 0:正常 1:删除 2:禁用
	CreateTime        time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //职位创建时间
	ModifyTime        time.Time `orm:"type(datetime);null" json:"modify_time"`         //最后修改时间

	Enterprise *Enterprise `orm:"rel(fk)" json:"enterprise"` //企业和职位的一对多关系
}

func (j *Job) TableName() string {
	return TNJob()
}

// GetAllJobs 获取所有职位信息
func GetAllJobs() ([]Job, error) {
	var jobs []Job
	var err error
	o := orm.NewOrm()
	_, err = o.QueryTable(TNJob()).Exclude("status", 1).All(&jobs)

	//同步对应企业数据
	for i := 0; i < len(jobs); i++ {
		_, err = o.LoadRelated(&jobs[i], "Enterprise")
	}

	return jobs, err
}

// GetJobByID 根据ID获取职位信息
func GetJobByID(id int) Job {
	o := orm.NewOrm()
	var job Job
	err := o.QueryTable(TNJob()).Filter("job_id", id).One(&job)
	if err != nil {
		logs.Error("Error get job: ", err)
		return Job{}
	}

	//同步对应企业数据
	_, err = o.LoadRelated(&job, "Enterprise")
	if err != nil {
		logs.Error("Error get enterprise for the job: ", err)
		return Job{}
	}
	return job
}

// IsDelivered 判断职位是否被投递
func IsDelivered(jobID int, memberID int) int {
	o := orm.NewOrm()
	var deliverance Deliverance
	err := o.QueryTable(TNDeliverance()).Filter("job_id", jobID).Filter("member_id", memberID).Exclude("status", "已撤销").One(&deliverance)
	if err == orm.ErrNoRows {
		return 0
	} else {
		return 1
	}
}

// FilterJobs 过滤职位信息
func FilterJobs(searchContent string, city string, requireExp string, requireEdu string, financingStage string, scale string, enterpriseType string) ([]Job, error) {
	o := orm.NewOrm()
	var jobs []Job
	var err error
	var jobsFiltered []Job
	qs := o.QueryTable(TNJob())
	cond := orm.NewCondition()
	if len(searchContent) != 0 {
		cond = cond.And("title__icontains", searchContent)
	}

	//按城市过滤
	if city != "不限" {
		cond = cond.And("location", city)
	}

	//工作经验过滤
	if requireExp != "不限" {
		cond = cond.And("require_experience", requireExp)
	}

	//按学历要求过滤
	if requireEdu != "不限" {
		cond = cond.And("require_education", requireEdu)
	}

	_, err = qs.SetCond(cond).Exclude("status", 1).All(&jobs)

	//针对企业数据筛选职位
	var enterpriseIDs = map[int]int{}
	enterprises, err := FilterEnterprises("不限", financingStage, scale, enterpriseType)
	if err != nil {
		logs.Error("Error job filter enterprises: ", err)
	}
	for n, enterprise := range enterprises {
		enterpriseIDs[enterprise.EnterpriseID] = n
	}

	//同步对应企业数据
	for i := 0; i < len(jobs); i++ {
		_, err = o.LoadRelated(&jobs[i], "Enterprise")
		if _, ok := enterpriseIDs[jobs[i].Enterprise.EnterpriseID]; ok {
			jobsFiltered = append(jobsFiltered, jobs[i])
		}
	}
	return jobsFiltered, err
}

// InsertOrUpdate 添加或更新职位信息
func (j *Job) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var job Job
	err := o.QueryTable(TNJob()).Filter("job_id", j.JobID).One(&job)
	if job.JobID > 0 {
		_, err = o.Update(j, fields...)
	} else {
		_, err = o.Insert(j)
	}
	return err
}

// IsCollected 判断职位是否被该用户收藏
func IsCollected(jobID int, memberID int) int {
	o := orm.NewOrm()
	var collection Collection
	err := o.QueryTable(TNCollection()).Filter("job_id", jobID).Filter("member_id", memberID).Filter("status", 0).One(&collection)
	if err == orm.ErrNoRows {
		return 0
	} else {
		return 1
	}
}
