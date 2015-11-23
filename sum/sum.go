package main

import (
	"fmt"
)

func sum(a []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	// 创建数组
	arr := []int{1, 2, 3, 4, 5}
	s := sum(arr)
	fmt.Println(s)
}
