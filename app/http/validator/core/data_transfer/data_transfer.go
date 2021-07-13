/**
* @Author: lik
* @Date: 2021/3/7 18:25
* @Version 1.0
 */
package data_transfer

import (
	"encoding/json"
	"gitee.com/open-product/dtcloud-api/app/http/validator/core/interf"
	"github.com/gin-gonic/gin"
)

/*
   bool, for JSON booleans
   float64, for JSON numbers
   string, for JSON strings
   []interface{}, for JSON arrays
   map[string]interface{}, for JSON objects
   nil for JSON null
*/

// 将验证器成员(字段)绑定到数据传输到上下文，方便控制器获取
/**
本函数参数说明：
validatorInterface 实现了验证器接口的结构体
extra_add_data_prefix  验证器绑定参数传递给控制器的数据前缀
context  gin上下文
*/
func DataAddContext(validatorInterface interf.ValidatorInterface, extraAddDataPrefix string, context *gin.Context) *gin.Context {
	var tempJson interface{}
	if tmpBytes, err1 := json.Marshal(validatorInterface); err1 == nil {
		if err2 := json.Unmarshal(tmpBytes, &tempJson); err2 == nil {
			if value, ok := tempJson.(map[string]interface{}); ok {
				for k, v := range value {
					context.Set(extraAddDataPrefix+k, v)
				}
				return context
			}
		}
	}
	return nil
}
