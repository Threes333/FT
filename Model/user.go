package Model

import (
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/GenerateId"
	"gorm.io/gorm"
	"log"
)

var IM = gid.NewIdMaker()

type User struct {
	Id       int    `json:"id" gorm:"type:int; primary_key"`
	UserName string `json:"username" gorm:"type:varchar(25)"`
	PassWord string `json:"password" gorm:"type:varchar(25)"`
	Identity string `json:"identity" gorm:"type:varchar(25)"`
}

//用户注册
func Register(user *User) int {
	var id int
	DB.Model(User{}).Select("id").Where("user_name = ?", user.UserName).First(&id)
	if id == 0 {
		//该用户名可用
		user.Id = IM.NewId()
		DB.Create(user)
		return emsg.Success
	}
	return emsg.UsernameExist
}

//用户登录
func Login(username, password string) (*User, int) {
	var user User
	if err := DB.Model(User{}).Where("user_name = ?", username).First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.UsernameNoExist //判断用户是否存在
	} else if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}

	if user.PassWord == password { //判断密码是否正确
		return &user, emsg.Success
	}
	return nil, emsg.PasswordWrong
}

//用户注销
func Logout(id int) int {
	if err := DB.Model(User{}).Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return emsg.UserLogoutFailed
	}
	return emsg.Success
}
