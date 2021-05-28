package models

type JobType struct {
	TypeID   int    `orm:"pk;auto;column(type_id)" json:"type_id"` //职位类型ID
	Name     string `json:"name"`                                  //类型名
	ParentID int    `orm:"column(parent_id)" json:"parent_id"`     //父标签ID
	Deleted  int    `orm:"default(0)" json:"deleted"`              //删除标记
}

func (j *JobType) TableName() string {
	return TNJobType()
}
