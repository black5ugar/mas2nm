package main

import (
	"flag"
	"fmt"
	"mas2nm/module"
)

var (
	input = flag.String("i", "ip.txt", "ip输入文件")
)

func main() {
	flag.Parse()
	if err := module.PortScan(*input); err != nil {
		fmt.Println(err)
	}
}
