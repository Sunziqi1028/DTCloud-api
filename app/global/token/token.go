/**
* @Author: lik
* @Date: 2021/3/7 18:45
* @Version 1.0
 */
package token

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/global/errno"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/model"
	"gitee.com/open-product/dtcloud-api/routers/middleware/my_jwt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 创建 userToken 工厂

func CreateUserFactory() *UserToken {
	return &UserToken{
		UserJwt: my_jwt.CreateMyJWT(constant.JwtTokenSignKey),
	}
}

type UserToken struct {
	UserJwt *my_jwt.JwtSign
}

//生成token
func (u *UserToken) GenerateToken(userid int64, username string, sqltype string, expireAt int64) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customClaims := my_jwt.CustomClaims{
		UserId:  userid,
		Name:    username,
		SqlType: sqltype,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 my_jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireAt, // 失效截止时间
		},
	}
	return u.UserJwt.CreateToken(customClaims)
}

// 用户login成功，记录用户token
func (u *UserToken) RecordLoginToken(userToken, clientIp string) bool {
	if customClaims, err := u.UserJwt.ParseToken(userToken); err == nil {
		userId := customClaims.UserId
		expiresAt := customClaims.ExpiresAt
		sqlType := customClaims.SqlType

		if sqlType == "mysql" {
			return model.CreateUserFactory(sqlType).OauthLoginTokens(userId, userToken, expiresAt, clientIp)
		} else if sqlType == "postgresql" {
			return model.CreateUserFactory(sqlType).OauthLoginToken(userId, userToken, expiresAt, clientIp)
		} else {
			return false
		}

	} else {
		return false
	}
}

// 刷新token的有效期（默认+3600秒，参见常量配置项）
func (u *UserToken) RefreshToken(oldToken, clientIp string) (newToken string, res bool) {

	// 解析用户token的数据信息
	_, code := u.isNotExpired(oldToken)
	switch code {
	case constant.JwtTokenOK, constant.JwtTokenExpired:
		//如果token已经过期，那么执行更新
		if newToken, err := u.UserJwt.RefreshToken(oldToken, variable.ConfigYml.GetInt64("Token.JwtTokenRefreshExpireAt")); err == nil {
			if customClaims, err := u.UserJwt.ParseToken(newToken); err == nil {
				userId := customClaims.UserId
				expiresAt := customClaims.ExpiresAt
				sqlType := customClaims.SqlType
				if model.CreateUserFactory(sqlType).OauthRefreshToken(userId, expiresAt, oldToken, newToken, clientIp) {
					return newToken, true
				}
			}
		}
	case constant.JwtTokenInvalid:
		variable.ZapLog.Error(errno.ErrorsTokenInvalid)
	}

	return "", false
}

// 销毁token，基本用不到，因为一个网站的用户退出都是直接关闭浏览器窗口，极少有户会点击“注销、退出”等按钮，销毁token其实无多大意义
func (u *UserToken) DestroyToken() {

}

// 判断token是否未过期
func (u *UserToken) isNotExpired(token string) (*my_jwt.CustomClaims, int) {
	if customClaims, err := u.UserJwt.ParseToken(token); err == nil {

		if time.Now().Unix()-customClaims.ExpiresAt < 0 {
			// token有效
			return customClaims, constant.JwtTokenOK
		} else {
			// 过期的token
			return customClaims, constant.JwtTokenExpired
		}
	} else {
		// 无效的token
		return nil, constant.JwtTokenInvalid
	}
}

// 判断token是否有效（未过期+数据库用户信息正常）
func (u *UserToken) IsEffective(token string) bool {
	customClaims, code := u.isNotExpired(token)
	sqlType := customClaims.SqlType

	if constant.JwtTokenOK == code {
		//if user_item := Model.CreateUserFactory("").ShowOneItem(customClaims.UserId); user_item != nil {
		if model.CreateUserFactory(sqlType).OauthCheckTokenIsOk(customClaims.UserId, token) {
			return true
		}
	}
	return false
}
