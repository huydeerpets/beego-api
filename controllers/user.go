package controllers

import (
	"beego-api/models"
	"beego-api/service"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"

	"github.com/astaxie/beego"
)

// 用户
type UserController struct {
	beego.Controller
}

// @Title 用户信息
// @Description 用户信息
// @router / [get]
func (u *UserController) Info() {
	token := u.Ctx.Input.Header("Authorization")

	_, t := service.CheckToken(token)
	claims, _ := t.Claims.(jwt.MapClaims)

	usersJson, err := service.AesDecrypt(claims["sign"].(string))
	if err != nil {
		u.Data["json"] = &RespError{
			Code: decryptError,
			Msg:  err.Error(),
		}
		u.ServeJSON()
	}

	var userInfo models.UserInfo
	err = json.Unmarshal([]byte(usersJson), &userInfo)
	if err != nil {
		u.Data["json"] = &RespError{
			Code: jsonError,
			Msg:  err.Error(),
		}
		u.ServeJSON()
	}

	user := new(models.User)
	user.Id = userInfo.Id
	user.Username = userInfo.Username

	res, err := models.GetUserInfo(user)
	if err != nil {
		u.Data["json"] = &RespError{
			Code: getUserInfoError,
			Msg:  err.Error(),
		}
		u.ServeJSON()
	}

	u.Data["json"] = res
	u.ServeJSON()
}
