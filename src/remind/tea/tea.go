package tea

// 后台运行
// 上班时间 10点～19点 每个整点提醒一次喝茶

// 定时任务 cron

import (
	"cryptocoin/src/public/config"
	"cryptocoin/src/public/lark"
	//"sync"

	//"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func TeaRemind() {
	c := cron.New(cron.WithSeconds())
	//每1小时执行一次
	c.AddFunc("0 0 * * * *", func(){
		Task1()
	})
	c.Start()
	//time.Sleep(300 * time.Minute)

	//阻塞主程序
	t := time.NewTimer(time.Second * 10)
	for{
		select{
		case <-t.C:
			t.Reset(time.Second * 10)
		}
	}
}

func Task1() {
	lark_config := config.GetLarkConfig()
	lark.SendLark(lark_config.AlarmId, lark_config.Title, lark_config.Content)
}

func Task2(){
	lark_config := config.GetLarkConfig2()
	lark.SendLark(lark_config.AlarmId, lark_config.Title, lark_config.Content)
}

func Timer(){
	t := time.NewTimer(time.Second * 10)
	for{
		select{
		case <-t.C:
			t.Reset(time.Second * 10)
			Task2()
		}
	}
}