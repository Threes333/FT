package Utils

import (
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/Response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	AccessTokenExpireDuration  = time.Hour * 2
	RefreshTokenExpireDuration = time.Hour * 24
)

var (
	AccessTokenKey  string = "Default AccessTokenKey"
	RefreshTokenKey string = "Default RefreshTokenKey"
)

type Jwt struct {
	SignKey []byte
}

type MyClaims struct {
	Id       int
	UserName string
	Identity string
	jwt.StandardClaims
}

func NewAccessTokenJwt() *Jwt {
	return &Jwt{
		[]byte(AccessTokenKey),
	}
}

func NewRefreshTokenJwt() *Jwt {
	return &Jwt{
		[]byte(RefreshTokenKey),
	}
}

func (j *Jwt) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

func (j *Jwt) ParseToken(token string) (*MyClaims, int) {
	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
	if err != nil {
		if er, ok := err.(*jwt.ValidationError); ok {
			if er.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, emsg.TokenErrorMalformed
				// ValidationErrorExpired表示Token过期
			} else if er.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, emsg.TokenErrorExpired
				// ValidationErrorNotValidYet表示无效token
			} else if er.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, emsg.TokenErrorNotValidYet
			} else {
				return nil, emsg.TokenCannotRecognized
			}

		}
	}
	if claim, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claim, emsg.Success
	}
	return nil, emsg.TokenInvalid
}

//identity 为空则所有身份都可通过,否则只有为identity的身份才能通过
func AccessTokenAuth(identity string) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := rep.NewResponseController(c)
		token := c.Request.Header.Get("token")
		if token == "" {
			res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(emsg.AccessTokenNoExist), emsg.AccessTokenNoExist)
			c.Abort()
			return
		}
		j := NewAccessTokenJwt()
		claims, code := j.ParseToken(token)
		if code != emsg.Success {
			res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
			c.Abort()
			return
		} else if identity != "" && identity != claims.Identity {
			res.Fail(http.StatusForbidden, emsg.GetErrorMsg(emsg.TokenNoPermission), emsg.TokenNoPermission)
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
		c.Set("username", claims.UserName)
		c.Set("identity", claims.Identity)
	}
}

func RefreshTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := rep.NewResponseController(c)
		token := c.Request.Header.Get("token")
		if token == "" {
			res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(emsg.RefreshTokenNoExist), emsg.RefreshTokenNoExist)
			c.Abort()
			return
		}
		j := NewRefreshTokenJwt()
		claims, code := j.ParseToken(token)
		if code != emsg.Success {
			res.Fail(http.StatusBadRequest, emsg.GetErrorMsg(code), code)
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
		c.Set("username", claims.UserName)
		c.Set("identity", claims.Identity)
	}
}

func GenerateAccessToken(id int, username string, identity string) (string, error) {
	j := NewAccessTokenJwt()
	claims := MyClaims{
		id,
		username,
		identity,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),
			Issuer:    "Threes",
		},
	}
	return j.CreateToken(claims)
}

func GenerateRefreshToken(id int, username string, identity string) (string, error) {
	j := NewRefreshTokenJwt()
	claims := MyClaims{
		id,
		username,
		identity,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(),
			Issuer:    "Threes",
		},
	}
	return j.CreateToken(claims)
}
