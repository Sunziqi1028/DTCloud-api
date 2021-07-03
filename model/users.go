/**
* @Author: lik
* @Date: 2021/3/7 17:43
* @Version 1.0
 */
package model

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/app/service/user/password"
	"go.uber.org/zap"
)

func CreateUserFactory(sqlType string) *MemberUsers {
	return &MemberUsers{BaseModel: BaseModel{DB: useDbConn(sqlType)}}
}

type MemberUsers struct {
	BaseModel `json:"-"`
}

// 测试
func (ml *MemberUsers) Demo() *ResUsers {
	var u ResUsers
	sql := "select * from res_users"
	result := ml.Raw(sql).First(&u)
	if result.Error == nil {
		return &u

	} else {
		variable.ZapLog.Error("根据账号查询单条记录出错:", zap.Error(result.Error))
	}
	return nil
}

// 查询用户
func (ml *MemberUsers) QueryUserLogin(login string) *ResUsers {
	var u ResUsers
	sql := "select id from res_users where login=?  limit 1"
	result := ml.Raw(sql, login).First(&u)
	if result.Error == nil {
		return &u

	} else {
		variable.ZapLog.Error("根据账号查询单条记录出错:", zap.Error(result.Error))
	}
	return nil
}

// 创建 ResPartner 返回 Partner ID
func (ml *MemberUsers) InsertMemberPartner(p ResPartner) *ResPartner {
	ml.Table("res_partner").Create(&p).Raw("select LAST_INSERT_ID()")
	return &p

}

// 创建  ResUsers 返回 Users ID
func (ml *MemberUsers) InsertMemberUsers(u ResUsers) *ResUsers {
	ml.Table("res_users").Create(&u).Raw("select LAST_INSERT_ID()")
	return &u

}

// 创建 ResCompanyUsersRel
func (ml *MemberUsers) InsertResCompanyUsersRel(cu ResCompanyUsersRel) *ResCompanyUsersRel {
	ml.Table("res_company_users_rel").Create(&cu).Raw("select LAST_INSERT_ID()")
	return &cu

}

// 创建 ResGroupsUsersRel
func (ml *MemberUsers) InsertResGroupsUsersRel(gu ResGroupsUsersRel) *ResGroupsUsersRel {
	ml.Table("res_groups_users_rel").Create(&gu).Raw("select LAST_INSERT_ID()")
	return &gu

}

// 更新 Partner UserId
func (ml *MemberUsers) UpdateMemberPartnerUserId(partnerId int64, userId int64) bool {
	sql := "UPDATE res_partner SET user_id=? WHERE id=?"
	if ml.Exec(sql, userId, partnerId).Error == nil {
		return true
	}
	return false
}

// 删除 Partner User
func (ml *MemberUsers) DeleteMemberPartnerUserId(partnerId int64, userId int64) {
	sql := "delete from res_users WHERE id=?"
	if ml.Exec(sql, userId).Error == nil {
		sql1 := "delete from res_partner WHERE id=?"
		ml.Exec(sql1, partnerId)
	}

}

// 用户登录,
func (ml *MemberUsers) Login(userLogin string, pass string) *ResUsers {
	var u ResUsers
	sql := "select id,login,partner_id,company_id,password,mobile from res_users where login=?  limit 1"
	result := ml.Raw(sql, userLogin).First(&u)
	if result.Error == nil {
		// 账号密码验证成功
		if len(u.Password) > 0 {
			err := password.SHA512Encrypt.Verify(pass, u.Password)
			if err != nil {
				return nil
			}
			return &u
		}
	} else {
		variable.ZapLog.Error("根据账号查询单条记录出错:", zap.Error(result.Error))
	}
	return nil
}

// 判断用户token是否在数据库存在+状态OK
func (ml *MemberUsers) OauthCheckTokenIsOk(userId int64, token string) bool {
	sql := "SELECT access_token  FROM  dtcloud_token  WHERE   user_id=?  AND  token_date>NOW() ORDER  BY  token_date  DESC  LIMIT ?"
	rows, err := ml.Raw(sql, userId, constant.JwtTokenOnlineUsers).Rows()
	if err == nil && rows != nil {
		for rows.Next() {
			var tempToken string
			err := rows.Scan(&tempToken)
			if err == nil {
				if tempToken == token {
					_ = rows.Close()
					return true
				}
			}
		}
		//  凡是查询类记得释放记录集
		_ = rows.Close()
	}
	return false
}

//记录用户登陆（login）生成的token，每次登陆记录一次token
func (ml *MemberUsers) OauthLoginToken(userId int64, token string, expiresAt int64, clientIp string) bool {

	sql := "INSERT INTO dtcloud_token(appid,secret,name,company_id,user_id,access_token,token_date) " +
		"SELECT 1,1,(select login from res_users where id = ? ),1,?,?,TO_TIMESTAMP(?) " +
		"FROM (select 1) DUAL WHERE NOT EXISTS (Select 1 FROM dtcloud_token " +
		"where user_id = ? and access_token=? )"
	//注意：token的精确度为秒，如果在一秒之内，一个账号多次调用接口生成的token其实是相同的，这样写入数据库，第二次的影响行数为0，知己实际上操作仍然是有效的。
	//所以这里只判断无错误即可，判断影响行数的话，>=0 都是ok的
	if ml.Exec(sql, userId, userId, token, expiresAt, userId, token).Error == nil {
		return true
	}
	return false
}

//TODO Mysql
func (ml *MemberUsers) OauthLoginTokens(userId int64, token string, expiresAt int64, clientIp string) bool {
	sql := "INSERT INTO dtcloud_token(appid,secret,name,company_id,user_id,access_token,token_date) " +
		"SELECT 1,1,(select login from res_users where id = ? ),1,?,?,FROM_UNIXTIME(?) " +
		"FROM (select 1) DUAL WHERE NOT EXISTS (Select 1 FROM dtcloud_token " +
		"where user_id = ? and access_token=?)"
	//注意：token的精确度为秒，如果在一秒之内，一个账号多次调用接口生成的token其实是相同的，这样写入数据库，第二次的影响行数为0，知己实际上操作仍然是有效的。
	//所以这里只判断无错误即可，判断影响行数的话，>=0 都是ok的
	if ml.Exec(sql, userId, userId, token, expiresAt, userId, token).Error == nil {
		return true
	}
	return false
}

//用户刷新 token
func (ml *MemberUsers) OauthRefreshToken(userId, expiresAt int64, oldToken, newToken, clientIp string) bool {
	sql := "UPDATE   dtcloud_token   SET  access_token=? ,token_date=TO_TIMESTAMP(?)  WHERE user_id=? and access_token=? "
	if ml.Exec(sql, newToken, expiresAt, userId, oldToken).Error == nil {
		return true
	}
	return false
}

//当用户更改密码后，所有的token都失效，必须重新登录
func (ml *MemberUsers) OauthResetToken(userId int64, newPass string) bool {
	//如果用户新旧密码一致，直接返回true，不需要处理
	//userItem, err := ml.ShowOneItem(userId)
	//if userItem != nil && err == nil && userItem.Password == newPass {
	//if userItem != nil && err == nil  {
	//	return true
	//} else if userItem != nil {
	sql := "UPDATE  dtcloud_token  SET  token_date=NOW()  WHERE  user_id=?"
	if ml.Exec(sql, userId).Error == nil {
		return true
	}
	//}
	return false
}

//根据用户ID查询一条信息
//func (ml *MemberUsers) ShowOneItem(userId int64) (*ResUsers, error) {
//	var u ResUsers
//	sql := "SELECT id,login,password FROM res_users WHERE  id=? LIMIT 1"
//	result := ml.Raw(sql, userId).First(u)
//	if result.Error == nil {
//		return &u, nil
//	} else {
//		return nil, result.Error
//	}
//}

// 修改密码
func (ml *MemberUsers) RefreshPassword(pass string, uid int64) bool {

	sql := "update res_users set password=? WHERE id=?"
	if ml.Exec(sql, pass, uid).RowsAffected >= 0 {
		if ml.OauthResetToken(uid, pass) {
			return true
		}
	}
	return false
}
