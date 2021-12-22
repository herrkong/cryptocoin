package main

import (
	"cryptocoin/src/public/config"
	"cryptocoin/src/public/lark"
)

func main() {
	config := config.GetLarkConfig()
	lark.SendLark(config.AlarmId,config.Title,config.Content)
}