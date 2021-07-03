/**
* @Author: lik
* @Date: 2021/3/9 11:28
* @Version 1.0
 */
package curd

import (
	"gitee.com/open-product/dtcloud-api/model"
	"gitee.com/open-product/dtcloud-api/routers/middleware/my_jwt"
)

func CreateUserCurdFactory() *UsersCurd {

	return &UsersCurd{}
}

type UsersCurd struct {
}

func (u *UsersCurd) PublicCreate(v map[string]interface{}, dat map[string]interface{}, claims *my_jwt.CustomClaims) bool {

	return model.CreateUserFactory(claims.SqlType).CreateParamsField(v, dat)
}

func (u *UsersCurd) PublicRead(fields string, ids string, page int64, limit int64, order string, table string) []map[string]interface{} {

	return model.CreateUserFactory("postgresql").PublicRead(fields, ids, page, limit, order, table)
}
