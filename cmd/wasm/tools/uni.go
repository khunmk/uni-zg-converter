package tools

import (
	"syscall/js"

	"github.com/khunmk/uni-zg-converter/internal/utils"
)

func toUni(this js.Value, i []js.Value) interface{} {
	str := js.ValueOf(i[0]).String()

	outputTxt, err := utils.Zg2uni(str)
	if err != nil {
		return map[string]interface{}{
			"code": 500,
		}
	}

	//fmt.Println(outputTxt)

	return map[string]interface{}{
		"code": 0,
		"data": outputTxt,
	}
}

func toZg(this js.Value, i []js.Value) interface{} {
	str := js.ValueOf(i[0]).String()

	outputTxt, err := utils.Zg2uni(str)
	if err != nil {
		return map[string]interface{}{
			"code": 500,
		}
	}

	return map[string]interface{}{
		"code": 0,
		"data": outputTxt,
	}
}
