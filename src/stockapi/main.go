package main

import (
	"cryptocoin/src/stockapi/hk"
	"fmt"
)

func main() {
	h := hk.NewHongKong()
	params_map := hk.SetParamsMap(hk.BILIBILI,string(hk.Real))
	stock_info, err := h.GetBiliBiliInfo(params_map)
	if err != nil {
		fmt.Printf("get stock info,err = %s\n", err)
	}
	fmt.Println(stock_info)

}
