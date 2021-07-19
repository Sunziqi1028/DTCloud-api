package curd

import (
	"encoding/json"
	"fmt"
	"gitee.com/open-product/dtcloud-api/app/http/controller/curd"
	"gitee.com/open-product/dtcloud-api/app/util"
	"gitee.com/open-product/dtcloud-api/app/util/odoo"
	"strings"
	"testing"
)

func TestPublicPage_CheckParams(t *testing.T) {
	var (
		bl   int = 2
		gUrl string
	)

	switch bl {
	case 0:
		gUrl = "http://122.51.164.176:8072/api/v1/login/0"
	case 1:
		gUrl = "http://127.0.0.1:8072/api/v1/login/0"
	case 2:
		gUrl = "http://127.0.0.1:8890/api/v1/login/0"

	}

	this := &util.RequestInfo{
		Url: gUrl,
		Data: map[string]string{
			"login":    "admin",
			"password": "1",
			"type":     "0",
		},
		DataInterface: map[string]interface{}{
			"ContentType": "application/x-www-form-urlencoded",
		},
	}
	body := this.PostUrlEncoded()
	fmt.Println(string(body))

	//whitelist := map[string]map[string]int{}
	m1 := make(map[string]interface{})
	err := json.Unmarshal(body, &m1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m1)
	fmt.Println(m1["data"])
	m2 := m1["data"].(map[string]interface{})
	fmt.Println(m2)
	fmt.Println(m2["access_token"])
	//fmt.Println(map2["data"])

	gUrl = "http://127.0.0.1:8890/api/v1/page/0"

	page := &util.RequestInfo{
		Url: gUrl,
		Data: map[string]string{
			"access_token": m2["access_token"].(string),
			"model":        "crm.lead",
			"uid":          "2",
			"partner_id":   "3",
			"domain":       "[('id', '=', 1)]",

			"search":        "",
			"search_fields": "name",                   // #查询关键字
			"fields":        "id,name,user_id,tag_id", //                #那几个字段
			"offset":        "0",                      //#从第几页开始
			"limit":         "10",                     //#每页显示数量
			"order":         "id desc",                //#排序

		},
		DataInterface: map[string]interface{}{
			"ContentType": "application/x-www-form-urlencoded",
		},
	}
	body = page.PostUrlEncoded()
	fmt.Println(string(body))
}
