package redis_test

import (
	"cryptocoin/src/public/redis"
	"fmt"
)

func redis_test() {

	redis.Init()

	reply,err := redis.Do("0", "keys","*")

	if err != nil{
		fmt.Print(err)
	}

	redis.Parse(reply)
	
}