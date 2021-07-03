/**
* @Author: lik
* @Date: 2021/3/8 9:34
* @Version 1.0
 */
package authorization

import (
	"gitee.com/open-product/dtcloud-api/app/global/errno"
	"gitee.com/open-product/dtcloud-api/app/global/token"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

func CheckAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//  模拟验证token
		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			variable.ZapLog.Error(errno.ErrorsValidatorBindParamsFail, zap.Error(err))
			context.Abort()
		}

		if len(headerParams.Authorization) >= 20 {
			accessToken := strings.Split(headerParams.Authorization, " ")
			if len(accessToken) == 2 && len(accessToken[1]) >= 20 {
				tokenIsEffective := token.CreateUserFactory().IsEffective(accessToken[1])
				if tokenIsEffective {
					context.Next()
				} else {
					response.ErrorAuthFail(context)
				}
			}
		} else {
			response.ErrorAuthFail(context)
		}

	}
}
