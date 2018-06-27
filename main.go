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

type GenericComponent struct {
	props    interface{}
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
	err := c.template.Execute(buf, c.props)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func NewComponent(templateID, targetID string, propsFn func(*GenericComponent) error) (Component, error) {
	cmp := &GenericComponent{}
	markup := js.Global().Get("document").Call("getElementById", templateID).Get("innerHTML").String()

	tmpl, err := template.New("component").Parse(markup)

	if err != nil {
		return cmp, err
	}

	cmp.template = tmpl
	cmp.targetID = targetID

	err = propsFn(cmp)

	if err != nil {
		return cmp, err
	}

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

	cmp, err := NewComponent("helloTemplate", "root", func(cmp *GenericComponent) error {
		cmp.props = struct {
			Label string
		}{
			Label: "markup from props",
		}

		return nil
	})

	checkErr(err)
	checkErr(cmp.Render())

	<-c
}
