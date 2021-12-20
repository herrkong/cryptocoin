package config

import(

)


//Redis配置
type RedisConfig struct {
	Host     string
	Port     string
	PassWord string
}


//数据库配置
type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	PassWord string
	Db       string
}

func GetRedisConfig() RedisConfig{
	return RedisConfig{
		Host: "127.0.0.1",
		Port: "6379",
		PassWord: "darwin",
	}
}

func GetMySqlConfig() MysqlConfig{
	return MysqlConfig{

	}
}


func InitAllConfig(){
}