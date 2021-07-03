/**
* @Author: lik
* @Date: 2021/3/7 16:43
* @Version 1.0
 */
package upload_file

import (
	//"gitee.com/open-product/dtcloud-api/global/variable"
	//"gitee.com/open-product/dtcloud-api/util/response"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"path"
	//"strings"
)

func Upload(context *gin.Context, savePath string) bool {

	////  1.获取上传的文件名(参数验证器已经验证完成了第一步错误，这里简化)
	//file, _ := context.FormFile(variable.ConfigYml.GetString("FileUploadSetting.UploadFileField")) //  file 是一个文件结构体（文件对象）
	//
	////  保存文件，原始文件名进行全局唯一编码加密、md5 加密，保证在后台存储不重复
	//var saveErr error
	//if sequence := variable.SnowFlake.GetId(); sequence > 0 {
	//	saveFileName := fmt.Sprintf("%d%s", sequence, file.Filename)
	//	saveFileName = md5_encrypt.MD5(saveFileName) + path.Ext(saveFileName)
	//
	//	if saveErr = context.SaveUploadedFile(file, savePath+saveFileName); saveErr == nil {
	//		//  上传成功,返回资源的相对路径，这里请根据实际返回绝对路径或者相对路径
	//		success := gin.H{
	//			"path": strings.ReplaceAll(savePath+saveFileName, variable.BasePath, ""),
	//		}
	//		response.Success(context, consts.CurdStatusOkMsg, success)
	//		return true
	//	}
	//} else {
	//	saveErr = errors.New(my_errors.ErrorsSnowflakeGetIdFail)
	//}
	//response.Fail(context, consts.FilesUploadFailCode, consts.FilesUploadFailMsg+", 文件保存失败!", saveErr.Error())
	return false

}
