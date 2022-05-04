package main

import (
	"errors"
	"fmt"
	"log"

	"syscall/js"

	"github.com/khunmk/uni-zg-converter/utils"
)

func main() {
	c := make(chan bool, 0)

	fmt.Println("Wasm initialized")

	document := js.Global().Get("document")
	if !document.Truthy() {
		log.Fatalln("unable to get document")
	}

	uniInputEle := document.Call("getElementById", "uni-input")
	if !uniInputEle.Truthy() {
		log.Fatalln("unable to get element")
	}

	zgInputEle := document.Call("getElementById", "zg-input")
	if !zgInputEle.Truthy() {
		log.Fatalln("unable to get element")
	}

	uniInputEle.Call("addEventListener", "input", toZg(zgInputEle))
	zgInputEle.Call("addEventListener", "input", toUni(uniInputEle))
	js.Global().Set("convertTo", convertTo())

	<-c
}

func toUni(outputEle js.Value) js.Func {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {

		inputStr := this.Get("value").String()

		outputTxt, err := utils.Zg2uni(inputStr)
		if err != nil {
			return errors.New("unable to convert")
		}

		outputEle.Set("value", outputTxt)

		return nil
	})

	return cb
}

func toZg(outputEle js.Value) js.Func {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		inputStr := this.Get("value").String()

		outputTxt, err := utils.Uni2zg(inputStr)
		if err != nil {
			return errors.New("unable to convert")
		}

		outputEle.Set("value", outputTxt)

		return nil
	})

	return cb
}

/**
 * Todo :::
 */
func convertTo() js.Func {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		// fmt.Println(args[0].Type().String())

		return nil
	})

	return cb
}
