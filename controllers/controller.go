package controllers

const paramError  = 10000
const registerError  = 10001
const loginError  = 10002
const tokenError  = 10003
const encryptError = 10004
const decryptError = 10005
const jsonError = 10006
const getUserInfoError = 10007

type RespError struct {
	Code int 			`json:"code"`
	Msg  string 		`json:"msg"`
}
