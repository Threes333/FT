package Model

import (
	emsg "Exe/Utils/ErrorMessage"
	"gorm.io/gorm"
	"log"
)

type Topic struct {
	Id        int     `gorm:"type:int;primary key;"`
	ClassId   int     `gorm:"type:int"`
	TeacherId int     `gorm:"type:int"`
	Title     string  `json:"title" gorm:"type:varchar(50)"` //话题标题
	Content   string  `json:"content" gorm:"type:longtext"`  //话题内容
	Reply     []Reply `gorm:"foreign:TopicId"`               //话题回复
}

//创建话题
func CreateTopic(topic *Topic) int {
	if err := DB.Create(topic).Error; err != nil {
		log.Println(err)
		return emsg.CreateTopicFailed
	}
	return emsg.Success
}

//删除话题
func DeleteTopic(topicId int, teacherId int) int {
	if err := DB.Where("id = ? and teacher_id = ?", topicId, teacherId).Delete(&Topic{}).Error; err != nil {
		log.Println(err)
		return emsg.DeleteTopicFailed
	}
	return emsg.Success
}

//更改话题
func UpdateTopic(id int, topic *Topic) int {
	var msg = map[string]interface{}{
		"Title":   topic.Title,
		"Content": topic.Content,
	}
	if err := DB.Model(Topic{}).Where("id = ?", id).Updates(&msg).Error; err != nil {
		return emsg.UpdateTopicFailed
	}
	return emsg.Success
}

//获取话题信息
func GetTopic(classId int) ([]Topic, int) {
	var topic []Topic
	if err := DB.Where("class_id = ?", classId).Preload("Reply").First(&topic).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.TopicNoExist
	} else if err != nil {
		return nil, emsg.GetTopicFailed
	}
	return topic, emsg.Success
}
