package Controller

import (
	"Exe/Service"
	"Exe/Utils/ErrorMessage"
	rep "Exe/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建课堂
func CreateClass(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.CreateClass(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//删除课堂
func DeleteClass(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.DeleteClass(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//更改课堂信息
func UpdateClass(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.UpdateClass(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取课堂信息
func GetClass(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetClass(c)
	if code == emsg.Success {
		res.Success(*data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}

}

//学生加入课堂
func JoinClass(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.JoinClass(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//学生退出课堂
func OutClass(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.OutClass(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//发布作业or资料
func ReleaseMaterial(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.ReleaseMaterial(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//删除作业or资料
func DeleteMaterial(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.DeleteMaterial(c)
	if code == emsg.Success {
		res.Success(nil, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取作业or资料的信息
func GetMaterialMsgByClassId(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetMaterialMsgByClassId(c)
	if code == emsg.Success {
		res.Success(data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取作业or资料文件
func GetMaterial(c *gin.Context) {
	res := rep.NewResponseController(c)
	code := Service.GetMaterial(c)
	if code != emsg.Success {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}

//获取学生已加入课程
func GetClassByStudentId(c *gin.Context) {
	res := rep.NewResponseController(c)
	data, code := Service.GetClassByStudentId(c)
	if code == emsg.Success {
		res.Success(data, emsg.Success)
	} else {
		res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
	}
}
