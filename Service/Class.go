package Service

import (
	"Exe/Model"
	"Exe/Utils/ErrorMessage"
	"Exe/Utils/GenerateId"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const FileNameSplitChar = " "

var ClsIM = gid.NewIdMaker()

//创建课堂
func CreateClass(c *gin.Context) (*Model.Class, int) {
	var class Model.Class
	class.ClassName = c.PostForm("class_name")
	class.BeginTime, _ = time.Parse("2006-01-02 15:04:05", c.PostForm("begin_time"))
	class.EndTime, _ = time.Parse("2006-01-02 15:04:05", c.PostForm("end_time"))
	class.Id = ClsIM.NewId()
	data, _ := c.Get("id")
	class.TeacherId = data.(int)
	class.TeacherName = c.GetString("username")
	return &class, Model.CreateClass(&class)
}

//删除课堂
func DeleteClass(c *gin.Context) int {
	id, _ := strconv.Atoi(c.Param("class_id"))
	data, _ := c.Get("id")
	teacherId := data.(int)
	return Model.DeleteClass(id, teacherId)
}

//更改课堂信息
func UpdateClass(c *gin.Context) int {
	var class Model.Class
	class.ClassName = c.PostForm("class_name")
	class.BeginTime, _ = time.Parse("2006-01-02 15:04:05", c.PostForm("begin_time"))
	class.EndTime, _ = time.Parse("2006-01-02 15:04:05", c.PostForm("end_time"))
	class.Id, _ = strconv.Atoi(c.Param("class_id"))
	data, _ := c.Get("id")
	class.TeacherId = data.(int)
	class.TeacherName = c.GetString("username")
	return Model.UpdateClass(&class)
}

//获取课堂信息
func GetClass(c *gin.Context) (*Model.Class, int) {
	id, _ := strconv.Atoi(c.Param("class_id"))
	return Model.GetClass(id)

}

//学生加入课堂
func JoinClass(c *gin.Context) int {
	msg, _ := c.Get("id")
	studentId := msg.(int)
	classId, _ := strconv.Atoi(c.Param("class_id"))
	return Model.JoinClass(studentId, classId)
}

//学生退出课堂
func OutClass(c *gin.Context) int {
	msg, _ := c.Get("id")
	studentId := msg.(int)
	classId, _ := strconv.Atoi(c.Param("class_id"))
	return Model.OutClass(studentId, classId)
}

//发布作业or资料
func ReleaseMaterial(c *gin.Context) int {
	strId := c.Param("class_id")
	classId, _ := strconv.Atoi(strId)
	title := c.PostForm("title")
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	nowTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	path := ".\\material\\" + strId + FileNameSplitChar + nowTime + FileNameSplitChar + file.Filename
	if err = c.SaveUploadedFile(file, path); err != nil {
		log.Println(err)
		return emsg.Error
	}
	return Model.ReleaseMaterial(classId, path, title)
}

//删除作业or资料
func DeleteMaterial(c *gin.Context) int {
	var path string
	var code int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return emsg.Error
	}
	if path, code = Model.DeleteMaterial(id); code != emsg.Success {
		return code
	}
	if err = os.Remove(path); err != nil {
		log.Println(err)
		return emsg.Error
	}
	return code
}

//获取作业or资料的信息
func GetMaterialMsgByClassId(c *gin.Context) ([]Model.ClassMaterial, int) {
	classId, _ := strconv.Atoi(c.Param("class_id"))
	return Model.GetMaterialByClassId(classId)
}

//获取作业or资料文件
func GetMaterial(c *gin.Context) int {
	id, _ := strconv.Atoi(c.Param("id"))
	if path, code := Model.GetMaterial(id); code != emsg.Success {
		return code
	} else {
		file, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return emsg.Error
		}
		defer file.Close()
		fileName := strings.Split(file.Name(), FileNameSplitChar)[2]
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		c.File(path)
		return emsg.Success
	}
}

//获取学生已加入课程
func GetClassByStudentId(c *gin.Context) ([]Model.Class, int) {
	msg, _ := c.Get("id")
	studentId := msg.(int)
	return Model.GetClassByStudentId(studentId)
}
