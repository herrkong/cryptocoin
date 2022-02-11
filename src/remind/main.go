package main

import (
	"cryptocoin/src/remind/tea"
)

func main() {
	tea.TeaRemind()
	//tea.TaskTimer()
	// if IsNowInTimeRange("12:01:00", "12:06:00"){
	// 	fmt.Print("in")
	// }else{
	// 	fmt.Println("out")
	// }

}

// //当前时间是否在指定范围内
// //参数为时间字符串，格式为"时:分:秒"
// func IsNowInTimeRange(startTimeStr, endTimeStr string) bool {
// 	//当前时间
// 	now := time.Now()
// 	//当前时间转换为"年-月-日"的格式
// 	format := now.Format("2006-01-02")
// 	//转换为time类型需要的格式
// 	layout := "2006-01-02 15:04:05"
// 	//将开始时间拼接“年-月-日 ”转换为time类型
// 	timeStart, _ := time.ParseInLocation(layout, format+" "+startTimeStr, time.Local)
// 	//将结束时间拼接“年-月-日 ”转换为time类型
// 	timeEnd, _ := time.ParseInLocation(layout, format+" "+endTimeStr, time.Local)
// 	//使用time的Before和After方法，判断当前时间是否在参数的时间范围
// 	return now.Before(timeEnd) && now.After(timeStart)
// }
