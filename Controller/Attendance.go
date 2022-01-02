package Controller

import (
	"Exe/Service"
	emsg "Exe/Utils/ErrorMessage"
	rep "Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建考勤
func CreateAttendance(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.CreateAttendance(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//学生签到
func StudentCheckIn(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.StudentCheckIn(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取考勤情况
func GetAttendanceStatus(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetAttendanceStatus(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
