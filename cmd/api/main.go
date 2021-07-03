/**
* @Author: lik
* @Date: 2021/3/5 11:41
* @Version 1.0
 */
package main

import (
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	_ "gitee.com/open-product/dtcloud-api/bootstrap"
	"gitee.com/open-product/dtcloud-api/routers"
)

func main() {
	router := routers.InitApiRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Api.Port"))
}
