package controllers

import (
	"beego-api/service"
	"encoding/json"
	"github.com/astaxie/beego"
	"net/url"
	"strings"
)

const appId = "dfvbdfbdfbdfbdfbdf"
const secret = "bfbdfbfdbfbfdbdfbdfbfbbfdbdfb"

const loginCertificateVerify = "https://developer.toutiao.com/api/apps/jscode2session"
const getAccessToken = "https://developer.toutiao.com/api/apps/token"
const contentSecurity = "https://developer.toutiao.com/api/v2/tags/text/antidirt"

// 今日头条
type ToutiaoController struct {
	beego.Controller
}

// @Title 登录凭证校验
// @Description https://developer.toutiao.com/docs/open/jscode2session.html
// @Param  code           body  string  true  "login接口返回的登录凭证"
// @Param  anonymousCode  body  string  true  "login接口返回的匿名登录凭证"
// @router /loginCertificateVerify [post]
func (t *ToutiaoController) LoginCertificateVerify() {
	code := t.GetString("code")
	anonymousCode := t.GetString("anonymousCode")

	query := "?appid" + appId + "&secret" + secret + "&code" + code + "&anonymousCode" + anonymousCode
	resp, err := service.HttpGet(loginCertificateVerify + query)
	if err != nil {
		t.Data["json"] = &RespError{
			Code: httpError,
			Msg:  err.Error(),
		}
		t.ServeJSON()
	}

	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(resp), &mapResult)
	if err != nil {
		t.Data["json"] = &RespError{
			Code: jsonError,
			Msg:  err.Error(),
		}
		t.ServeJSON()
	}
	t.Data["json"] = mapResult
	t.ServeJSON()
}

// @Title 获取access_token
// @Description https://developer.toutiao.com/docs/open/accessToken.html
// @router /getAccessToken [post]
func (t *ToutiaoController) GetAccessToken() {
	query := "?appid" + appId + "&secret" + secret + "&grant_type=client_credential"
	resp, err := service.HttpGet(getAccessToken + query)
	if err != nil {
		t.Data["json"] = &RespError{
			Code: httpError,
			Msg:  err.Error(),
		}
		t.ServeJSON()
	}

	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(resp), &mapResult)
	if err != nil {
		t.Data["json"] = &RespError{
			Code: jsonError,
			Msg:  err.Error(),
		}
		t.ServeJSON()
	}
	t.Data["json"] = mapResult
	t.ServeJSON()
}

// @Title 内容安全检测
//
//  postValue := url.Values{
//		"accessToken": {accessToken},
//		"content": {content},
//	}
//	postString := postValue.Encode()
//	postReader := strings.NewReader(postString)
//
// @Description https://developer.toutiao.com/docs/open/textCheck.html
// @Param  accessToken  body  string  true  "小程序access_token"
// @Param  content      body  string  true  "要检测的文本"
// @router /contentSecurity [post]
func (t *ToutiaoController) ContentSecurity() {
	accessToken := t.GetString("accessToken")
	content := t.GetString("content")

	postValue := url.Values{
		"tasks": {content},
	}
	postString := postValue.Encode()
	postReader := strings.NewReader(postString)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Token": accessToken,
	}

	resp, err := service.HttpDo("POST", contentSecurity, postReader, header)
	if err != nil {
		t.Data["json"] = &RespError{
			Code: httpError,
			Msg:  err.Error(),
		}
		t.ServeJSON()
	}

	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(resp), &mapResult)
	if err != nil {
		t.Data["json"] = &RespError{
			Code: jsonError,
			Msg:  err.Error(),
		}
		t.ServeJSON()
	}
	t.Data["json"] = mapResult
	t.ServeJSON()
}
