package Service

import (
	"Exe/Model"
	emsg "Exe/Utils/ErrorMessage"
	gid "Exe/Utils/GenerateId"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var scrIM = gid.NewIdMaker()

//创建成绩
func CreateScore(c *gin.Context) (*gin.H, int) {
	var score Model.Score
	if err := c.ShouldBindJSON(&score); err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	score.ClassId, _ = strconv.Atoi(c.Param("class_id"))
	score.Id = scrIM.NewId()
	data, _ := c.Get("id")
	score.TeacherId = data.(int)
	return &gin.H{
		"score_id": score.Id,
	}, Model.CreateScore(&score)
}

//更新成绩
func UpdateScore(c *gin.Context) int {
	var score Model.Score
	err := c.ShouldBindJSON(&score)
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	score.Id, _ = strconv.Atoi(c.Param("score_id"))
	data, _ := c.Get("id")
	score.TeacherId = data.(int)
	return Model.UpdateScore(&score)
}

//获取成绩信息
func GetScoreByStudentId(c *gin.Context) (*gin.H, int) {
	data, _ := c.Get("id")
	studentId := data.(int)
	score, code := Model.GetScoreByStudentId(studentId)
	if code != emsg.Success {
		return nil, code
	} else {
		return &gin.H{
			"score": score,
		}, emsg.Success
	}
}
