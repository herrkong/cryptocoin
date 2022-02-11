package main

import "fmt"

const (
	KlineType1min   = 1
	KlineType5min   = 5
	KlineType15min  = 15
	KlineType30min  = 30
	KlineType1hour  = 60
	KlineType4hour  = 240
	KlineType12hour = 720
	KlineType1day   = 1440
	KlineType1week  = 10080
)



func main() {
	klineType := int64(10080)

	if klineType == KlineType1week{
		fmt.Print("yes")
	}else{
		fmt.Print("no")
	}

}