package main

import (
	"fmt"

	"github.com/khunmk/uni-zg-converter/cmd/wasm/tools"
)

func main() {
	c := make(chan bool)

	fmt.Println("wasm initialized...")

	tools.RegisterCallBack()

	<-c
}
