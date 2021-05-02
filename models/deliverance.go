package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Deliverance struct {
	DeliveranceID int       `orm:"pk;auto;column(deliverance_id)" json:"deliverance_id"` //投递ID
	MemberID      int       `orm:"column(member_id)" json:"member_id"`                   //用户ID
	JobID         int       `orm:"column(job_id)" json:"job_id"`                         //职位ID
	DeliverTime   time.Time `orm:"type(datetime);null" json:"deliver_time"`              //投递时间
	ModifyTime    time.Time `orm:"type(datetime);null" json:"modify_time"`               //状态修改时间
	Status        string    `orm:"default(待查看)" json:"status"`                           //投递状态
}

func (d *Deliverance) TableName() string {
	return TNDeliverance()
}

func (d *Deliverance) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var deliverance Deliverance
	err := o.QueryTable(TNDeliverance()).Filter("deliverance_id", d.DeliveranceID).One(&deliverance)
	if deliverance.DeliveranceID > 0 {
		_, err = o.Update(d, fields...)
	} else {
		_, err = o.Insert(d)
	}
	return err
}

// GetAllDeliveranceByMemberID 根据用户ID获取拥有企业对应的所有投递信息
func GetAllDeliveranceByMemberID(memberID int) []Deliverance {
	o := orm.NewOrm()
	var delivers []Deliverance
	_, err := o.QueryTable(TNDeliverance()).Filter("member_id", memberID).All(&delivers)
	if err != nil {
		logs.Error("Error get all deliverance: ", err)
		return nil
	} else {
		return delivers
	}
}

// GetAllResumesByJobID 根据职位ID获取所有的投递简历数据
func GetAllResumesByJobID(jobID int) ([]Resume, error) {
	o := orm.NewOrm()
	//查询该职位投递的所有用户ID
	sql := "select member_id from t_deliverance join t_job tj on t_deliverance.job_id = tj.job_id where tj.job_id = ?"
	var ids []int
	_, err := o.Raw(sql, jobID).QueryRows(&ids)
	if err != nil {
		logs.Error("Error GetAllResumeByJobID: ", err)
		return nil, err
	}

	var resumes []Resume
	for _, id := range ids {
		resume, err := GetResumeByMemberID(id)
		if err != nil {
			logs.Error("Error GetResumeByMemberID: ", err)
			return nil, err
		}
		//为简历加载经历信息
		resume.LoadExperiences()
		resumes = append(resumes, resume)
	}
	return resumes, nil

}
