package main

import (
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService()
	service.Init()

	for i := 0; i < 10; i++ {
		call(i, service.Client())
	}

}
