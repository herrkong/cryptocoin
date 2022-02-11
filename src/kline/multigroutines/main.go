package main

import (
	"cryptocoin/src/multigroutines/tasks"
)

var (
	prefix string = "1k"
	ignore_case bool = false
	workers int = 4
)


func main() {
	task := tasks.Tasks{
		Prefix: prefix,
		IgnoreCase: ignore_case,
		Workers: workers,
		ResultChan: make(chan tasks.Result),
		TestChan: make(chan int),
	}
	task.Run()
}