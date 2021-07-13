/**
* @Author: lik
* @Date: 2021/3/7 18:10
* @Version 1.0
 */
package user

type Base struct {
	Type string `form:"type" json:"type"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
}
