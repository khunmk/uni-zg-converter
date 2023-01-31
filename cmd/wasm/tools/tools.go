package tools

import (
	"fmt"
	"syscall/js"
)

func RegisterCallBack() {
	js.Global().Set("add", js.FuncOf(add))

	//uni
	js.Global().Set("toUni", js.FuncOf(toUni))
	js.Global().Set("toZg", js.FuncOf(toZg))

	//txt
	js.Global().Set("convertTextFile", js.FuncOf(convertTextFile))
}

func add(this js.Value, i []js.Value) interface{} {
	param1 := js.ValueOf(i[0]).Int()
	param2 := js.ValueOf(i[1]).Int()

	fmt.Println("result is", param1+param2)

	return nil
}
