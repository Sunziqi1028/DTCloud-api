package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestInfo struct {
	Url           string
	Data          map[string]string //post要传输的数据，必须key value必须都是string
	DataInterface map[string]interface{}
}

//适用于 application/x-www-form-urlencoded
func (this RequestInfo) PostUrlEncoded() []byte {
	client := &http.Client{}
	//post要提交的数据
	DataUrlVal := url.Values{}
	for key, val := range this.Data {
		DataUrlVal.Add(key, val)
	}
	req, err := http.NewRequest("POST", this.Url, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//伪装头部
	//req.Header.Set("Accept","application/json, text/javascript, */*; q=0.01")
	//req.Header.Add("Accept-Encoding","gzip, deflate, br")
	//req.Header.Add("Accept-Language","zh-CN,zh;q=0.8,en-US;q=0.6,en;q=0.4")
	//req.Header.Add("Connection","keep-alive")
	//req.Header.Add("Content-Length","25")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("Cookie","user_trace_token=20170425200852-dfbddc2c21fd492caac33936c08aef7e; LGUID=20170425200852-f2e56fe3-29af-11e7-b359-5254005c3644; showExpriedIndex=1; showExpriedCompanyHome=1; showExpriedMyPublish=1; hasDeliver=22; index_location_city=%E5%85%A8%E5%9B%BD; JSESSIONID=CEB4F9FAD55FDA93B8B43DC64F6D3DB8; TG-TRACK-CODE=search_code; SEARCH_ID=b642e683bb424e7f8622b0c6a17ffeeb; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1493122129,1493380366; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1493383810; _ga=GA1.2.1167865619.1493122129; LGSID=20170428195247-32c086bf-2c09-11e7-871f-525400f775ce; LGRID=20170428205011-376bf3ce-2c11-11e7-8724-525400f775ce; _putrc=AFBE3C2EAEBB8730")
	//req.Header.Add("Host","www.lagou.com")
	//req.Header.Add("Origin","https://www.lagou.com")
	//req.Header.Add("Referer","https://www.lagou.com/jobs/list_python?labelWords=&fromSearch=true&suginput=")
	//req.Header.Add("X-Anit-Forge-Code","0")
	//req.Header.Add("X-Anit-Forge-Token","None")
	//req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	//req.Header.Add("X-Requested-With","XMLHttpRequest")
	//提交请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//读取返回值
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(result))
	return result
}
