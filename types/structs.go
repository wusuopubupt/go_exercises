package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 10
	fmt.Println(v)
	p := &v
	p.X = 11
	fmt.Println(v)
}
