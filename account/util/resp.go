package util

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github/XiaoXinTao/goproject/account/consts"
)

func RetNil(ctx *gin.Context) {
	re := gin.H{
		"code":    consts.ErrSuccess,
		"message": consts.ErrStrSuccess,
	}
	log.Printf("result=%v", re)
	ctx.JSON(http.StatusOK, re)
}

func RetJson(ctx *gin.Context, data interface{}) {
	re := gin.H{
		"code":    consts.ErrSuccess,
		"message": consts.ErrStrSuccess,
		"data":    data,
	}
	log.Printf( "result=%v", re)
	ctx.JSON(http.StatusOK, re)
}

func Ret(ctx *gin.Context, code int, msg string, data interface{}) {
	re := gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
	log.Printf( "result=%v", re)
	ctx.JSON(http.StatusOK, re)
}

func RetErrJson(ctx *gin.Context, code int, msg string, data interface{}) {
	re := gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
	log.Printf( "result=%v", re)
	ctx.JSON(http.StatusOK, re)
}

func RetErr(ctx *gin.Context, code int, msg string) {
	re := gin.H{
		"code":    code,
		"message": msg,
	}
	log.Printf( "result=%v", re)
	ctx.JSON(http.StatusOK, re)
}

func RetStr(ctx *gin.Context, msg string) {
	ctx.String(http.StatusOK, msg)
}
