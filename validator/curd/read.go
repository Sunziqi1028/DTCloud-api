/**
* @Author: lik
* @Date: 2021/3/10 9:29
* @Version 1.0
 */
package curd

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"gitee.com/open-product/dtcloud-api/handler/curd"
	"gitee.com/open-product/dtcloud-api/validator/core/data_transfer"
	"github.com/gin-gonic/gin"
)

type PublicRead struct {
	AccessToken  string `form:"access_token" json:"access_token" binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	Model        string `form:"model" json:"model" binding:"required,min=1"`
	Uid          int16  `form:"uid" json:"uid" binding:"required,min=1"`
	PartnerId    int16  `form:"partner_id" json:"partner_id" binding:"required,min=1"`
	Lang         string `form:"lang" json:"lang"`
	Cache        string `form:"cache" json:"cache"`
	DataType     string `form:"data_type" json:"data_type"`
	Ids          string `form:"ids" json:"ids"`
	Fields       string `form:"fields" json:"fields"`
	FunctionName string `form:"function_name" json:"function_name"`
	Page         int16  `form:"page" json:"page"`
	Limit        int16  `form:"limit" json:"limit"`
	Order        string `form:"order" json:"order"`
}

func (c PublicRead) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&c); err != nil {
		errs := gin.H{
			"tips": "参数校验失败，参数不符合规定",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(c, constant.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&curd.ReadData{}).PublicRead(extraAddBindDataContext)
	}
}
