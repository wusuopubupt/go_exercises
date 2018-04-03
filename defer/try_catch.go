package main

import "fmt"

// defer: 函数结束时执行的语句
// panic : 抛RuntimeException
// recover : 捕获处理异常

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main() {
	// 抛出一个panic的异常，然后在defer中通过recover捕获这个异常
	Try(func() {
		panic("exception message")
	}, func(e interface{}) {
		fmt.Println(e)
	})
}
