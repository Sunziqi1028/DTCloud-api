/**
* @Author: lik
* @Date: 2021/3/7 16:31
* @Version 1.0
 */
package upload_files

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/app/util/files"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"gitee.com/open-product/dtcloud-api/handler/upload"
	"github.com/gin-gonic/gin"
	"strings"
)

type UpFiles struct {
}

// 文件上传公共模块表单参数验证器
func (u UpFiles) CheckParams(context *gin.Context) {
	tmpFile, err := context.FormFile(variable.ConfigYml.GetString("FileUploadSetting.UploadFileField")) //  file 是一个文件结构体（文件对象）
	var isPass bool
	//获取文件发生错误，可能上传了空文件等
	if err != nil {
		response.Fail(context, constant.FilesUploadFailCode, constant.FilesUploadFailMsg, err.Error())
		return
	}
	//超过系统设定的最大值：32M，tmpFile.Size 的单位是 bytes 和我们定义的文件单位M 比较，就需要将我们的单位*1024*1024(即2的20次方)，一步到位就是 << 20
	if tmpFile.Size > variable.ConfigYml.GetInt64("FileUploadSetting.Size")<<20 {
		response.Fail(context, constant.FilesUploadMoreThanMaxSizeCode, constant.FilesUploadMoreThanMaxSizeMsg+variable.ConfigYml.GetString("FileUploadSetting.Size"), "")
		return
	}
	//不允许的文件mime类型
	if fp, err := tmpFile.Open(); err == nil {
		mimeType := files.GetFilesMimeByFp(fp)

		for _, value := range variable.ConfigYml.GetStringSlice("FileUploadSetting.AllowMimeType") {
			if strings.ReplaceAll(value, " ", "") == strings.ReplaceAll(mimeType, " ", "") {
				isPass = true
				break
			}
		}
		_ = fp.Close()
	} else {
		response.ErrorSystem(context, constant.ServerOccurredErrorMsg, "")
		return
	}
	//凡是存在相等的类型，通过验证，调用控制器
	if !isPass {
		response.Fail(context, constant.FilesUploadMimeTypeFailCode, constant.FilesUploadMimeTypeFailMsg, "")
	} else {
		(&upload.Upload{}).StartUpload(context)
	}
}
