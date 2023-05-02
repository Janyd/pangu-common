package response

import (
	common "gitee.com/ChuckChan/pangu-common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ReturnJson(c *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// ReturnJsonFromString 将json字符窜以标准json格式返回（例如，从redis读取json格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(c *gin.Context, httpCode int, jsonStr string) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(httpCode, jsonStr)
}

// 语法糖函数封装

// Success 直接返回成功
func Success(c *gin.Context, data interface{}) {
	ReturnJson(c, http.StatusOK, common.SuccessCode, common.SuccessMsg, data)
}

// Biz 业务错误
func Biz(c *gin.Context, code int, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, code, msg, data)
}

// Fail 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

// ErrorParam 参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, common.ParameterVerificationError, common.ParameterVerificationErrorMsg, wrongParam)
	c.Abort()
}

// Error 服务端错误
func Error(c *gin.Context, code int) {
	ReturnJson(c, http.StatusInternalServerError, code, common.ErrorCode[code], nil)
	c.Abort()
}

func Unauthorized(c *gin.Context, code int) {
	ReturnJson(c, http.StatusOK, code, common.ErrorCode[code], nil)
	c.Abort()
}

func ErrorErr(c *gin.Context, err error) {
	var msg string
	var code int
	var secondMsg string
	switch e := err.(type) {
	case *common.MessageError:
		msg = e.Error()
		code = e.Code()
		secondMsg = e.SecondMsg()
	default:
		logrus.WithError(err).Error("未知响应错误")
		code = common.SystemErrorCode
		msg = common.ErrMsg(code)
	}

	ReturnJson(c, http.StatusOK, code, msg+"; "+secondMsg, nil)
	c.Abort()
}
