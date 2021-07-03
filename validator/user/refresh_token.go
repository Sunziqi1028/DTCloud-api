/**
* @Author: lik
* @Date: 2021/3/8 9:42
* @Version 1.0
 */
package user

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"gitee.com/open-product/dtcloud-api/handler/user"
	"github.com/gin-gonic/gin"
	"strings"
)

type RefreshToken struct {
	Authorization string `json:"access_token" header:"Authorization" binding:"required,min=20"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (r RefreshToken) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBindHeader(&r); err != nil {
		errs := gin.H{
			"tips": "Token参数校验失败，参数不符合规定，token 长度>=20",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}
	token := strings.Split(r.Authorization, " ")
	if len(token) == 2 {
		context.Set(constant.ValidatorPrefix+"token", token[1])
		(&user.Users{}).RefreshToken(context)
	} else {
		errs := gin.H{
			"tips": "Token不合法，token请放置在header头部分，按照按=>键提交，例如：Authorization：Bearer 你的实际token....",
		}
		response.Fail(context, constant.JwtTokenFormatErrCode, constant.JwtTokenFormatErrMsg, errs)
	}

}
