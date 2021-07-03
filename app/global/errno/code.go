/**
* @Author: lik
* @Date: 2021/3/5 17:48
* @Version 1.0
 */
package errno

const (
	SUCCESS int = 0 //操作成功
	FAILED  int = 1 //操作失败

	JwtTokenOK            int = 2001 // token有效
	JwtTokenInvalid       int = 2002 // 无效的token
	JwtTokenExpired       int = 2003 // 过期的token
	JwtTokenFormatErrCode int = 2004 // 提交的 token 格式错误

	DeleteFail int = 1001 //删除失败
	UpdateFail int = 1002 //修改失败
	SelectFail int = 1003 //查询失败
	CreateFail int = 1004 //创建失败

)
