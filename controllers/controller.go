package controllers

const paramError = 10000
const registerError = 10001
const loginError = 10002
const CheckError = 10003
const tokenError = 10004
const encryptError = 10005
const decryptError = 10006
const jsonError = 10007
const getUserInfoError = 10008
const httpError = 10009

type RespError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
