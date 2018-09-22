package main

import (
	"fmt"

	"github.com/toddlerya/learngo/retriever/mock"
	"github.com/toddlerya/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fakeimooc.com."}
	r = real.Retriever{}
	fmt.Println(download(r))
}
