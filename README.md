# beego-api
beego api框架实践

## 第一次运行
app.conf修改端口，bee run，访问http://192.168.20.188:8000/v1/user

## 加密方案
密码认证：先AES加再Md5

签名认证：用户信息AES加密后存入Token
