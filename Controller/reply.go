package Controller

import (
	"Exe/Service"
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建回复
func CreateReply(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.CreateReply(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//删除回复
func DeleteReply(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.DeleteReply(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取回复信息
func GetReplyByStudentId(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetReplyByStudentId(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
