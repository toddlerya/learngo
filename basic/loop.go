package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("==> print from file")
	printFile("basic/basic/branch/abc.txt")
	fmt.Println("\n==> print from NewReader")
	s := `asdfadf
	我哈得发抖
	~！！发生的
	------
	adf sdf
	  dsaf士大夫`
	printFileContents(strings.NewReader(s))
}
