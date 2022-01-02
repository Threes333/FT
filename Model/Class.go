package Model

import (
	emsg "Exe/Utils/ErrorMessage"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type Class struct {
	Id          int       `json:"id" gorm:"type:int;primary key"`
	ClassName   string    `json:"class_name" gorm:"type:varchar(200);not null;"`
	TeacherId   int       `json:"teacher_id"`
	TeacherName string    `json:"teacher_name" gorm:"type:varchar(200);"`
	BeginTime   time.Time `json:"begin_time"`
	EndTime     time.Time `json:"end_time"`
}

type ClassMember struct {
	Class     Class `json:"class" gorm:"foreign key:ClassId"`
	ClassId   int   `json:"class_id" gorm:"type:int"`
	Student   User  `json:"student" gorm:"foreign key:StudentId"`
	StudentId int   `json:"student_id" gorm:"type:int"`
}

type ClassMaterial struct {
	Id           int    `json:"id" gorm:"type:int;primary key;auto_increment"`
	Class        Class  `json:"class" gorm:"foreign key:ClassId"`
	ClassId      int    `json:"class_id" gorm:"type:int"`
	MaterialName string `json:"material_name" gorm:"type:varchar(25)"`
	Material     string `json:"material" gorm:"type:varchar(200)"`
}

//创建课堂
func CreateClass(class *Class) int {
	if err := DB.Model(&Class{}).Create(class).Error; err != nil {
		return emsg.CreateClassFailed
	} else {
		return emsg.Success
	}
}

//删除课堂
func DeleteClass(classId int, teacherId int) int {
	if err := DB.Debug().Model(&Class{}).Where("id = ? and teacher = ?", classId, teacherId).Delete(&Class{}).Error; err != nil {
		return emsg.DeleteClassFailed
	} else {
		return emsg.Success
	}
}

//更改课堂信息
func UpdateClass(class *Class) int {
	var msg = map[string]interface{}{
		"ClassName":   class.ClassName,
		"TeacherId":   class.TeacherId,
		"TeacherName": class.TeacherName,
		"BeginTime":   class.BeginTime,
		"EndTime":     class.EndTime,
	}
	if err := DB.Debug().Model(Class{}).Where("id = ?", class.Id).Updates(&msg).Error; err != nil {
		return emsg.UpdateClassFailed
	}
	return emsg.Success
}

//获取课堂信息
func GetClass(id int) (*Class, int) {
	var class Class
	if err := DB.Where("id = ?", id).First(&class).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.ClassNoExist
	} else if err != nil {
		return nil, emsg.GetClassFailed
	}
	return &class, emsg.Success
}

//学生加入课堂
func JoinClass(studentId, classId int) int {
	var msg = ClassMember{
		ClassId:   classId,
		StudentId: studentId,
	}
	if err := DB.Model(&ClassMember{}).Create(&msg).Error; err != nil {
		return emsg.JoinClassFailed
	} else {
		return emsg.Success
	}
}

//学生退出课堂
func OutClass(studentId, classId int) int {
	if err := DB.Debug().Model(&ClassMember{}).Where("student_id = ? and class_id = ?", studentId, classId).Delete(&ClassMember{}).Error; err != nil {
		return emsg.OutClassFailed
	} else {
		return emsg.Success
	}
}

//发布作业or资料
func ReleaseMaterial(classId int, path, title string) int {
	var material = ClassMaterial{
		ClassId:      classId,
		Material:     path,
		MaterialName: title,
	}
	if err := DB.Model(&ClassMaterial{}).Create(&material).Error; err != nil {
		log.Println(err)
		return emsg.ReleaseMaterialFailed
	} else {
		return emsg.Success
	}
}

//删除作业or资料
func DeleteMaterial(id int) (string, int) {
	var path string
	if err := DB.Model(&ClassMaterial{}).Select("material").Where("id = ?", id).First(&path).Error; err != nil {
		log.Println(err)
		return "", emsg.MaterialNoExist
	}
	if err := DB.Model(&ClassMaterial{}).Where("id = ?", id).Delete(&ClassMaterial{}).Error; err != nil {
		log.Println(err)
		return "", emsg.DeleteMaterialFailed
	} else {
		return path, emsg.Success
	}
}

//获取作业or资料信息
func GetMaterialByClassId(classId int) ([]ClassMaterial, int) {
	var materials []ClassMaterial
	if err := DB.Debug().Preload("Class").Where("class_id = ?", classId).Find(&materials).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.MaterialNoExist
	} else if err != nil {
		log.Println(err)
		return nil, emsg.GetMaterialFailed
	}
	return materials, emsg.Success
}

//获取作业or资料路径
func GetMaterial(id int) (string, int) {
	var path string
	if err := DB.Model(&ClassMaterial{}).Select("material").Where("id = ?", id).First(&path).Error; err == gorm.ErrRecordNotFound {
		return "", emsg.MaterialNoExist
	} else if err != nil {
		log.Println(err)
		return "", emsg.GetMaterialFailed
	}
	return path, emsg.Success
}

//获取学生已加入的课程
func GetClassByStudentId(studentId int) ([]Class, int) {
	var classes []Class
	if err := DB.Debug().Model(&ClassMember{}).Joins("Class").Select("`Class`.`id`,`Class`.`class_name`,`Class`.`teacher_id`,`Class`.`teacher_name`,`Class`.`begin_time`,`Class`.`end_time`").Where("student_id = ?", studentId).Find(&classes).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.ClassNoExist
	} else if err != nil {
		log.Println(err)
		return nil, emsg.GetClassFailed
	}
	fmt.Println(classes)
	return classes, emsg.Success
}
