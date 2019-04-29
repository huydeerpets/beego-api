package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
)

var cryptKey = []byte(beego.AppConfig.String("cryptkey"))

func AesEncrypt(text string) (string, error) {
	var iv = cryptKey[:aes.BlockSize]
	encrypted := make([]byte, len(text))

	block, err := aes.NewCipher(cryptKey)
	if err != nil {
		return "", err
	}

	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))

	return hex.EncodeToString(encrypted), nil
}

func AesDecrypt(encrypted string) (string, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	var iv = cryptKey[:aes.BlockSize]
	decrypted := make([]byte, len(src))

	var block cipher.Block
	block, err = aes.NewCipher([]byte(cryptKey))
	if err != nil {
		return "", err
	}

	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)

	return string(decrypted), nil
}

func Md5(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
