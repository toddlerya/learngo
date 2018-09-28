package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/toddlerya/learngo/functional/fib"
)

/*
1. 确保调用在函数结束时发生
2. 参数在defer语句时计算
3. defer列表为后进先出
*/

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
