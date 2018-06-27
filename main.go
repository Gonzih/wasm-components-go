package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	cb := js.NewCallback(func(args []js.Value) {
		v := js.Global().Get("document").Call("getElementById", "myText").Get("value").String()
		fmt.Printf("v: %s\r\n", v)
	})

	js.Global().Get("document").Call("getElementById", "myText").Call("addEventListener", "input", cb)
	js.Global().Get("document").Call("getElementById", "runButton").Set("disabled", true)

	<-c
}
