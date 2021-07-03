/**
* @Author: lik
* @Date: 2021/3/7 16:46
* @Version 1.0
 */
package user

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/global/token"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/app/service/user"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"gitee.com/open-product/dtcloud-api/model"
	"github.com/gin-gonic/gin"
	"time"
)

type Users struct {
}

func (u *Users) CeS(context *gin.Context) {
	member := model.CreateUserFactory("").Demo()
	response.Success(context, constant.CurdStatusOkMsg, member)
	return
}

// 1.用户注册
func (u *Users) Signup(context *gin.Context) {
	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、GetInt64()、GetFloat64（）等快捷获取需要的数据类型，注意：相关键名规则：  前缀+验证器结构体中的 json 标签
	// 当然也可以通过gin框架的上下文原始方法获取，例如： context.PostForm("user_name") 获取，这样获取的数据格式为文本，需要自己继续转换
	userName := context.GetString(constant.ValidatorPrefix + "login")
	pass := context.GetString(constant.ValidatorPrefix + "password")
	sqlType := context.GetString(constant.ValidatorPrefix + "sqlType")
	//userIp := context.ClientIP()
	user.CreateUserCurdFactory().Register(userName, pass, sqlType, context)

}

//  2.用户登录
func (u *Users) Login(context *gin.Context) {
	userLogin := context.GetString(constant.ValidatorPrefix + "login")
	userPassword := context.GetString(constant.ValidatorPrefix + "password")
	loginType := context.GetString(constant.ValidatorPrefix + "type")
	sqlType := context.GetString(constant.ValidatorPrefix + "sqlType")

	if sqlType == "" {
		sqlType = variable.ConfigGormYml.GetString("Gormv2.UseDbType")
	}

	var userModel *model.ResUsers

	if loginType == "0" {
		userModel = user.CreateUserCurdFactory().UserLogin(userLogin, userPassword, sqlType)
	} else if loginType == "1" {
		userModel = user.CreateUserCurdFactory().SmsLogin(userLogin, userPassword, sqlType, context)
	}

	if userModel != nil {
		userTokenFactory := token.CreateUserFactory()
		userToken, err := userTokenFactory.GenerateToken(userModel.Id, userModel.Login, sqlType, variable.ConfigYml.GetInt64("Token.JwtTokenCreatedExpireAt"))
		if err == nil {
			if userTokenFactory.RecordLoginToken(userToken, context.ClientIP()) {
				data := gin.H{
					"uid":          userModel.Id,
					"partner_id":   userModel.PartnerId,
					"company_id":   userModel.CompanyId,
					"mobile":       userModel.Mobile,
					"login":        userLogin,
					"access_token": userToken,
					"create_date":  time.Now().Format("2006-01-02 15:04:05"),
				}
				response.Success(context, constant.CurdStatusOkMsg, data)
				return
			}
		}
	}
	response.Fail(context, constant.CurdLoginFailCode, constant.CurdLoginFailMsg, "")
}

// 刷新用户token
func (u *Users) RefreshToken(context *gin.Context) {
	oldToken := context.GetString(constant.ValidatorPrefix + "token")
	if newToken, ok := token.CreateUserFactory().RefreshToken(oldToken, context.ClientIP()); ok {
		res := gin.H{
			"access_token": newToken,
		}
		response.Success(context, constant.CurdStatusOkMsg, res)
	} else {
		response.Fail(context, constant.CurdRefreshTokenFailCode, constant.CurdRefreshTokenFailMsg, "")
	}
}

// 修改密码
func (u *Users) RefreshPassword(context *gin.Context) {

	if user.CreateUserCurdFactory().RefreshPassword(context) {
		response.Success(context, constant.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, constant.CurdUpdateFailCode, constant.CurdUpdateFailMsg, "")
	}

}
