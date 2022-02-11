package map_save_struct


import ()

//map[int]struct {}不能修改struct内的属性
//传入指针能修改结构体的值

type Kline struct {
	Open       string `gorm:"column:open" json:"open"`
	Close      string `gorm:"column:close" json:"close"`
	High       string `gorm:"column:high" json:"high"`
	Low        string `gorm:"column:low" json:"low"`
	Volume     string `gorm:"column:volume" json:"volume"`
}


func TestMapSaveStruct(){
	
	m := make(map[int64]*Kline)




}