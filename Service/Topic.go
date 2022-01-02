package Service

import (
	"Exe/Model"
	emsg "Exe/Utils/ErrorMessage"
	gid "Exe/Utils/GenerateId"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var tpcIM = gid.NewIdMaker()

//创建话题
func CreateTopic(c *gin.Context) (*gin.H, int) {
	var topic Model.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	topic.Id = tpcIM.NewId()
	topic.ClassId, _ = strconv.Atoi(c.Param("class_id"))
	data, _ := c.Get("id")
	topic.TeacherId = data.(int)
	return &gin.H{
		"topic_id": topic.Id,
	}, Model.CreateTopic(&topic)
}

//删除话题
func DeleteTopic(c *gin.Context) int {
	topicId, err := strconv.Atoi(c.Param("topic_id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	data, _ := c.Get("id")
	teacherId := data.(int)
	return Model.DeleteTopic(topicId, teacherId)
}

//获取话题信息
func GetTopicByClassId(c *gin.Context) (*gin.H, int) {
	classId, err := strconv.Atoi(c.Param("class_id"))
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	topic, code := Model.GetTopic(classId)
	if code != emsg.Success {
		return nil, code
	} else {
		return &gin.H{
			"topic": topic,
		}, emsg.Success
	}
}

//更改话题信息
func UpdateTopic(c *gin.Context) int {
	topicId, err := strconv.Atoi(c.Param("topic_id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	var topic Model.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		return emsg.Error
	}
	return Model.UpdateTopic(topicId, &topic)
}
