package main

import (
	"cryptocoin/src/public/redis"
	"fmt"
)

func main() {

	redis.Init()

	reply,err := redis.Do("0", "keys","*")

	if err != nil{
		fmt.Print(err)
	}

	redis.Parse(reply)
	
}