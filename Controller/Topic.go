package Controller

import (
	"Exe/Service"
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建话题
func CreateTopic(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.CreateTopic(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//删除话题
func DeleteTopic(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.DeleteTopic(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//更新话题信息
func UpdateTopic(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.UpdateTopic(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取话题信息
func GetTopic(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetTopicByClassId(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
