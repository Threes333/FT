package Service

import (
	"Exe/Model"
	ut "Exe/Utils"
	emsg "Exe/Utils/ErrorMessage"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(c *gin.Context) (*gin.H, int) {
	var user Model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	if code := Model.Register(&user); code != emsg.Success {
		return nil, code
	} else if actoken, err := ut.GenerateAccessToken(user.Id, user.UserName, user.Identity); err != nil {
		return nil, emsg.GenerateAccessTokenFailed
	} else if rftoken, err := ut.GenerateRefreshToken(user.Id, user.UserName, user.Identity); err != nil {
		return nil, emsg.GenerateRefreshTokenFailed
	} else {
		res := &gin.H{
			"user":          user,
			"access-token":  actoken,
			"refresh-token": rftoken,
		}
		return res, code
	}
}

func Login(c *gin.Context) (*gin.H, int) {
	username := c.Query("username")
	password := c.Query("password")
	if user, code := Model.Login(username, password); code != emsg.Success {
		return nil, code
	} else if actoken, err := ut.GenerateAccessToken(user.Id, user.UserName, user.Identity); err != nil {
		return nil, emsg.GenerateAccessTokenFailed
	} else if rftoken, err := ut.GenerateRefreshToken(user.Id, user.UserName, user.Identity); err != nil {
		return nil, emsg.GenerateRefreshTokenFailed
	} else {
		res := &gin.H{
			"access-token":  actoken,
			"refresh-token": rftoken,
		}
		return res, code
	}
}

//用户注销
func Logout(c *gin.Context) int {
	if data, ok := c.Get("id"); !ok {
		return emsg.Error
	} else {
		id, _ := data.(int)
		code := Model.Logout(id)
		return code
	}
}
