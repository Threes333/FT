package Service

import (
	"Exe/Model"
	emsg "Exe/Utils/ErrorMessage"
	gid "Exe/Utils/GenerateId"
	"github.com/gin-gonic/gin"
	"strconv"
)

var AtdIM = gid.NewIdMaker()

//创建考勤
func CreateAttendance(c *gin.Context) (*gin.H, int) {
	classId, _ := strconv.Atoi(c.Param("class_id"))
	attendanceId := AtdIM.NewId()
	if code := Model.CreateAttendance(classId, attendanceId); code != emsg.Success {
		return nil, code
	} else {
		return &gin.H{
			"attendance_id": attendanceId,
		}, emsg.Success
	}
}

//学生签到
func StudentCheckIn(c *gin.Context) int {
	data, _ := c.Get("id")
	studentId := data.(int)
	attendanceId, _ := strconv.Atoi(c.Param("attendance_id"))
	return Model.StudentCheckIn(attendanceId, studentId)
}

//获取考勤情况
func GetAttendanceStatus(c *gin.Context) (*gin.H, int) {
	attendanceId, _ := strconv.Atoi(c.Param("attendance_id"))
	if status, code := Model.GetAttendanceStatus(attendanceId); code != emsg.Success {
		return nil, code
	} else {
		return &gin.H{
			"attendance_status": status,
		}, emsg.Success
	}
}
