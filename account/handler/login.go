package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"github/XiaoXinTao/goproject/account/model"
)

func Login(ctx *gin.Context) {
	var LoginReq model.LoginReq
	if err := ctx.Bind(LoginReq); err != nil {
		log.Fatalf("登录请求, 绑定参数错误, err = %+v", err)

	}
}