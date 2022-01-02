package Service

import (
	"Exe/Model"
	emsg "Exe/Utils/ErrorMessage"
	gid "Exe/Utils/GenerateId"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var repIM = gid.NewIdMaker()

//创建回复
func CreateReply(c *gin.Context) (*gin.H, int) {
	var reply Model.Reply
	if err := c.ShouldBindJSON(&reply); err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	reply.TopicId, _ = strconv.Atoi(c.Param("topic_id"))
	reply.Id = repIM.NewId()
	data, _ := c.Get("id")
	reply.StudentId = data.(int)
	return &gin.H{
		"reply_id": reply.Id,
	}, Model.CreateReply(&reply)
}

//删除回复
func DeleteReply(c *gin.Context) int {
	replyId, err := strconv.Atoi(c.Param("reply_id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	data, _ := c.Get("id")
	studentId := data.(int)
	return Model.DeleteReply(replyId, studentId)
}

//获取回复信息
func GetReplyByStudentId(c *gin.Context) (*gin.H, int) {
	data, _ := c.Get("id")
	studentId := data.(int)
	reply, code := Model.GetReply(studentId)
	if code != emsg.Success {
		return nil, code
	} else {
		return &gin.H{
			"reply": reply,
		}, emsg.Success
	}
}
