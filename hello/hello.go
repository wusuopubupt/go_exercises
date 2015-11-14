package main

// 导入标准库包
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 快速变量声明
	who := "World!"
	if len(os.Args) > 1 {
		// slice[n:] 返回从第N个元素开始到最后一个元素的切片
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}
