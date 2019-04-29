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
func (u *AuthController) Register() {
	Username := u.GetString("username")
	password := u.GetString("password")

	user := new(models.User)

	user.Username = Username
	user.Password = password

	valid := validation.Validation{}
	valid.Required(user.Username, "username")
	valid.Required(user.Password, "password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			u.Data["json"] = &RespError{
				Code: paramError,
				Msg:  err.Message,
			}
			u.ServeJSON()
		}
	}

	passwordNew, err := service.AesEncrypt(password)
	if err != nil {
		u.Data["json"] = &RespError{
			Code: encryptError,
			Msg:  err.Error(),
		}
		u.ServeJSON()
	}

	user.Password = service.Md5(passwordNew)
	res, err := models.Register(user)
	if err != nil {
		u.Data["json"] = &RespError{
			Code: registerError,
			Msg:  err.Error(),
		}
		u.ServeJSON()
	}

	if res != nil {
		userInfo := models.UserInfo{
			Id:       res.Id,
			Username: res.Username,
		}
		token, err := service.GenToken(userInfo)
		if err != nil {
			u.Data["json"] = &RespError{
				Code: tokenError,
				Msg:  err.Error(),
			}
			u.ServeJSON()
		}

		u.Data["json"] = models.TokenInfo{
			Id:       res.Id,
			Username: res.Username,
			Token:    token,
		}
		u.ServeJSON()
	}
}

// @Title 注册
// @Description 用户注册
// @Param  username  body  string  true  "手机号"
// @Param  password  body  string  true  "密码"
// @router /login [post]
func (u *AuthController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")

	user := new(models.User)
	user.Username = username
	user.Password = password

	valid := validation.Validation{}
	valid.Required(user.Username, "username")
	valid.Required(user.Password, "password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			u.Data["json"] = &RespError{
				Code: paramError,
				Msg:  err.Message,
			}
			u.ServeJSON()
		}
	}

	passwordNew, err := service.AesEncrypt(password)
	if err != nil {
		u.Data["json"] = &RespError{
			Code: encryptError,
			Msg:  err.Error(),
		}
		u.ServeJSON()
	}

	user.Password = service.Md5(passwordNew)
	res, err := models.Login(user)
	if err != nil {
		u.Data["json"] = &RespError{
			Code: loginError,
			Msg:  err.Error(),
		}
		u.ServeJSON()
	}

	if res != nil {
		userInfo := models.UserInfo{
			Id:       res.Id,
			Username: res.Username,
		}

		token, err := service.GenToken(userInfo)
		if err != nil {
			u.Data["json"] = &RespError{
				Code: tokenError,
				Msg:  err.Error(),
			}
			u.ServeJSON()
		}

		u.Data["json"] = models.TokenInfo{
			Id:       res.Id,
			Username: res.Username,
			Token:    token,
		}
		u.ServeJSON()
	}

}

// @Title 退出
// @Description 用户退出
// @router /login [post]
func (u *AuthController) Logout() {

}
