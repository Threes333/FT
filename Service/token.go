package Service

import (
	ut "Exe/Utils"
	emsg "Exe/Utils/ErrorMessage"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) (*gin.H, int) {
	msg, _ := c.Get("id")
	id := msg.(int)
	username := c.GetString("username")
	identity := c.GetString("identity")
	actoken, err := ut.GenerateAccessToken(id, username, identity)
	if err != nil {
		return nil, emsg.GenerateAccessTokenFailed
	} else {
		return &gin.H{
			"access_token": actoken,
		}, emsg.Success
	}
}
