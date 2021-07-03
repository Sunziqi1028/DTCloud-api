/**
* @Author: lik
* @Date: 2021/3/8 14:26
* @Version 1.0
 */
package user

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"gitee.com/open-product/dtcloud-api/handler/user"
	"gitee.com/open-product/dtcloud-api/validator/core/data_transfer"
	"github.com/gin-gonic/gin"
)

type Signup struct {
	Login     string `form:"login" json:"login"  binding:"required,min=1"`                   // 必填、对于文本,表示它的长度>=1
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`       //  密码为 必填，长度>=6
	Passwords string `form:"passwords" json:"passwords" binding:"required,eqfield=Password"` //  密码为 必填，长度>=6
	SqlType   string `form:"sqlType" json:"sqlType"  binding:"required,min=1"`
}

func (l Signup) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&l); err != nil {
		errs := gin.H{
			"tips": "UserRegister参数校验失败，参数不符合规定，Login 长度(>=1)、pass长度[6,20]、二次密码一致，不允许注册",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(l, constant.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "userLogin表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&user.Users{}).Signup(extraAddBindDataContext)
	}

}
