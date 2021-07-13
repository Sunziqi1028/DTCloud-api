/**
* @Author: lik
* @Date: 2021/3/10 9:30
* @Version 1.0
 */
package curd

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/service/curd"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type ReadData struct {
}

func (c *ReadData) PublicRead(context *gin.Context) {
	fields := context.GetString(constant.ValidatorPrefix + "fields")
	ids := context.GetString(constant.ValidatorPrefix + "ids")
	page := context.GetInt64(constant.ValidatorPrefix + "page")
	limit := context.GetInt64(constant.ValidatorPrefix + "limit")
	order := context.GetString(constant.ValidatorPrefix + "order")
	table := strings.Replace(context.GetString(constant.ValidatorPrefix+"model"), ".", "_", -1)

	result := curd.CreateUserCurdFactory().PublicRead(fields, ids, page, limit, order, table)
	if result != nil {
		response.Success(context, constant.CurdStatusOkMsg, result)
		return
	}
	response.Fail(context, constant.CurdSelectFailCode, constant.CurdSelectFailMsg, "")
	return

}
