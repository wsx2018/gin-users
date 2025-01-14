package main

import (
	"os"
	"user/kernel"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		panic("请输入命令 ./exec mode args...")
	}
	switch args[1] {
	case "http":
		kernel.StartHttp()
	default:
		kernel.Exec(args[1:])
	}
}
