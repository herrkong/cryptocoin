package redis

import (
	"cryptocoin/src/public/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

var (
	redis_pool *redis.Pool
)


func Init(){
	//实例化一个连接池
	redis_config := config.GetRedisConfig()
	address := fmt.Sprintf("%s:%s",redis_config.Host,redis_config.Port)
	redis_pool = &redis.Pool{
		//最初的连接数量
		MaxIdle: 16,
		//最大连接数量
		//连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		MaxActive: 2000,
		//连接关闭时间 300秒 （300秒不使用自动关闭）
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			//要连接的redis数据库
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err	
			}
			if _, err := c.Do("AUTH", redis_config.PassWord); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	fmt.Println("init redis success.")
}	


func Do(db string,command_name string ,args ...interface{})(reply interface{},err error){
	conn := redis_pool.Get()
	defer conn.Close()

	_, err = conn.Do("SELECT",db)

	if err != nil {
		return nil,err
	}

	return conn.Do(command_name,args...)

}


//数据解析
func Parse(item interface{}) (data map[string]interface{}) {
	dataBytes := item.([]interface{})[0].([]byte)
	fmt.Print(string(dataBytes))
	if err := json.Unmarshal(dataBytes, &data); err == nil {
		return data
	}
	return nil
}
