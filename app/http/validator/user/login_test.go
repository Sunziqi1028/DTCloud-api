package user

import (
	"encoding/json"
	"fmt"
	"gitee.com/open-product/dtcloud-api/app/util"
	"testing"
)

func TestLogin_CheckParams(t *testing.T) {
	var (
		bl   int = 1
		gUrl string
	)

	if bl == 1 {
		gUrl = "http://122.51.164.176:8072/api/v1/login/0"
	} else {
		gUrl = "http://127.0.0.1:8072/api/v1/login/0"
	}
	this := &util.RequestInfo{
		Url: gUrl,
		Data: map[string]interface{}{
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
}
