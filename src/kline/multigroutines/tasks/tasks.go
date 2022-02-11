package tasks

import (
	"bufio"
	btchelper "cryptocoin/src/coin_address/btc"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	//"time"
	//"sync"
)

// var (
// 	prefix      string = "1da" // 靓号前缀
// 	ignore_case bool   = true  // 是否区分大小写
// 	workers     int    = 20    // 协程数
// )

// 一个任务分发协程 解析结果 收到成功的结果 停止分发任务

// 若干个worker协程 处理具体的业务逻辑 计算一次 将计算结果 写入ResultChan

type Tasks struct {
	Prefix     string // 靓号前缀
	IgnoreCase bool   // 是否区分大小写
	Workers    int    // 协程数
	ResultChan chan Result
	TestChan   chan int
}

type Result struct {
	Prikey   string   `json:"privkey"`
	Pubkey   string	  `json:"pubkey"`
	Address  string	  `json:"address"`
}

// 任务处理函数
// 如果完成目标 需要向通道写入true
func (t *Tasks) Handletask()(result Result,err error) {
	btc := new(btchelper.Btc)
	privateKey, publicKey, err := btc.GetBtcKeyPair()
	if err != nil {
		fmt.Printf("Get btc key pair error: %s\n", err)
		return Result{},nil
	}

	prikey := hex.EncodeToString(privateKey.Serialize())
	pubkey := hex.EncodeToString(publicKey.SerializeCompressed())

	//fmt.Printf("prikey= %s\n,pubkey= %s\n", prikey, pubkey)

	chains := []string{"btc", "btctest", "btc_p2wpkh", "btctest_p2wpkh", "btc_p2sh-p2wpkh", "btctest_p2sh-p2wpkh"}
	chain := chains[0]
	address, err := btc.GetBitcoinAddress(publicKey, chain)
	if err != nil {
		fmt.Printf("Get btc address error: %s\n", err)
		return Result{},err
	}
	//fmt.Printf("address= %s\n", address)

	// flag := t.JudgeAddress(address)

	// if !flag {
	// 	return
	// }

	result = Result{
		//Finished: flag,
		Prikey:   prikey,
		Pubkey:   pubkey,
		Address:  address,
	}


	//t.WriteFile(result)

	//fmt.Print(result)
	return result,nil

}

//结构体2jsonstring
func ResultToJson(result * Result) (str string,err error){
	jsonbyte,_ := json.Marshal(*result)
	str = string(jsonbyte)
	return str,nil
}


func (t *Tasks) Run() {
	defer func() {
		if errs := recover(); errs != nil {
			fmt.Printf("errs=%v\n", errs)
		}
	}()

	// 不需要协程同步 直接启动若干个协程dotask
	results := []string{}

	wg := sync.WaitGroup{}

	for i := 0; i < t.Workers; i++ {
		wg.Add(1)
		go func(){
			result ,err := t.Handletask()
			if err != nil{
				fmt.Print(err)
			}
			//fmt.Print(result)
			result_str,_ := ResultToJson(&result)
			results = append(results,result_str)
			wg.Done()
		}()
	}
	wg.Wait()
	//time.Sleep(time.Second * 3)
	fmt.Print(results)
	t.WriteJsonToFile(results)

}

// //发送者
// func sender(c chan int) {
// 	for i := 0; i < 20; i++ {
// 		c <- i
// 		if i >= 5 {
// 			time.Sleep(time.Second * 7)
// 		} else {
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	go sender(c)
// 	timeout := time.After(time.Second * 3)
// 	for {
// 		select {
// 		case d := <-c:
// 			fmt.Println(d)
// 		case <-timeout:
// 			fmt.Println("执行定时操作任务")
// 		case dd := <-time.After(time.Second * 3):
// 			fmt.Println(dd.Format("2006-01-02 15:04:05"), "执行超时动作")
// 		}
// 		fmt.Println("for end")
// 	}
// }

func (t *Tasks) WriteJsonToFile(results []string)(err error){
	result_file, _ := os.OpenFile("multi_groutines_result.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	defer result_file.Close()
	result_file_writer := bufio.NewWriter(result_file)
	defer result_file_writer.Flush()
	for _,v := range results{
		_,err := result_file_writer.WriteString(fmt.Sprintf(v + "\n"))
		if err !=nil {
			return err
		}
	}
	return nil
}

func (t *Tasks) WriteFile(result Result) {
	result_file, _ := os.OpenFile("multi_groutines_result.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	defer result_file.Close()
	result_file_writer := bufio.NewWriter(result_file)
	defer result_file_writer.Flush()
	result_file_writer.Write([]byte(fmt.Sprintf("private_key=%s,pubkey=%s,address=%s\n", result.Prikey, result.Pubkey, result.Address)))
}

func (t *Tasks) JudgeAddress(address string) bool {
	if t.IgnoreCase {
		return address[:len(t.Prefix)] == t.Prefix
	}
	return strings.EqualFold(address[:len(t.Prefix)], t.Prefix)
}

// 初始化

// func (t *Tasks) Init(prefix string, ignore_case bool, workers int) {
// 	t = &Tasks{
// 		Prefix:     prefix,
// 		IgnoreCase: ignore_case,
// 		Workers:    workers,
// 		ResultChan: make(chan Result, 10000),
// 		TestChan:   make(chan int, 10),
// 	}
// }
