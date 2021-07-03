/**
* @Author: lik
* @Date: 2021/3/8 11:34
* @Version 1.0
 */
package chaptcha

import (
	"encoding/json"
	"fmt"
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/app/util/cache"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

type Captcha struct {
	Mobile string `json:"mobile"`
	N      int8   `json:"n"`
}

func (c *Captcha) CheckCode(context *gin.Context) {
	mobile, _ := context.GetQuery("mobile")
	n, _ := context.GetQuery("n")
	N, _ := strconv.ParseInt(n, 10, 8)

	if mobile == "" || len(mobile) != 11 {
		response.Fail(context, constant.CaptchaCheckParamsInvalidCode, constant.CaptchaCheckParamsInvalidMsg, "")
		return
	}
	if N == 0 {
		N = 6
	}
	// 生成验证码
	code := ""
	if N == 4 {
		code = fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	} else if N == 6 {
		code = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	} else {
		response.Fail(context, constant.CaptchaGetParamsInvalidCode, constant.CaptchaGetParamsInvalidMsg, "")
		return
	}
	regionId := variable.ConfigYml.GetString("Sms.RegionId")
	appKey := variable.ConfigYml.GetString("Sms.AppKey")
	appSecret := variable.ConfigYml.GetString("Sms.AppSecret")
	signName := variable.ConfigYml.GetString("Sms.SignName")
	templateCode := variable.ConfigYml.GetString("Sms.TemplateCode")
	client, err := dysmsapi.NewClientWithAccessKey(regionId, appKey, appSecret)
	if err != nil {
		response.Fail(context, constant.CaptchaCheckFailCode, constant.CaptchaCheckFailMsg, "")
		return
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = signName
	request.TemplateCode = templateCode
	request.PhoneNumbers = mobile
	par, _ := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	res, errs := client.SendSms(request)
	if errs != nil {
		response.Fail(context, constant.CaptchaCheckFailCode, constant.CaptchaCheckFailMsg, "")
		return
	}
	if res.Code == "OK" {
		// 缓存验证码 并设置默认过期时间
		cache.SetCache(mobile, code, 2*time.Minute)
		response.Success(context, constant.CaptchaCheckOkMsg, res)
		return
	}

}
