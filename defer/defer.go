package main

import "fmt"
import "os"
import "io"

func copyFile(src string, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	defer func() {
		srcFile.Close()
	}()
	if err != nil {
		return
	}

	dstFile, err := os.Create(dst)
	defer func() {
		dstFile.Close()
	}()
	if err != nil {
		return
	}

	return io.Copy(dstFile, srcFile)
}

func fn(n int) int {
	// defer: FILO
	defer func() {
		n++
		fmt.Println("3rd", n)
	}()

	defer func() {
		n++
		fmt.Println("2nd", n)
	}()

	defer func() {
		n++
		fmt.Println("1st", n)
	}()
	// 不改变值
	return n
}

func main() {
	copyFile("./a.txt", "./b.txt")
	fmt.Println("function return: ", fn(0))
}
