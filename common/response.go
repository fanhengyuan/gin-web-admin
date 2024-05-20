package common

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Success: true,
		Code:    errCode,
		Msg:     msg,
		Data:    data,
	})

	return
}
