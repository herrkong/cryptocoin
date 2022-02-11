package main

import (
	"fmt"
)

type ListData struct {
	Value int
	Str   string
}

func main() {
	list1 := make([]*ListData, 0)
	d1 := new(ListData)
	d1.Value = 1
	d1.Str = "e"

	list1 = append(list1, d1)

	d2 := new(ListData)
	d2.Value = 2
	d2.Str = "fe"

	list1 = append(list1, d2)
	
	templist := list1

	TestList(&templist)

}

func UpdateList() []*ListData {
	list1 := make([]*ListData, 0)
	d := new(ListData)
	d.Value = 1
	d.Str = "e"
	list1 = append(list1, d)
	return list1
}

func TestList(List *[]*ListData) {
	
	fmt.Println(*List)

}

// 这里确保list里的数据都被清空

// func (w *Worker) UpdateCurKline(updateKline func(tc *[]*tablemodels.Trade)) {
// 	go func() {
// 		ticker := time.Tick(time.Second)
// 		for {
// 			<-ticker
// 			TC.Lock()
// 			tmpCache := TC.List
// 			TC.List = TradeCachePool.Get().([]*tablemodels.Trade)
// 			TC.Unlock()

// 			go func() {
// 				updateKline(&tmpCache)
// 				tmpCache = tmpCache[0:0]
// 				TradeCachePool.Put(tmpCache)
// 			}()
// 		}
// 	}()
// }
