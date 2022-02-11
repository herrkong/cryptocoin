package tea

// 后台运行
// 上班时间 10点～19点 每个整点提醒一次喝茶

// 定时任务 cron

import (
	"cryptocoin/src/public/config"
	"cryptocoin/src/public/lark"

	"time"

	"github.com/robfig/cron/v3"
)

func TeaRemind() {
	c := cron.New(cron.WithSeconds())
	//每1小时执行一次
	spec := "0/5 * * * * *"

	c.AddFunc(spec, func() {
		TaskForTea()
		TaskSleep()
	})
	c.Start()
	defer c.Stop()
	//阻塞主程序
	select {}
}


func TaskTimer() {
    for {
        now := time.Now()
        // 计算下一个零点
        next := now.Add(time.Hour * 24)
        next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
        t := time.NewTimer(next.Sub(now))
        <-t.C
        //以下为定时执行的操作
        TaskSleep()
    }
}


// 每小时饮茶提醒
func TaskForTea() {
	lark_config := config.GetLarkTeaConfig()
	lark.SendLark(lark_config.AlarmId, lark_config.Title, lark_config.Content)
}

// 11点起床提醒
func TaskGetUp() {
	lark_config := config.GetLarkGetUpConfig()
	lark.SendLark(lark_config.AlarmId, lark_config.Title, lark_config.Content)
}

// 11点30遛狗提醒
func TaskGoForDog() {
	lark_config := config.GetLarkDogConfig()
	lark.SendLark(lark_config.AlarmId, lark_config.Title, lark_config.Content)
}

// 12点30午饭提醒
func TaskForLunch() {
	lark_config := config.GetLarkLunchConfig()
	lark.SendLark(lark_config.AlarmId, lark_config.Title, lark_config.Content)
}

// 13点35上班提醒
func TaskGoForWork() {
	lark_config := config.GetLarkWorkConfig()
	lark.SendLark(lark_config.AlarmId, lark_config.Title, lark_config.Content)
}

// 24点 睡觉提醒
func TaskSleep() {
	lark_config := config.GetLarkSleepConfig()
	lark.SendLark(lark_config.AlarmId, lark_config.Title, lark_config.Content)
}
