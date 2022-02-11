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

func GetLarkTeaConfig() LarkConfig {
	return LarkConfig{
		AlarmId: "20ef678c-7a06-4660-a7c2-450bd7414fbb",
		Title:   "饮茶先啦",
		Content: "起来活动下,去打一杯茶",
	}
}

func GetLarkGetUpConfig() LarkConfig {
	return LarkConfig{
		AlarmId: "20ef678c-7a06-4660-a7c2-450bd7414fbb",
		Title:   "起床先啦",
		Content: "起床活动下,烧水,洗脸,刷牙",
	}
}

func GetLarkDogConfig() LarkConfig {
	return LarkConfig{
		AlarmId: "20ef678c-7a06-4660-a7c2-450bd7414fbb",
		Title:   "遛狗先啦",
		Content: "起来活动下,带idol去溜溜",
	}
}

func GetLarkLunchConfig() LarkConfig {
	return LarkConfig{
		AlarmId: "20ef678c-7a06-4660-a7c2-450bd7414fbb",
		Title:   "午饭时间到啦",
		Content: "起来活动下,食午饭啦",
	}
}

func GetLarkSleepConfig() LarkConfig {
	return LarkConfig{
		AlarmId: "20ef678c-7a06-4660-a7c2-450bd7414fbb",
		Title:   "睡觉时间到啦",
		Content: "放下手机,睡觉啦",
	}
}

func GetLarkWorkConfig() LarkConfig {
	return LarkConfig{
		AlarmId: "20ef678c-7a06-4660-a7c2-450bd7414fbb",
		Title:   "返工时间到啦",
		Content: "起来活动下,返工时间啦",
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

func GetMySqlConfigNoPwd() MysqlConfig {
	return MysqlConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		PassWord: "",
		Db:       "darwin",
	}
}

func InitAllConfig() {
}
