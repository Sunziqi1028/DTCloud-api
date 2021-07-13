/**
* @Author: lik
* @Date: 2021/3/5 19:08
* @Version 1.0
 */
package controller

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/http/validator/core/factory"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) HandlerRouter(vApi *gin.RouterGroup) {

	// 刷新token，当token过期，用旧token换取新token
	vApi.POST("access_token", factory.Create(constant.ValidatorPrefix+"RefreshToken"))
	vApi.POST("update_password", factory.Create(constant.ValidatorPrefix+"RefreshPassword"))
	vApi.POST("create/:id", factory.Create(constant.ValidatorPrefix+"PublicCreate"))
	vApi.POST("read/:id", factory.Create(constant.ValidatorPrefix+"PublicRead"))
	vApi.POST("page/:id", factory.Create(constant.ValidatorPrefix+"PublicPage"))
	//vApi.POST("write", factory.Create(constant.ValidatorPrefix+"PublicWrite"))
	//vApi.POST("unlink", factory.Create(constant.ValidatorPrefix+"PublicUnlink"))
	//vApi.POST("getattr", factory.Create(constant.ValidatorPrefix+"Getattr"))
	//vApi.POST("public_getattr", factory.Create(constant.ValidatorPrefix+"PublicGetattr"))
	//vApi.POST("upload_img", factory.Create(constant.ValidatorPrefix+"UploadFiles"))

}
