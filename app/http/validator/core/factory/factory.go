/**
* @Author: lik
* @Date: 2021/3/7 16:19
* @Version 1.0
 */
package factory

import (
	"gitee.com/open-product/dtcloud-api/app/global/errno"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/app/http/validator/core/container"
	"gitee.com/open-product/dtcloud-api/app/http/validator/core/interf"
	"github.com/gin-gonic/gin"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}
	variable.ZapLog.Error(errno.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
