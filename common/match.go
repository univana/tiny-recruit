package common

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"myApp/models"
	"strconv"
	"strings"
)

const ()

var (
	eduMap = map[string]float64{
		"博士":    1,
		"硕士":    0.8,
		"本科":    0.6,
		"大专":    0.4,
		"高中及以下": 0.2,
	}
	locationMap = map[string]float64{
		"同城":   1,
		"同省跨城": 0.50,
		"跨省":   0,
	}
)

// matchEducation 计算教育经历匹配值
func matchEducation(resume models.Resume) float64 {
	//确定最高教育经历
	var res = 0.0
	for _, edu := range resume.EducationExperiences {
		tmp := eduMap[edu.Education]
		if tmp > res {
			res = tmp
		}
	}
	return res
}

// matchSkillTag 计算技能标签匹配度
func matchSkillTag(resume string, job string) float64 {
	resumeSkills := strings.Split(resume, "-")
	jobSkills := strings.Split(job, "-")

	//职位的技能标签向量
	x1 := make([]float64, len(jobSkills))
	// 简历的技能标签向量
	x2 := make([]float64, len(jobSkills))
	for i := 0; i < len(x1); i++ {
		x1[i] = 1.0
		for j := 0; j < len(resumeSkills); j++ {
			if resumeSkills[j] == jobSkills[i] {
				x2[i] = 1.0
			}
		}
	}
	//计算余弦相似度
	var (
		up    = 0.0
		lenX1 = 0.0
		lenX2 = 0.0
	)
	for i := 0; i < len(x1); i++ {
		up = up + x1[i]*x2[i]
		lenX1 = lenX1 + math.Pow(x1[i], 2)
		lenX2 = lenX2 + math.Pow(x2[i], 2)
	}
	lenX1 = math.Sqrt(lenX1)
	lenX2 = math.Sqrt(lenX2)

	cos := up / (lenX1 * lenX2)

	return cos
}

// isDirectCity 判断是否是直辖市
func isDirectCity(city string) bool {
	direct := []string{"北京市", "天津市", "上海市", "重庆市"}
	for _, d := range direct {
		if city == d {
			return true
		}
	}
	return false
}

// isInSameProvince 判断两个城市是否同省
func isInSameProvince(c1 string, c2 string) bool {
	o := orm.NewOrm()
	sql := "select pro_id from t_city where city_name = ?"
	var pro1 string
	var pro2 string
	err := o.Raw(sql, c1).QueryRow(&pro1)
	if err != nil {
		logs.Error("Error get pro_id:", err)

	}
	err = o.Raw(sql, c2).QueryRow(&pro2)
	if err != nil {
		logs.Error("Error get pro_id:", err)

	}
	return pro1 == pro2
}

// matchLocation 计算工作点匹配值
func matchLocation(l1 string, l2 string) float64 {
	//检测是否同城
	if l1 == l2 {
		return locationMap["同城"]
	} else if isDirectCity(l1) || isDirectCity(l2) {
		return locationMap["跨省"]
	} else if isInSameProvince(l1, l2) {
		return locationMap["同省跨城"]
	} else {
		return locationMap["跨省"]
	}
}

// matchSalary 计算薪水匹配度
func matchSalary(hope int, min int, max int) float64 {
	if hope < min {
		return 1.0
	} else if hope > max {
		return 0.0
	} else {
		return float64(hope-min) / float64(max-min)
	}
}

// GetMatchingDegree 计算简历与职位的匹配程度
func GetMatchingDegree(resume models.Resume, job models.Job, wEdu float64, wLoc float64, wSkill float64, wSal float64) float64 {
	var (
		matchingDegree float64 //待返回的匹配度结果
		mEducation     float64 //学历匹配度
		mLocation      float64 //工作地点匹配度
		mSkill         float64 // 技能匹配度
		mSalary        float64 //期望薪水匹配度
	)

	//空值检测
	if resume.EducationExperiences == nil {
		mEducation = 0
	} else {
		mEducation = matchEducation(resume)
	}

	if resume.City == "" {
		mLocation = 0
	} else {
		mLocation = matchLocation(resume.City, job.Location)
	}

	if resume.SkillTags == "" {
		mSkill = 0
	} else {
		mSkill = matchSkillTag(resume.SkillTags, job.SkillTags)
	}
	if resume.HopeSalary == "" {
		mSalary = 0
	} else {
		hopeSalary, _ := strconv.Atoi(resume.HopeSalary[:len(resume.HopeSalary)-1])
		mSalary = matchSalary(hopeSalary, job.MinMonthlySalary, job.MaxMonthlySalary)
	}
	matchingDegree = mEducation*wEdu + mLocation*wLoc + mSkill*wSkill + mSalary*wSal

	return matchingDegree
}
