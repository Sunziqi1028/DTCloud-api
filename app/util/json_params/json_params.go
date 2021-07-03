/**
* @Author: lik
* @Date: 2021/3/9 14:02
* @Version 1.0
 */
package json_params

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}

func QueryParams(query url.Values, v *map[string]interface{}) error {

	jsonStr, err := json.Marshal(query)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(jsonStr), &v)

}

func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func SliceValues(values []interface{}) string {
	var paramSlice string
	for _, param := range values {
		switch v := param.(type) {
		case string:
			paramSlice = v
		default:
			panic("params type not supported")
		}

	}

	return paramSlice
}

// 拼接字符串
func GetKeyValue(cf map[string]interface{}) string {
	j := 0
	kv := make([]interface{}, len(cf))
	for k, v := range cf {
		kv[j] = k + "=" + strconv.Quote(v.(string))
		j++
	}

	var paramSlice []string
	for _, param := range kv {
		switch v := param.(type) {
		// TODO
		case string:
			paramSlice = append(paramSlice, v)
		case int:
			strV := strconv.FormatInt(int64(v), 10)
			paramSlice = append(paramSlice, strV)
		default:
			panic("params type not supported")
		}

	}
	result := strings.Replace(strings.Join(paramSlice, ","), "\"", "'", -1)
	return result
}

// key - map value - map 用于sql插入
func GetKeys(cf map[string]interface{}) ([]string, []interface{}) {
	j := 0
	keys := make([]string, len(cf))
	values := make([]interface{}, len(cf))
	for k, v := range cf {
		keys[j] = k
		values[j] = SliceValues(v.([]interface{}))
		j++
	}
	return keys, values
}

// value - string
func GetValues(values []interface{}) string {
	var paramSlice []string
	for _, param := range values {
		switch v := param.(type) {
		case string:
			paramSlice = append(paramSlice, strconv.Quote(v))
		case int:
			strV := strconv.FormatInt(int64(v), 10)
			paramSlice = append(paramSlice, strV)
		default:
			panic("params type not supported")
		}

	}
	var result = strings.Replace(strings.Join(paramSlice, ","), "\"", "'", -1)
	return result
}

func CloneTags(tags map[string]interface{}) map[string]interface{} {
	cloneTags := make(map[string]interface{})
	for k, v := range tags {
		if regexp.MustCompile(`form-field-m2m`).FindAllStringSubmatch(k, -1) != nil {
			cloneTags["form-field-m2m"] = v
		}

	}
	return cloneTags
}

func ImagesToBase64(ff multipart.File) []byte {
	//读原图片
	defer ff.Close()
	sourceBuffer := make([]byte, 804254644)
	n, _ := ff.Read(sourceBuffer)
	//base64压缩
	sourceString := base64.StdEncoding.EncodeToString(sourceBuffer[:n])
	//dist, _ := base64.StdEncoding.DecodeString(sourceString)
	return []byte(sourceString)
}

// 结构体-map
func JSONMethod(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}

// 驼峰 转 下划线 json
type JsonSnakeCase struct {
	Value interface{}
}

func (c JsonSnakeCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

type Person struct {
}

//用map填充结构
func FillStruct(data map[string]interface{}, obj interface{}) error {
	for k, v := range data {
		err := SetField(obj, Case2Camel(k), SliceValues(v.([]interface{})))
		if err != nil {
			return err
		}
	}
	return nil
}

//用map的值替换结构的值
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}

	structFieldValue.Set(val)
	return nil
}

//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	//else if .......增加其他一些类型的转换

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
