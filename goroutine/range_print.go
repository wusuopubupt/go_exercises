package main

import "fmt"
import "time"

var values = []string{"a", "b", "c"}

func main() {
	for _, v := range values {
		/*
			输出：
			c
			c
			c
			原因：
			只是将goroutine放到调度队列，而不是立即执行
		*/
		go func() {
			fmt.Println(v) //隐式的调用v的地址传递
		}()
		/*
			输出：
			a
			b
			c
		*/
		/*
			go func(s string) {
				fmt.Println(s)
			}(v) //明确复制值，作为栈变量传入goroutine
		*/
	}
	time.Sleep(time.Second)
}
