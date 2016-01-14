// The type [n]T is an array of n values of type T.
package main

import "fmt"

func main() {
	// a is array of type string, num of elements if 2
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
