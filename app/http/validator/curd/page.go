package curd

import (
	"gitee.com/open-product/dtcloud-api/app/global/constant"
	"gitee.com/open-product/dtcloud-api/app/http/controller/curd"
	"gitee.com/open-product/dtcloud-api/app/http/validator/core/data_transfer"
	"gitee.com/open-product/dtcloud-api/app/util/response"
	"github.com/gin-gonic/gin"
)

/*
   'access_token': v['data']['access_token'],  #获取access_token 每天不一样目前
   'model': 'crm.lead',               #当前对象
   'uid': v['data']['uid'],               # 当前用户 uid，      后台二次校验   如果有修改修改帐号停用，并进入黑名单
   'partner_id': v['data']['partner_id'],  # 当前用户partner_id，后台二次校验   如果有修改修改帐号停用，并进入黑名单

   # 'lang': 'zh_CN',                    #当前用户语言
   # 'cache': True,                        #用户缓存调用，用户调用参数可以设置为默认调用5分钟前的数据，后台代码可以控制
   # 'data_type': 'list',                 #返回格式
   #
   'domain': "[('id', '=', 1)]",        #可以选项
   # 'domain': "[('id', '=', 11)]",        #可以选项

   #过滤条件 可以不写
   'search': '',                         #查询关键字
   'search_fields': 'name', #查询关键字
   'fields': 'id,name,user_id,tag_ids',                  #那几个字段
   'offset': 0,                          #从第几页开始
   'limit': 10,                          #每页显示数量
   'order': 'id desc',                   #排序
*/

type PublicPage struct {
	AccessToken string `form:"access_token" json:"access_token" binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	Model       string `form:"model" json:"model" binding:"required,min=1"`
	Uid         int    `form:"uid" json:"uid" binding:"required,min=1"`
	PartnerId   int    `form:"partner_id" json:"partner_id" binding:"required,min=1"`

	Lang     string `form:"lang" json:"lang"`
	Cache    string `form:"cache" json:"cache"`
	DataType string `form:"data_type" json:"data_type"`

	Domain       string `form:"domain" json:"domain"`
	SearchFields string `form:"search_fields" json:"search_fields"`

	Fields       string `form:"fields" json:"fields"`
	FunctionName string `form:"function_name" json:"function_name"`
	Page         int    `form:"page" json:"page"`
	Limit        int    `form:"limit" json:"limit"`
	Order        string `form:"order" json:"order"`
}

func (c PublicPage) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&c); err != nil {
		errs := gin.H{
			"tips": "参数校验失败，参数不符合规定",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(c, constant.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&curd.PageData{}).PulickPage(extraAddBindDataContext)
	}

}
