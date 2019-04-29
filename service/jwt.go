package service

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var jwtKey = beego.AppConfig.String("jwtkey")

func GenToken(encryptData interface{}) (string, error) {
	encryptJson, err := json.Marshal(encryptData)
	if err != nil {
		return "", err
	}

	sign, err := AesEncrypt(string(encryptJson))
	if err != nil {
		return "", err
	}

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sign"] = sign

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckToken(token string) (b bool, t *jwt.Token) {
	kv := strings.Split(token, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		return false, nil
	}

	t, err := jwt.Parse(kv[1], func(*jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return false, nil
	}

	return true, t
}
