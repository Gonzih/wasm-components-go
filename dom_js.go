//+build js

package main

import "syscall/js"

type DomHelper struct {
}

func (d *DomHelper) SetInnerHTMLByID(id, content string) error {
	el := js.Global().Get("document").Call("getElementById", id)
	el.Set("innerHTML", content)

	return nil
}

func (d *DomHelper) GetInnerHTMLByID(id string) (string, error) {
	content := js.Global().Get("document").Call("getElementById", id).Get("innerHTML").String()

	return content, nil
}

func init() {
	domHelper = &DomHelper{}
}
