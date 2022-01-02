package Model

import (
	emsg "Exe/Utils/ErrorMessage"
	"gorm.io/gorm"
	"log"
)

type Score struct {
	Id        int `gorm:"type:int;primary_key"`
	ClassId   int `gorm:"type:int"`
	TeacherId int `gorm:"type:int"`
	StudentId int `json:"student_id" gorm:"type:varchar(200)"`
	Score     int `json:"score" gorm:"type:longtext"`
}

//创建成绩
func CreateScore(score *Score) int {
	if err := DB.Create(score).Error; err != nil {
		log.Println(err)
		return emsg.CreateScoreFailed
	}
	return emsg.Success
}

//更新成绩
func UpdateScore(score *Score) int {
	var msg = map[string]interface{}{
		"score": score.Score,
	}
	if err := DB.Debug().Model(&Score{}).Where("id = ? and teacher_id = ?", score.Id, score.TeacherId).Updates(&msg).Error; err != nil {
		log.Println(err)
		return emsg.UpdateScoreFailed
	}
	return emsg.Success
}

//获取成绩
func GetScoreByStudentId(studentId int) ([]Score, int) {
	var score []Score
	if err := DB.Debug().Where("student_id = ?", studentId).Find(&score).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.ScoreNoExist
	} else if err != nil {
		log.Println(err)
		return nil, emsg.GetScoreFailed
	}
	return score, emsg.Success
}
