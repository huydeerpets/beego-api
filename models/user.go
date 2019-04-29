package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var db orm.Ormer
var sqlConn = beego.AppConfig.String("sqlconn")

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sqlConn)

	orm.RegisterModel(new(User))

	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	db = orm.NewOrm()
	db.Using("default")
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenInfo struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserInfo struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func Register(User *User) (*User, error) {
	_, err := db.Insert(User)
	if err != nil {
		return nil, err
	}

	return User, nil
}

func Login(user *User) (*TokenInfo, error) {
	var TokenInfo *TokenInfo
	err := db.Raw("SELECT id,username FROM user WHERE username = ? AND password = ?", user.Username, user.Password).QueryRow(&TokenInfo)
	if err != nil {
		return nil, err
	}

	return TokenInfo, nil
}

func GetUserInfo(user *User) (*UserInfo, error) {
	var userInfo *UserInfo
	err := db.Raw("SELECT id,username FROM user WHERE id = ? AND username = ?", user.Id, user.Username).QueryRow(&userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
