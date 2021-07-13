/**
* @Author: lik
* @Date: 2021/3/7 17:41
* @Version 1.0
 */
package user

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/global/token"
	"gitee.com/open-product/dtcloud-api/app/model"
	"gitee.com/open-product/dtcloud-api/app/service/user/password"
	"gitee.com/open-product/dtcloud-api/app/util/cache"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateUserCurdFactory() *UsersCurd {

	return &UsersCurd{}
}

type UsersCurd struct {
}

func (u *UsersCurd) Register(userName, pass, sqlType string, context *gin.Context) {
	pwd, _ := password.SHA512Encrypt.Hash(pass) // 预先处理密码加密，然后存储在数据库

	member := model.CreateUserFactory(sqlType).QueryUserLogin(userName)

	if member.Id != 0 {
		response.Fail(context, constant.UsersLoginFailCode, constant.UsersLoginFailMsg, "")
		return
	}

	partner := model.ResPartner{}
	partner.Name = userName
	partner.CompanyId = 1
	partner.CreateDate = time.Now()
	partner.DisplayName = userName
	partner.Lang = "zh_CN"
	partner.Tz = "Asia/Shanghai"
	partner.Active = true
	partner.Type = "contact"
	partner.CountryId = 48
	partner.Mobile = userName
	partner.IsCompany = false
	partner.Color = 0
	partner.PartnerShare = false
	partner.UserId = 1
	memberPartner := model.CreateUserFactory(sqlType).InsertMemberPartner(partner)
	if memberPartner.Id == 0 {
		response.Fail(context, constant.UsersLoginPartnerFailCode, constant.UsersLoginPartnerFailMsg, "")
		return
	}
	user := model.ResUsers{}
	user.Login = userName
	user.Password = pwd
	user.CreateDate = time.Now()
	user.CompanyId = 1
	user.PartnerId = memberPartner.Id
	user.NotificationType = "inbox"
	user.Active = true
	user.Mobile = userName
	user.IsBackstage = true
	memberUsers := model.CreateUserFactory(sqlType).InsertMemberUsers(user)
	if memberUsers.Id == 0 {
		response.Fail(context, constant.UsersLoginUsersFailCode, constant.UsersLoginUsersFailMsg, "")
		return
	}

	resCompanyUsersRel := model.ResCompanyUsersRel{}
	resCompanyUsersRel.Cid = 1
	resCompanyUsersRel.UserId = memberUsers.Id
	model.CreateUserFactory(sqlType).InsertResCompanyUsersRel(resCompanyUsersRel)

	resGroupsUsersRel := model.ResGroupsUsersRel{}
	resGroupsUsersRel.Gid = 1
	resGroupsUsersRel.Uid = memberUsers.Id
	model.CreateUserFactory(sqlType).InsertResGroupsUsersRel(resGroupsUsersRel)

	partnerUserId := model.CreateUserFactory(sqlType).UpdateMemberPartnerUserId(memberPartner.Id, memberUsers.Id)
	if partnerUserId {
		response.Success(context, constant.CurdStatusOkMsg, "")
	} else {
		model.CreateUserFactory(sqlType).DeleteMemberPartnerUserId(memberPartner.Id, memberUsers.Id)
		response.Fail(context, constant.CurdRegisterFailCode, constant.CurdRegisterFailMsg, "")
	}
}

func (u *UsersCurd) UserLogin(login, pass, sqlType string) *model.ResUsers {
	userModel := model.CreateUserFactory(sqlType).Login(login, pass)
	return userModel
}
func (u *UsersCurd) SmsLogin(login, pass, sqlType string, context *gin.Context) *model.ResUsers {
	verifyCode := context.GetString(constant.ValidatorPrefix + "verifyCode")
	code, _ := cache.GetCache(login)
	if code != nil && verifyCode == code {

		pwd, _ := password.SHA512Encrypt.Hash(pass)

		member := model.CreateUserFactory(sqlType).QueryUserLogin(login)
		if member.Id != 0 {
			userModel := model.CreateUserFactory(sqlType).Login(login, pass)
			return userModel
		}

		partner := model.ResPartner{}
		partner.Name = login
		partner.CompanyId = 1
		partner.CreateDate = time.Now()
		partner.DisplayName = login
		partner.Lang = "zh_CN"
		partner.Tz = "Asia/Shanghai"
		partner.Active = true
		partner.Type = "contact"
		partner.CountryId = 48
		partner.Mobile = login
		partner.IsCompany = false
		partner.Color = 0
		partner.PartnerShare = false
		partner.UserId = 1

		memberPartner := model.CreateUserFactory(sqlType).InsertMemberPartner(partner)
		if memberPartner.Id == 0 {
			//response.Fail(context, constant.UsersLoginPartnerFailCode, constant.UsersLoginPartnerFailMsg, "")
			return nil
		}
		user := model.ResUsers{}
		user.Login = login
		user.Password = pwd
		user.CreateDate = time.Now()
		user.CompanyId = 1
		user.PartnerId = memberPartner.Id
		user.NotificationType = "inbox"
		user.Active = true
		user.Mobile = login
		user.IsBackstage = true
		memberUsers := model.CreateUserFactory(sqlType).InsertMemberUsers(user)
		if memberUsers.Id == 0 {
			//response.Fail(context, constant.UsersLoginUsersFailCode, constant.UsersLoginUsersFailMsg, "")
			return nil
		}

		resCompanyUsersRel := model.ResCompanyUsersRel{}
		resCompanyUsersRel.Cid = 1
		resCompanyUsersRel.UserId = memberUsers.Id
		model.CreateUserFactory(sqlType).InsertResCompanyUsersRel(resCompanyUsersRel)

		resGroupsUsersRel := model.ResGroupsUsersRel{}
		resGroupsUsersRel.Gid = 1
		resGroupsUsersRel.Uid = memberUsers.Id
		model.CreateUserFactory(sqlType).InsertResGroupsUsersRel(resGroupsUsersRel)

		partnerUserId := model.CreateUserFactory(sqlType).UpdateMemberPartnerUserId(memberPartner.Id, memberUsers.Id)
		if partnerUserId {
			//response.Success(context, constant.CurdStatusOkMsg, "")
			return nil
		} else {
			model.CreateUserFactory(sqlType).DeleteMemberPartnerUserId(memberPartner.Id, memberUsers.Id)
			//response.Fail(context, constant.CurdRegisterFailCode, constant.CurdRegisterFailMsg, "")
			return nil
		}

	}

	return nil
}

func (u *UsersCurd) RefreshPassword(context *gin.Context) bool {
	pass := context.GetString(constant.ValidatorPrefix + "password")
	accessToken := context.GetString(constant.ValidatorPrefix + "access_token")

	userTokenFactory := token.CreateUserFactory()
	customClaims, _ := userTokenFactory.UserJwt.ParseToken(accessToken)
	uid := customClaims.UserId
	sqlType := customClaims.SqlType

	pwd, _ := password.SHA512Encrypt.Hash(pass)

	return model.CreateUserFactory(sqlType).RefreshPassword(pwd, uid)
}
