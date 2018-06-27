package main

import (
	"bytes"
	"fmt"
	"html/template"
	"syscall/js"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %s\r\n", err)
		panic(err)
	}
}

type Component interface {
	Render() error
	RenderToString() (string, error)
}

type Props struct {
}

type GenericComponent struct {
	props    Props
	template *template.Template
	targetID string
}

func (c *GenericComponent) Render() error {
	html, err := c.RenderToString()
	if err != nil {
		return err
	}

	js.Global().Get("document").Call("getElementById", c.targetID).Set("innerHTML", html)
	return nil
}

func (c *GenericComponent) RenderToString() (string, error) {
	buf := new(bytes.Buffer)
	err := c.template.Execute(buf, "")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func NewComponent(templateID, targetID string) (Component, error) {
	cmp := &GenericComponent{}
	markup := js.Global().Get("document").Call("getElementById", templateID).Get("innerHTML").String()

	tmpl, err := template.New("component").Parse(markup)

	if err != nil {
		return cmp, err
	}

	cmp.template = tmpl
	cmp.targetID = targetID

	return cmp, nil
}

func main() {
	c := make(chan struct{}, 0)
	cb := js.NewCallback(func(args []js.Value) {
		v := js.Global().Get("document").Call("getElementById", "myText").Get("value").String()
		fmt.Printf("v: %s\r\n", v)
	})

	js.Global().Get("document").Call("getElementById", "myText").Call("addEventListener", "input", cb)
	js.Global().Get("document").Call("getElementById", "runButton").Set("disabled", true)

	cmp, err := NewComponent("helloTemplate", "root")
	checkErr(err)
	checkErr(cmp.Render())

	<-c
}
