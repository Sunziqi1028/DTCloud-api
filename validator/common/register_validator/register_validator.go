/**
* @Author: lik
* @Date: 2021/3/7 15:28
* @Version 1.0
 */
package register_validator

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/validator/common/upload_files"
	"gitee.com/open-product/dtcloud-api/validator/core/container"
	"gitee.com/open-product/dtcloud-api/validator/curd"
	"gitee.com/open-product/dtcloud-api/validator/user"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func RegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string

	// 文件上传
	key = constant.ValidatorPrefix + "UploadFiles"
	containers.Set(key, upload_files.UpFiles{})

	// Users 模块表单验证器按照 key => value 形式注册在容器，方便路由模块中调用

	key = constant.ValidatorPrefix + "UsersLogin"
	containers.Set(key, user.Login{})
	key = constant.ValidatorPrefix + "UsersSignup"
	containers.Set(key, user.Signup{})
	key = constant.ValidatorPrefix + "RefreshToken"
	containers.Set(key, user.RefreshToken{})
	key = constant.ValidatorPrefix + "RefreshPassword"
	containers.Set(key, user.RefreshPassword{})
	key = constant.ValidatorPrefix + "PublicCreate"
	containers.Set(key, curd.PublicCreate{})
	key = constant.ValidatorPrefix + "PublicRead"
	containers.Set(key, curd.PublicRead{})

	key = constant.ValidatorPrefix + "CeS"
	containers.Set(key, user.CeS{})

}
