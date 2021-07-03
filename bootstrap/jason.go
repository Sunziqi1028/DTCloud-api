package bootstrap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	base        string = "http://122.51.164.176:8072"
	contentType string = "application/x-www-form-urlencoded"
)

func check() {
	url := base + "/api/v1/login/0"
	//login:admin
	//password:123
	//type:0
	data := `login=admin&password=123&type=0`
	res, err := http.Post(
		url,
		contentType,
		strings.NewReader(data),
	)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
