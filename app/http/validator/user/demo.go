/**
* @Author: lik
* @Date: 2021/3/11 21:40
* @Version 1.0
 */
package user

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/http/controller/user"
	"gitee.com/open-product/dtcloud-api/app/http/validator/core/data_transfer"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"github.com/gin-gonic/gin"
)

type CeS struct {
}

func (c CeS) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&c); err != nil {
		errs := gin.H{
			"tips": "UserRegister参数校验失败，参数不符合规定，user_name、pass、 长度有问题，不允许登录",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	extraAddBindDataContext := data_transfer.DataAddContext(c, constant.ValidatorPrefix, context)

	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "userLogin表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&user.Users{}).CeS(extraAddBindDataContext)
	}
}
