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

	// fileInputEle := document.Call("getElementById", "file_input")
	// if !fileInputEle.Truthy() {
	// 	log.Fatalln("unable to get element")
	// }

	uniInputEle.Call("addEventListener", "input", toZg(zgInputEle))
	zgInputEle.Call("addEventListener", "input", toUni(uniInputEle))

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

// func onFileInput() js.Func {
// 	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
// 		length := this.Get("files").Length()

// 		for i := 0; i < length; i++ {
// 			fileEle := this.Get("files").Call("item", i)

// 			fileName := fileEle.Get("name").String()

// 			fmt.Println(fileName)

// 			// file, err := ioutil.TempFile("dir", "prefix")
// 			// if err != nil {
// 			// 	log.Fatal(err)
// 			// }
// 			// defer os.Remove(file.Name())

// 			// fileEle.Call("arrayBuffer").Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
// 			// 	data := js.Global().Get("Uint8Array").New(args[0])
// 			// 	dst := make([]byte, data.Get("length").Int())
// 			// 	js.CopyBytesToGo(dst, data)

// 			// 	r := bytes.NewReader(dst)

// 			// 	buf := make([]byte, 4)
// 			// 	for {
// 			// 		n, err := r.Read(buf)
// 			// 		if err == io.EOF {
// 			// 			break
// 			// 		}

// 			// 		fmt.Println(string(buf[:n]))
// 			// 	}

// 			// 	// f, _ := os.Create("temp.txt")
// 			// 	// defer f.Close()
// 			// 	// f.Write()

// 			// 	return nil
// 			// }))
// 		}

// 		return nil
// 	})

// 	return cb
// }
