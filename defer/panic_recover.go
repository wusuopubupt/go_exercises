package main

import "fmt"

func f() {
	fmt.Println("1")
	panic("exception message")
	// panic后面的代码不再执行
	fmt.Println("2")
}

/**
输出：

1
3
exception message
4
 */
func main() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("3")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("4")
	}()

	f()
}


