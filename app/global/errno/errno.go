/**
* @Author: lik
* @Date: 2021/3/5 17:45
* @Version 1.0
 */
package errno

const (
	//系统部分
	ErrorsContainerKeyAlreadyExists string = "该键已经注册在容器中了"
	ErrorsPublicNotExists           string = "public 目录不存在"
	ErrorsConfigYamlNotExists       string = "config.yaml 配置文件不存在"
	ErrorsConfigGormNotExists1      string = "gorm_v2.yaml 配置文件不存在"
	ErrorsConfigGormNotExists2      string = "gorm_v2.yaml 配置文件不存在"
	ErrorsStorageLogsNotExists      string = "storage/logs 目录不存在"
	ErrorsConfigInitFail            string = "初始化配置文件发生错误"
	ErrorsFuncEventAlreadyExists    string = "注册函数类事件失败，键名已经被注册"
	ErrorsFuncEventNotRegister      string = "没有找到键名对应的函数"
	ErrorsFuncEventNotCall          string = "注册的函数无法正确执行"
	ErrorsBasePath                  string = "初始化项目根目录失败"
	ErrorsNoAuthorization           string = "token鉴权未通过，请通过token授权接口重新获取token,"
	ErrorsGormInitFail              string = "Gorm 数据库驱动、连接初始化失败"
	ErrorsGormNotInitGlobalPointer  string = "%s 数据库全局变量指针没有初始化，请在配置文件 Gormv2.yml 设置 Gormv2.%s.IsInitGolobalGormMysql = 1, 并且保证数据库配置正确 \n"
	// 数据库部分
	ErrorsDbDriverNotExists string = "数据库驱动类型不存在,目前支持的数据库类型：mysql、sqlserver、postgresql，您提交数据库类型："
	ErrorsDialectDbInitFail string = "gorm 初始化失败:"

	// token部分
	JwtTokenFormatErrMsg string = "提交的 token 格式错误" //提交的 token 格式错误
	ErrorsTokenInvalid   string = "无效的token"

	// login
	LoginPasswordErrMsg    string = "账号密码错误" //未知的账号
	VerificationCodeErrMsg string = "验证码错误"  //验证码错误

	// 验证器错误
	ErrorsValidatorNotExists      string = "不存在的验证器"
	ErrorsValidatorBindParamsFail string = "验证器绑定参数失败"

	//文件上传
	ErrorsFilesUploadOpenFail string = "打开文件失败，详情："
	ErrorsFilesUploadReadFail string = "读取文件32字节失败，详情："
)
