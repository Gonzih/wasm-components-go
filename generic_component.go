package main

import (
	"bytes"
	"html/template"
	"syscall/js"
)

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

	tmpl, err := template.New(templateID).Parse(markup)

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
