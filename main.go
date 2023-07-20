package main

import (
	"strconv"
	"syscall/js"
)

var document js.Value

func init() {
	document = js.Global().Get("document")
}

func getValue(elem string) string {
	return document.Call("getElementById", elem).Get("value").String()
}

func sum1(this js.Value, args []js.Value) any {
	first := args[0].Int()
	second := args[1].Int()
	return map[string]any{"message": first + second}
}

func sum2(this js.Value, args []js.Value) any {
	first, _ := strconv.ParseInt(getValue("first"), 10, 0)
	second, _ := strconv.ParseInt(getValue("second"), 10, 0)

	return map[string]any{"message": first + second}
}

func main() {
	js.Global().Set("goSum1", js.FuncOf(sum1))
	js.Global().Set("goSum2", js.FuncOf(sum2))
	select {} // keep running
}
