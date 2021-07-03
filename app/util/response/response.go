/**
* @Author: lik
* @Date: 2021/3/7 16:35
* @Version 1.0
 */
package response

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/global/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnJson2(Context *gin.Context, httpCode int, errcode int, errmsg string, data interface{}, message string) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
		"data":    data,
		"message": message,
	})
}

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

// 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, constant.CurdStatusOkCode, msg, data)
}

// 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

//权限校验失败
func ErrorAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, errno.ErrorsNoAuthorization, "")
	//暂停执行
	c.Abort()
}

//参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, constant.ValidatorParamsCheckFailCode, constant.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, constant.ServerOccurredErrorCode, constant.ServerOccurredErrorMsg+msg, data)
	c.Abort()
}
