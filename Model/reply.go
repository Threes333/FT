package Model

import (
	"Exe/Utils/ErrorMessage"
	"gorm.io/gorm"
	"log"
)

type Reply struct {
	Id        int    `gorm:"type:int;primary_key"`
	TopicId   int    `json:"topic_id" gorm:"type:int"`
	StudentId int    `gorm:"type:varchar(200)"`
	Content   string `json:"content" gorm:"type:longtext"`
}

//创建回复
func CreateReply(reply *Reply) int {
	if err := DB.Create(reply).Error; err != nil {
		log.Println(err)
		return emsg.CreateReplyFailed
	}
	return emsg.Success
}

//删除回复
func DeleteReply(replyId int, studentId int) int {
	if err := DB.Where("id = ? and student_id = ?", replyId, studentId).Delete(Reply{}).Error; err != nil {
		log.Println(err)
		return emsg.DeleteReplyFailed
	}
	return emsg.Success
}

//获取回复信息
func GetReply(studentId int) ([]Reply, int) {
	var reply []Reply
	if err := DB.Debug().Where("student_id = ?", studentId).Find(&reply).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.ReplyNoExist
	} else if err != nil {
		log.Println(err)
		return nil, emsg.GetReplyFailed
	}
	return reply, emsg.Success
}
