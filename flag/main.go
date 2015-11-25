// go build main.go
// ./main -c ../conf_path -l ../log_path -id 24 -ok true
// ./main -h
package main

import (
	"flag"
	"fmt"
)

func main() {
	// refer : http://www.01happy.com/golang-command-line-arguments/
	// 方法一： flag.StringVar(),传入指针，直接给confPath赋值
	var confPath string
	flag.StringVar(&confPath, "c", "../conf", "config file path")
	// 方法二：flag.String()，把函数调用的返回值赋值给logPath
	logPath := flag.String("l", "../log", "log file path")

	id := flag.Int("id", 1, "id")
	ok := flag.Bool("ok", false, "is ok")

	flag.Parse()

	fmt.Println("conf path: ", confPath)
	fmt.Println("log path: ", *logPath)
	fmt.Println("id is", *id)
	fmt.Println("ok is", *ok)
}
