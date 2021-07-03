/**
* @Author: lik
* @Date: 2021/3/5 19:09
* @Version 1.0
 */
package model

import (
	"time"
)

type ResUsers struct {
	Id               int64     `json:"id"`
	Login            string    `json:"login"`
	Password         string    `json:"password"`
	CreateDate       time.Time `json:"create_date"`
	CompanyId        int64     `json:"company_id"`
	PartnerId        int64     `json:"partner_id"`
	NotificationType string    `json:"notification_type"`
	Mobile           string    `json:"mobile"`
	Active           bool      `json:"active"`
	IsBackstage      bool      `json:"is_backstage"`
}

type ResPartner struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	CompanyId    int64     `json:"company_id"`
	CreateDate   time.Time `json:"create_date"`
	DisplayName  string    `json:"display_name"`
	Lang         string    `json:"lang"`
	Tz           string    `json:"tz"`
	Active       bool      `json:"active"`
	Type         string    `json:"type"`
	CountryId    int64     `json:"country_id"`
	UserId       int64     `json:"user_id"`
	Mobile       string    `json:"mobile"`
	IsCompany    bool      `json:"is_company"`
	Color        int16     `json:"color"`
	PartnerShare bool      `json:"partner_share"`
}

type ResGroupsUsersRel struct {
	Gid int64 `json:"gid"`
	Uid int64 `json:"uid"`
}

type ResCompanyUsersRel struct {
	Cid    int64 `json:"cid"`
	UserId int64 `json:"user_id"`
}
