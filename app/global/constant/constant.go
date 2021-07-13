/**
* @Author: lik
* @Date: 2021/3/5 17:49
* @Version 1.0
 */
package constant

const (

	// 表单验证器前缀
	ValidatorPrefix              string = "Form_Validator_"
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	// token部分
	JwtTokenSignKey       string = "DTCloudAPI"     // 设置默认Key
	JwtTokenOK            int    = 200100           //token有效
	JwtTokenInvalid       int    = -400100          //无效的token
	JwtTokenExpired       int    = -400101          //过期的token
	JwtTokenFormatErrCode int    = -400102          //提交的 token 格式错误
	JwtTokenFormatErrMsg  string = "提交的 token 格式错误" //提交的 token 格式错误
	JwtTokenOnlineUsers   int    = 10               // 设置一个账号最大允许几个用户同时在线，默认为10

	//文件上传
	FilesUploadFailCode            int    = -400250
	FilesUploadFailMsg             string = "文件上传失败, 获取上传文件发生错误!"
	FilesUploadMoreThanMaxSizeCode int    = -400251
	FilesUploadMoreThanMaxSizeMsg  string = "长传文件超过系统设定的最大值,系统允许的最大值（M）："
	FilesUploadMimeTypeFailCode    int    = -400252
	FilesUploadMimeTypeFailMsg     string = "文件mime类型不允许"

	//服务器代码发生错误
	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误, "

	// CURD 常用业务状态码
	CurdStatusOkCode         int    = 200
	CurdStatusOkMsg          string = "Success"
	CurdCreatFailCode        int    = -400200
	CurdCreatFailMsg         string = "新增失败"
	CurdUpdateFailCode       int    = -400201
	CurdUpdateFailMsg        string = "更新失败"
	CurdDeleteFailCode       int    = -400202
	CurdDeleteFailMsg        string = "删除失败"
	CurdSelectFailCode       int    = -400203
	CurdSelectFailMsg        string = "查询无数据"
	CurdRegisterFailCode     int    = -400204
	CurdRegisterFailMsg      string = "注册失败"
	CurdLoginFailCode        int    = -400205
	CurdLoginFailMsg         string = "登录失败"
	CurdRefreshTokenFailCode int    = -400206
	CurdRefreshTokenFailMsg  string = "刷新Token失败"

	//验证码
	CaptchaGetParamsInvalidMsg    string = "验证码长度错误"
	CaptchaGetParamsInvalidCode   int    = -400350
	CaptchaCheckParamsInvalidMsg  string = "参数错误"
	CaptchaCheckParamsInvalidCode int    = -400351
	CaptchaCheckOkMsg             string = "短信发送成功"
	CaptchaCheckOkCode            int    = 200
	CaptchaCheckFailCode          int    = -400355
	CaptchaCheckFailMsg           string = "短信发送失败"

	// 用户
	UsersLoginFailMsg  string = "用户已存在"
	UsersLoginFailCode int    = -400207

	UsersLoginPartnerFailMsg  string = "Partner创建失败"
	UsersLoginPartnerFailCode int    = -400207

	UsersLoginUsersFailMsg  string = "Users创建失败"
	UsersLoginUsersFailCode int    = -400207

	UsersLoginPWDErrcode int    = 1000
	UsersLoginPWDerrmsg  string = "no"
	UsersLoginPWDMessage string = "帐号密码错误! Access Denied"

	ParamsFailMsg      string = "参数解析失败"
	ErrorsTokenInvalid string = "无效的token"
)
