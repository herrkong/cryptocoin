package mysql

import (
	"cryptocoin/src/public/config"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/grace"
)

func InitDb() {
	orm.DefaultRowsLimit = -1
	grace.DefaultTimeout = 120 * time.Second
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		panic(err)
	}
	//初始化数据库
	mysql_config := config.GetMySqlConfigNoPwd()
	DataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		mysql_config.User, mysql_config.PassWord, mysql_config.Host, mysql_config.Port, mysql_config.Db)
	if err := orm.RegisterDataBase("darwin", "mysql", DataSource); err != nil {
		panic(err)
	}
	logs.Info("init mysql success.")

}
