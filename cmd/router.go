package cmd

import (
	"Exe/Controller"
	ut "Exe/Utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	//用户相关
	r.POST("/register", Controller.Register)
	r.GET("/login", Controller.Login)
	r.GET("/token", ut.RefreshTokenAuth(), Controller.RefreshToken)
	user := r.Group("/")
	user.Use(ut.AccessTokenAuth(""))
	{
		//用户相关
		r.DELETE("user", Controller.Logout)
		//课堂相关
		user.GET("class/:class_id", Controller.GetClass)
		user.GET("class/:class_id/material", Controller.GetMaterialMsgByClassId)
		user.GET("class/material/:id", Controller.GetMaterial)
		user.GET("class/:class_id/topic", Controller.GetTopic)
		user.POST("class/topic/:topic_id/reply", Controller.CreateReply)
		user.DELETE("class/topic/reply/:reply_id", Controller.DeleteReply)
	}
	teacher := r.Group("/teacher/")
	teacher.Use(ut.AccessTokenAuth("teacher"))
	{
		//课堂相关
		teacher.POST("class", Controller.CreateClass)
		teacher.DELETE("class/:class_id", Controller.DeleteClass)
		teacher.PUT("class/:class_id", Controller.UpdateClass)
		teacher.POST("class/:class_id/material/", Controller.ReleaseMaterial)
		teacher.DELETE("class/material/:id", Controller.DeleteMaterial)
		teacher.POST("class/:class_id/attendance", Controller.CreateAttendance)
		teacher.GET("class/attendance/:attendance_id", Controller.GetAttendanceStatus)
		teacher.POST("class/:class_id/topic", Controller.CreateTopic)
		teacher.DELETE("class/topic/:topic_id", Controller.DeleteTopic)
		teacher.PUT("class/topic/:topic_id", Controller.UpdateTopic)
		teacher.POST("class/:class_id/score", Controller.CreateScore)
		teacher.PUT("class/score/:score_id", Controller.UpdateScore)
	}
	student := r.Group("/student/")
	student.Use(ut.AccessTokenAuth("student"))
	{
		student.POST("class/:class_id", Controller.JoinClass)
		student.DELETE("class/:class_id", Controller.OutClass)
		student.GET("/class/", Controller.GetClassByStudentId)
		student.PUT("/class/attendance/:attendance_id", Controller.StudentCheckIn)
		student.GET("/class/topic/reply", Controller.GetReplyByStudentId)
		student.GET("/score", Controller.GetScoreByStudentId)
	}
	_ = r.Run(":8080")
}
