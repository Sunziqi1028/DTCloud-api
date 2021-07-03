/**
* @Author: lik
* @Date: 2021/3/7 15:18
* @Version 1.0
 */
package gorm_v2

type ConfigParams struct {
	Write ConfigParamsDetail
	Read  ConfigParamsDetail
}
type ConfigParamsDetail struct {
	Host     string
	DataBase string
	Port     int
	Prefix   string
	User     string
	Pass     string
	Charset  string
}
