package main

import (
	"fmt"

	"github.com/mix-go/xcli"
	"github.com/mix-go/xcli/flag"
)

// 采用对象
// type FooCommand struct {
// }
// func (t *FooCommand) Main() {
//     // do something
// }
// cmd := &xcli.Command{
//     Name:  "hello",
//     Short: "Echo demo",
//     RunI:  &FooCommand{},
// }

func main() {
	xcli.SetName("app").SetVersion("0.0.0-alpha")
	cmd := &xcli.Command{
		Name:  "hello",
		Short: "Echo demo",
		//闭包
		RunF: func() {
			name := flag.Match("n", "name").String("default")
			// do something
			fmt.Println(name)
		},
	}
	opt := &xcli.Option{
		Names: []string{"n", "name"},
		Usage: "Your name",
	}
	cmd.AddOption(opt)
	xcli.AddCommand(cmd).Run()
}
