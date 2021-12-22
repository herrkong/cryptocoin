package request

import (
	"github.com/beego/beego/v2/client/httplib"
	"fmt"
)

//发送post json请求
func PostJson(url string, body map[string]interface{}) (map[string]interface{}, error) {
	//logs.Info("send post json, ", url, body)
	req := Post(url)
	req.Header("content-type", "application/json")
	req.JSONBody(body)

	var data map[string]interface{}
	err := req.ToJSON(&data)
	fmt.Printf("data=%v\n",data)
	return data, err
}


//重写beego原有http请求的Post方法
func Post(url string) *httplib.BeegoHTTPRequest {
	req := httplib.Post(url)
	return req
}