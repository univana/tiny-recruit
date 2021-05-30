package models

import "github.com/astaxie/beego/orm"

type JobType struct {
	TypeID   int    `orm:"pk;auto;column(type_id)" json:"type_id"` //职位类型ID
	Name     string `json:"name"`                                  //类型名
	ParentID int    `orm:"column(parent_id)" json:"parent_id"`     //父标签ID
	Level    string `json:"level"`                                 //标签的层级
	Deleted  int    `orm:"default(0)" json:"deleted"`              //删除标记
}

func (j *JobType) TableName() string {
	return TNJobType()
}

// GetAllTypes 获取所有职位类型
func GetAllTypes() ([]JobType, error) {
	o := orm.NewOrm()
	var types []JobType
	_, err := o.QueryTable(TNJobType()).Filter("deleted", 0).All(&types)
	if err != nil {
		return nil, err
	} else {
		return types, nil
	}
}

// GetParentName 获取父类型名
func (j *JobType) GetParentName() (string, error) {
	if j.Level == "一级" {
		return "无", nil
	}
	o := orm.NewOrm()
	var jobType JobType
	err := o.QueryTable(TNJobType()).Filter("type_id", j.ParentID).One(&jobType)
	if err != nil {
		return "", err
	} else {
		return jobType.Name, nil
	}
}

// InsertOrUpdate 添加或更新职位类型
func (j *JobType) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var jobType JobType
	err := o.QueryTable(TNJobType()).Filter("type_id", j.TypeID).Filter("deleted", 0).One(&jobType)

	if jobType.TypeID > 0 {
		_, err = o.Update(j, fields...)
	} else {
		_, err = o.Insert(j)
	}
	return err
}

// GetTypeByID 根据ID获取职位类型数据
func GetTypeByID(typeID int) (JobType, error) {
	o := orm.NewOrm()
	var jobType JobType
	err := o.QueryTable(TNJobType()).Filter("type_id", typeID).Filter("deleted", 0).One(&jobType)
	if err != nil {
		return JobType{}, err
	} else {
		return jobType, nil
	}
}

// GetFirstLevelJobTypes 获取所有一级职位分类
func GetFirstLevelJobTypes() ([]JobType, error) {
	o := orm.NewOrm()
	var jobTypes []JobType
	_, err := o.QueryTable(TNJobType()).Filter("level", "一级").Filter("deleted", 0).All(&jobTypes)
	if err != nil {
		return nil, err
	} else {
		return jobTypes, nil
	}
}

// GetSecondLevelJobTypesByParentID 获取所有一级职位分类
func GetSecondLevelJobTypesByParentID(parentID int) ([]JobType, error) {
	o := orm.NewOrm()
	var jobTypes []JobType
	_, err := o.QueryTable(TNJobType()).Filter("level", "二级").Filter("deleted", 0).Filter("parent_id", parentID).All(&jobTypes)
	if err != nil {
		return nil, err
	} else {
		return jobTypes, nil
	}
}

// GetThirdLevelJobTypesByParentID 获取所有一级职位分类
func GetThirdLevelJobTypesByParentID(parentID int) ([]JobType, error) {
	o := orm.NewOrm()
	var jobTypes []JobType
	_, err := o.QueryTable(TNJobType()).Filter("level", "三级").Filter("deleted", 0).Filter("parent_id", parentID).All(&jobTypes)
	if err != nil {
		return nil, err
	} else {
		return jobTypes, nil
	}
}
