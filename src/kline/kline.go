package kline

import (
	"fmt"
	"time"
)

const (
	KlineType1min   = 1
	KlineType5min   = 5
	KlineType15min  = 15
	KlineType30min  = 30
	KlineType1hour  = 60
	KlineType4hour  = 240
	KlineType12hour = 720
	KlineType1day   = 1440
	KlineType1week  = 10080
)

type SaveMtime struct {
	Mtime int64
	Type  int
}


// SaveKline 落地k线
func SaveKline(saveKline func([]SaveMtime)) {
	timeUnix := time.Now().Unix()
	// 默认当前时区为+8
	var calcTimeDelta int64 = 8 * 3600

	mtime1Min := timeUnix - timeUnix%60
	mtime5Min := timeUnix - (timeUnix+calcTimeDelta)%(5*60)
	mtime15Min := timeUnix - (timeUnix+calcTimeDelta)%(15*60)
	mtime30Min := timeUnix - (timeUnix+calcTimeDelta)%(30*60)
	mtime1Hour := timeUnix - (timeUnix+calcTimeDelta)%(60*60)
	mtime4Hour := timeUnix - (timeUnix+calcTimeDelta)%(240*60)
	mtime12Hour := timeUnix - (timeUnix+calcTimeDelta)%(720*60)
	mtime1Day := timeUnix - (timeUnix+calcTimeDelta)%(1440*60)
	next := mtime1Min + 60
	timer := time.NewTimer(time.Unix(next, 0).Sub(time.Now()))
	for {
		select {
		case <-timer.C:
			saveList := make([]SaveMtime, 0)

			// 1min
			saveList = append(saveList, SaveMtime{
				Mtime: mtime1Min,
				Type:  KlineType1min,
			})
			mtime1Min = next

			// 5min
			if next%(5*60) == 0 {
				saveList = append(saveList, SaveMtime{
					Mtime: mtime5Min,
					Type:  KlineType5min,
				})
				mtime5Min = next
			}

			// 15min
			if next%(15*60) == 0 {
				saveList = append(saveList, SaveMtime{
					Mtime: mtime15Min,
					Type:  KlineType15min,
				})
				mtime15Min = next
			}

			// 30min
			if next%(30*60) == 0 {
				saveList = append(saveList, SaveMtime{
					Mtime: mtime30Min,
					Type:  KlineType30min,
				})
				mtime30Min = next
			}

			// 1hour
			if next%(60*60) == 0 {
				saveList = append(saveList, SaveMtime{
					Mtime: mtime1Hour,
					Type:  KlineType1hour,
				})
				mtime1Hour = next
			}

			// 4hour
			if next%(240*60) == 0 {
				saveList = append(saveList, SaveMtime{
					Mtime: mtime4Hour,
					Type:  KlineType4hour,
				})
				mtime4Hour = next
			}

			// 12hour
			if next%(720*60) == 0 {
				saveList = append(saveList, SaveMtime{
					Mtime: mtime12Hour,
					Type:  KlineType12hour,
				})
				mtime12Hour = next
			}

			// 1day
			if next%(1440*60) == 0 {
				saveList = append(saveList, SaveMtime{
					Mtime: mtime1Day,
					Type:  KlineType1day,
				})
				mtime1Day = next
			}

			// 1week
			fmt.Print(saveList)

			saveKline(saveList)

			next += 60
			timer = time.NewTimer(time.Unix(next, 0).Sub(time.Now()))
		}
	}
}