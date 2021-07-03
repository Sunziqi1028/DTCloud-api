/**
* @Author: lik
* @Date: 2021/3/5 20:02
* @Version 1.0
 */
package interfconfig

import (
	"time"
)

type YmlConfig interface {
	ConfigFileChangeListen()
	Clone(fileName string) YmlConfig
	Get(keyName string) interface{}
	GetString(keyName string) string
	GetBool(keyName string) bool
	GetInt(keyName string) int
	GetInt32(keyName string) int32
	GetInt64(keyName string) int64
	GetFloat64(keyName string) float64
	GetDuration(keyName string) time.Duration
	GetStringSlice(keyName string) []string
}
