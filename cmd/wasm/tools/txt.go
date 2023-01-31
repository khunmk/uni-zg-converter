package tools

import (
	"fmt"
	"syscall/js"

	myanmartools "github.com/google/myanmar-tools/clients/go"
	"github.com/khunmk/uni-zg-converter/internal/utils"
)

func convertTextFile(this js.Value, i []js.Value) interface{} {
	file := js.ValueOf(i[0])

	//fmt.Println(file.Get("name"))

	return file.Call("arrayBuffer").Call("then", js.FuncOf(convertTxt))
}

func convertTxt(v js.Value, x []js.Value) any {
	data := js.Global().Get("Uint8Array").New(x[0])
	dst := make([]byte, data.Get("length").Int())

	js.CopyBytesToGo(dst, data)

	zgDetector := myanmartools.NewZawgyiDetector()
	score := zgDetector.GetZawgyiProbability(string(dst))
	fmt.Println(score)
	if score >= 0.97 {
		fmt.Println("zg to uni")
		outTxt, err := utils.Zg2uni(string(dst))
		if err != nil {
			return map[string]interface{}{
				"code": 500,
			}
		}
		return map[string]interface{}{
			"code": 0,
			"data": outTxt,
		}
	} else {
		fmt.Println("uni to zg")
		outTxt, err := utils.Uni2zg(string(dst))
		if err != nil {
			return map[string]interface{}{
				"code": 500,
			}
		}
		return map[string]interface{}{
			"code": 0,
			"data": outTxt,
		}
	}
}
