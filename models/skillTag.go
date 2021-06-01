package models

import "github.com/astaxie/beego/orm"

// SkillTag 技能标签
type SkillTag struct {
	TagID   int    `orm:"pk;auto;column(tag_id)" json:"tag_id"` //标签ID
	Name    string `json:"name"`                                //标签名
	Type    string `json:"type"`                                //标签类型名
	Deleted int    `orm:"default(0)" json:"deleted"`            //删除标记
}

func (s *SkillTag) TableName() string {
	return TNSkillTag()
}

// GetAllSkillTags 获取所有的技能标签
func GetAllSkillTags() ([]SkillTag, error) {
	o := orm.NewOrm()
	var tags []SkillTag
	_, err := o.QueryTable(TNSkillTag()).Filter("deleted", 0).All(&tags)
	if err != nil {
		return nil, err
	} else {
		return tags, nil
	}
}

// InsertOrUpdate 添加或更新技能标签
func (s *SkillTag) InsertOrUpdate(fields ...string) error {
	o := orm.NewOrm()
	var tag SkillTag
	err := o.QueryTable(TNSkillTag()).Filter("tag_id", s.TagID).Filter("deleted", 0).One(&tag)
	if tag.TagID > 0 {
		_, err = o.Update(s, fields...)
	} else {
		_, err = o.Insert(s)
	}
	return err
}

// GetAllSkillTagTypes 获取所有的技能标签类型
func GetAllSkillTagTypes() ([]string, error) {
	o := orm.NewOrm()
	var types []string
	sql := "select distinct type from t_skill_tag where deleted = 0"
	_, err := o.Raw(sql).QueryRows(&types)
	if err != nil {
		return nil, err
	} else {
		return types, nil
	}
}

// GetAllSkillTagsByType 根据类型名获取所有的技能标签
func GetAllSkillTagsByType(tagType string) ([]SkillTag, error) {
	o := orm.NewOrm()
	var tags []SkillTag
	_, err := o.QueryTable(TNSkillTag()).Filter("type", tagType).Filter("deleted", 0).All(&tags)
	if err != nil {
		return nil, err
	} else {
		return tags, nil
	}

}
