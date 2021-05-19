package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"opiece/server/global"
	"opiece/server/middleware"
	"opiece/server/model"
	service2 "opiece/server/service"
	"opiece/server/service/response"
	"time"
)

// Login 用户登录接口
// @Summary 用户登录接口
// @Description 用户登录，返回JWT
// @Tags 用户管理相关接口
// @Accept json
// @Param user query model.User true "用户登录信息"
// @Success 200 object model.User
// @Router /user/login [post]
func Login(c *gin.Context) {
	var L model.User
	_ = c.ShouldBind(&L)
	var user model.User
	user.Username = L.Username
	user.Password = L.Password
	if userRet, err := service2.Login(&user); err != nil {
		response.FailWithDetailed(nil, "登录失败，用户名不存在或密码错误", c)
	} else {
		userRet.Password = ""
		tokenNext(c, *userRet)
	}
}

func Register(c *gin.Context) {
	var R model.User
	_ = c.ShouldBind(&R)
	user := model.User{
		Username: R.Username,
		Password: R.Password,
		Email:    R.Email,
	}
	userRet, err := service2.Register(user)
	if err != nil {
		response.FailWithDetailed(nil, "注册用户失败", c)
		c.Abort()
	} else {
		userRet.Password = ""
		response.OkWithDetailed(userRet, "注册成功", c)
	}

}

func tokenNext(c *gin.Context, user model.User) {
	j := middleware.NewJWT() // 唯一签名
	claims := model.CustomClaims{
		ID:         user.ID,
		Username:   user.Username,
		BufferTime: global.OConfig.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.OConfig.JWT.ExpireTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                          // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithDetailed(nil, "获取token失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"user":       user,
		"token":      token,
		"expires_at": claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}
