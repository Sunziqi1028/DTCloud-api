package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestStruct(t *testing.T) {
	pe := NewBuilder().
		AddString("Name").
		AddInt64("Age").
		Build()
	p := pe.New()

	//手工赋值
	p.SetString("Name", "你好")
	p.SetInt64("Age", 32)
	//fmt.Printf("%+v\n",p)
	//fmt.Printf("%T，%+v\n",p.Interface(),p.Interface())
	//fmt.Printf("%T，%+v\n",p.Addr(),p.Addr())
	//fmt.Printf("================\n")
	//fmt.Printf("%T,%v,%+v,%#+v\n",p.Addr(),p.Addr(),p.Addr(),p.Addr())

	fmt.Printf("s================\n")
	//JSON赋值
	data := `{"age":18,"name":"标签111"}`
	json.Unmarshal([]byte(data), p.Addr())
	fmt.Printf("%v\n", p.Addr())
	name, err := p.Field("Name")
	if err != nil {
		fmt.Println(name)
	}
	fmt.Println(p.Interface())
	fmt.Println(p.Interface())

	fmt.Printf("e================\n")

	a := 12
	fmt.Printf("%T,%v,%+v,%#+v\n", a, a, a, a)
	fmt.Printf("%T,%v,%+v,%#+v\n", &a, &a, &a, &a)

	a1 := struct {
		id   int    `json:"id"`
		age  int    `json:"age"`
		name string `json:"name"`
	}{
		1,
		2,
		"libai",
	}

	fmt.Printf("%T,%v,%+v,%#+v\n", &a1, &a1, &a1, &a1)
	fmt.Printf("%T,%v,%+v,%#+v\n", a1, a1, a1, a1)

	//
	//data := `{"id":10,"name":"标签111"}`

}

func Test2Struct(t *testing.T) {
	var v interface{}
	data := `{"code":200,"data":[{"id":10,"name":"标签111"},{"id":1,"name":"My Company"},{"id":2,"name":"OdooBot"},{"id":3,"name":"Administrator"}],"msg":"Success"}`
	json.Unmarshal([]byte(data), v)

	t.Log(v)

}

func Test3Saaa(t *testing.T) {
	u := &User{
		id:   1,
		name: "libai",
		age:  18,
	}

	u1 := User{
		id:   2,
		name: "l2",
		age:  18,
	}

	u2 := new(User)
	u2.id = 2
	u2.name = "xiaoli"
	u2.age = 17

	t.Log(u)
	t.Log(u1)
	t.Log(u2)
	fmt.Println("tt=================")
	uType := reflect.TypeOf(u1)
	//uValue := reflect.ValueOf(u1)

	fmt.Printf("%T-%T\n", uType, uType.String())

}
