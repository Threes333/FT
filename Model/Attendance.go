package Model

import (
	emsg "Exe/Utils/ErrorMessage"
	"gorm.io/gorm"
	"log"
)

type Attendance struct {
	Id      int      `json:"id" gorm:"type:int;primary key"`
	ClassId int      `json:"class_id" gorm:"type:int"`
	Status  []Status `gorm:"foreignkey:AttendanceId"`
}

type Status struct {
	Id           int `json:"id" gorm:"type:int;primary key;auto_increment"`
	AttendanceId int `gorm:"type:int"`
	StudentId    int `gorm:"type:int"`
	IsAttend     int `gorm:"type:varchar(200);default:0"`
}

//创建考勤
func CreateAttendance(classId int, attendanceId int) int {
	var studentIds []int
	var status []Status
	if err := DB.Model(&ClassMember{}).Select("student_id").Where("class_id = ?", classId).Find(&studentIds).Error; err == gorm.ErrRecordNotFound {
		return emsg.ClassHasNoStudent
	} else if err != nil {
		log.Println(err)
		return emsg.CreateAttendanceFailed
	}
	for _, id := range studentIds {
		status = append(status, Status{
			StudentId: id,
			IsAttend:  0,
		})
	}
	if err := DB.Create(&Attendance{Id: attendanceId, ClassId: classId, Status: status}).Error; err != nil {
		log.Println(err)
		return emsg.CreateAttendanceFailed
	} else {
		return emsg.Success
	}
}

//学生签到
func StudentCheckIn(attendanceId int, studentId int) int {
	var msg = map[string]interface{}{
		"IsAttend": 1,
	}
	if err := DB.Debug().Model(Status{}).Where("attendance_id = ? and student_id = ?", attendanceId, studentId).Updates(&msg).Error; err != nil {
		return emsg.CheckInFailed
	}
	return emsg.Success
}

//获取考勤情况
func GetAttendanceStatus(attendanceId int) (*Attendance, int) {
	var status Attendance
	if err := DB.Preload("Status").Where("id = ?", attendanceId).Find(&status).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.AttendanceNoExist
	} else if err != nil {
		return nil, emsg.GetAttendanceStatusFailed
	}
	return &status, emsg.Success
}
