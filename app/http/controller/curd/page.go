package curd

import (
	"fmt"
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/service/curd"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

type PageData struct {
}

func (this PageData) PulickPage(context *gin.Context) {
	uid := int(context.GetFloat64(constant.ValidatorPrefix + "uid"))
	partnerId := int(context.GetFloat64(constant.ValidatorPrefix + "partner_id"))
	table := strings.Replace(context.GetString(constant.ValidatorPrefix+"model"), ".", "_", -1)

	domain := context.GetString(constant.ValidatorPrefix + "domain")

	search := context.GetString(constant.ValidatorPrefix + "search")
	searchfields := context.GetString(constant.ValidatorPrefix + "search_fields")

	fields := context.GetString(constant.ValidatorPrefix + "fields")
	offset := int(context.GetFloat64(constant.ValidatorPrefix + "page"))
	limit := int(context.GetFloat64(constant.ValidatorPrefix + "limit"))
	order := context.GetString(constant.ValidatorPrefix + "order")

	fmt.Print(uid, partnerId, search, searchfields, fields, offset, limit, order, table)

	//type Domain struct{
	//	filed1 string
	//	sign string
	//	value interface{}
	//}
	//a1 := []string{}
	//json.Unmarshal([]byte(domain),a1)
	//
	//a := "[('id','=',1),('id','=',1),('id','=',1),('id','=',1)]"
	//
	//b := []int{1,2,3,4,5}
	//fmt.Println(a)
	//fmt.Println(a[0])
	//fmt.Println(a[1])
	//fmt.Println(b)
	////a11 := &Domain{}
	////fmt.Println(domain)
	////json.Unmarshal([]byte(domain),a11)
	//FindAllSubmatch
	//k:= regexp.MustCompile(`/\\((.+?)\\)/g`).FindAllStringSubmatch(domain, -1)
	//

	for _, match := range regexp.MustCompile("\\((.+?)\\)").FindAllString(domain, -1) {
		fmt.Printf("%v\n", match)
	}

	//for k,v := range domain{
	//	fmt.Println(k,v)
	//	a1 := v.(Domain)
	//	//strings.Split(domain)
	//}

	result := curd.CreateUserCurdFactory().PublicPage(uid, partnerId, domain, search, searchfields, fields, offset, limit, order, table)
	if result != nil {
		response.Success(context, constant.CurdStatusOkMsg, result)
		return
	}
	response.Fail(context, constant.CurdSelectFailCode, constant.CurdSelectFailMsg, "")
	return

}
