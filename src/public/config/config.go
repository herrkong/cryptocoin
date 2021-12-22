package config



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

//Lark配置
type LarkConfig struct {
	AlarmId string
	Title   string
	Content string
}

func GetLarkConfig() LarkConfig {
	return LarkConfig{
		AlarmId: "e070e5e4-7a22-4302-b1d9-60d5a4dd8503",
		Title:   "饮茶先啦 提醒第一次",
		Content: "起来活动下,去打一杯茶",
	}
}

func GetLarkConfig2() LarkConfig {
	return LarkConfig{
		AlarmId: "e070e5e4-7a22-4302-b1d9-60d5a4dd8503",
		Title:   "饮茶先啦 提醒第二次",
		Content: "起来活动下,去打一杯茶",
	}
}



func GetRedisConfig() RedisConfig {
	return RedisConfig{
		Host:     "127.0.0.1",
		Port:     "6379",
		PassWord: "darwin",
	}
}

func GetMySqlConfig() MysqlConfig {
	return MysqlConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "darwin",
		PassWord: "loginmysql",
		Db:       "testdb",
	}
}

func InitAllConfig() {
}
