package main

import "fmt"

func main() {
	// new
	s := new(string)
	fmt.Println("%s\n", s)
	fmt.Println(*s)

	// make
	// 内建函数make(T, args)与new(T)的用途不一样。它只用来创建slice，map和channel，
	// 并且返回一个初始化的(而不是置零)，类型为T的值（而不是*T）
	v := make([]int, 100)

}
