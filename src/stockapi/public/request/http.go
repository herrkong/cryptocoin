package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func NewHttpRequest(client *http.Client, reqType string, reqUrl string, postData string, requstHeaders map[string]string) ([]byte, error) {
	req, err := http.NewRequest(reqType, reqUrl, strings.NewReader(postData))
	if err != nil {
		return nil, err
	}
	fmt.Printf("send reqeust=%v\n", req)

	for k, v := range requstHeaders {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http request failed, url:%s, err:%s", reqUrl, err)
		return nil, err
	}
	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read http body failed, url:%s, err:%s", reqUrl, err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HttpStatusCode:%d ,Desc:%s", resp.StatusCode, string(bodyData)))
	}

	return bodyData, nil
}

func HttpGet(client *http.Client, reqUrl string) (map[string]interface{}, error) {
	respData, err := NewHttpRequest(client, "GET", reqUrl, "", nil)
	if err != nil {
		return nil, err
	}

	var bodyDataMap map[string]interface{}
	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		return nil, err
	}

	return bodyDataMap, nil
}

func HttpGet2(client *http.Client, reqUrl string, headers map[string]string) ([]map[string]interface{}, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	//headers["Content-Type"] = "application/x-www-form-urlencoded"
	respData, err := NewHttpRequest(client, "GET", reqUrl, "", headers)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("respData=%v\n", respData)

	var bodyDataMap []map[string]interface{}

	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		return nil, err
	}

	return bodyDataMap, nil
}

func HttpGet3(client *http.Client, reqUrl string, headers map[string]string) (map[string]interface{}, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	//headers["Content-Type"] = "application/x-www-form-urlencoded"
	respData, err := NewHttpRequest(client, "GET", reqUrl, "", headers)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("respData=%v\n", respData)

	var bodyDataMap map[string]interface{}

	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		return nil, err
	}

	return bodyDataMap, nil
}

func HttpGet4(client *http.Client, reqUrl string, headers map[string]string) (map[string]interface{}, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	//headers["Content-Type"] = "application/x-www-form-urlencoded"
	respData, err := NewHttpRequest(client, "POST", reqUrl, "", headers)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("respData=%v\n", respData)

	var bodyDataMap map[string]interface{}

	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		return nil, err
	}

	return bodyDataMap, nil
}

func HttpDelete(client *http.Client, reqUrl string, postData url.Values, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	//headers["Content-Type"] = "application/x-www-form-urlencoded"
	return NewHttpRequest(client, "DELETE", reqUrl, postData.Encode(), headers)
}
