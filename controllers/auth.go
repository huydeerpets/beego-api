package controllers

import (
	"beego-api/models"
	"beego-api/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	_ "github.com/astaxie/beego/validation"
)

// 认证
type AuthController struct {
	beego.Controller
}

// @Title 注册
// @Description 用户注册
// @Param  username  body  string  true  "手机号"
// @Param  password  body  string  true  "密码"
// @router /register [post]
func (a *AuthController) Register() {
	username := a.GetString("username")
	password := a.GetString("password")

	user := new(models.User)

	user.Username = username
	user.Password = password

	valid := validation.Validation{}
	valid.Required(user.Username, "username")
	valid.Required(user.Password, "password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			a.Data["json"] = &RespError{
				Code: paramError,
				Msg:  err.Message,
			}
			a.ServeJSON()
		}
	}

	passwordNew, err := service.AesEncrypt(password)
	if err != nil {
		a.Data["json"] = &RespError{
			Code: encryptError,
			Msg:  err.Error(),
		}
		a.ServeJSON()
	}

	user.Password = service.Md5(passwordNew)
	res, err := models.Register(user)
	if err != nil {
		a.Data["json"] = &RespError{
			Code: registerError,
			Msg:  err.Error(),
		}
		a.ServeJSON()
	}

	if res != nil {
		userInfo := models.UserInfo{
			Id:       res.Id,
			Username: res.Username,
		}
		token, err := service.GenToken(userInfo)
		if err != nil {
			a.Data["json"] = &RespError{
				Code: tokenError,
				Msg:  err.Error(),
			}
			a.ServeJSON()
		}

		a.Data["json"] = models.TokenInfo{
			Id:       res.Id,
			Username: res.Username,
			Token:    token,
		}
		a.ServeJSON()
	}
}

// @Title 登录
// @Description 用户登录
// @Param  username  body  string  true  "手机号"
// @Param  password  body  string  true  "密码"
// @router /login [post]
func (a *AuthController) Login() {
	username := a.GetString("username")
	password := a.GetString("password")

	user := new(models.User)
	user.Username = username
	user.Password = password

	valid := validation.Validation{}
	valid.Required(user.Username, "username")
	valid.Required(user.Password, "password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			a.Data["json"] = &RespError{
				Code: paramError,
				Msg:  err.Message,
			}
			a.ServeJSON()
		}
	}

	passwordNew, err := service.AesEncrypt(password)
	if err != nil {
		a.Data["json"] = &RespError{
			Code: encryptError,
			Msg:  err.Error(),
		}
		a.ServeJSON()
	}

	user.Password = service.Md5(passwordNew)
	res, err := models.Login(user)
	if err != nil {
		a.Data["json"] = &RespError{
			Code: loginError,
			Msg:  err.Error(),
		}
		a.ServeJSON()
	}

	if res != nil {
		userInfo := models.UserInfo{
			Id:       res.Id,
			Username: res.Username,
		}

		token, err := service.GenToken(userInfo)
		if err != nil {
			a.Data["json"] = &RespError{
				Code: tokenError,
				Msg:  err.Error(),
			}
			a.ServeJSON()
		}

		a.Data["json"] = models.TokenInfo{
			Id:       res.Id,
			Username: res.Username,
			Token:    token,
		}
		a.ServeJSON()
	}

}
