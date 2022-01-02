package Controller

import (
	"Exe/Service"
	emsg "Exe/Utils/ErrorMessage"
	rep "Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建成绩
func CreateScore(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.CreateScore(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//更改成绩
func UpdateScore(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.UpdateScore(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取成绩信息
func GetScoreByStudentId(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetScoreByStudentId(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
