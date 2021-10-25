package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义程序中的响应
type ResponseDate struct {
	Code ResCode     `json:"code"`           //程序中的错误码
	Msg  interface{} `json:"msg"`            //错误的提示信息
	Data interface{} `json:"data,omitempty"` //要存放的数据
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseDate{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	//自定义错误
	c.JSON(http.StatusOK, &ResponseDate{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	//响应成功时可以传递一些数据
	c.JSON(http.StatusOK, &ResponseDate{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
