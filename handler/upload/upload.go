/**
* @Author: lik
* @Date: 2021/3/7 16:42
* @Version 1.0
 */
package upload

import (
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/app/service/upload_file"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

//  文件上传是一个独立模块，给任何业务返回文件上传后的存储路径即可。
// 开始上传
func (u *Upload) StartUpload(context *gin.Context) bool {
	savePath := variable.BasePath + variable.ConfigYml.GetString("FileUploadSetting.UploadFileSavePath")
	return upload_file.Upload(context, savePath)
}
