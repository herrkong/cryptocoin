package hk

import (
	request "cryptocoin/src/stockapi/public/request"
	"fmt"
	"net/http"
)

type HongKong struct {
	HttpClient *http.Client
	Url        string
}

type StockInfo struct {
	StockName string
	MarkPrice string
}

type MethodName string

const (
	Real MethodName = "real"
)

const (
	BILIBILI = "9626.HKM"
)

type ParamsMap map[string]interface{}

func NewHongKong() *HongKong {
	client := &http.Client{}
	url := "https://sandbox.hscloud.cn/quote/v1/"
	return &HongKong{
		HttpClient: client,
		Url:        url,
	}
}

func SetParamsMap(en_prod_code string, method_name string) ParamsMap {
	params_map := make(ParamsMap)
	params_map["method_name"] = method_name
	params_map["en_prod_code"] = en_prod_code
	return params_map
}

//method name :real
//param :
//en_prod_code : 代码
//example : /quote/v1/real?en_prod_code=00001.HKM
func (h *HongKong) GetBiliBiliInfo(ParamsMap ParamsMap) (StockInfo *StockInfo, err error) {
	en_prod_code := ParamsMap["en_prod_code"].(string)
	method_name := ParamsMap["method_name"].(string)
	path := fmt.Sprintf("%s%s?%s", h.Url, method_name, en_prod_code)
	respmap, err := request.HttpGet(h.HttpClient, path)
	if err != nil {
		return nil, err
	}
	fmt.Printf("respmap= %v\n", respmap)

	return StockInfo, nil
}
