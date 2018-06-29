package main

import (
	"bytes"
	"html/template"
	"syscall/js"

	"github.com/google/uuid"
)

type GenericComponent struct {
	props       interface{}
	propsFn     func(*GenericComponent) error
	template    *template.Template
	targetID    string
	componentID string
}

func (c *GenericComponent) Render() error {
	globalObserver.SetContext(c.componentID)

	html, err := c.RenderToString()
	if err != nil {
		return err
	}

	js.Global().Get("document").Call("getElementById", c.targetID).Set("innerHTML", html)
	return nil
}

func (c *GenericComponent) RenderToString() (string, error) {
	err := c.propsFn(c)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = c.template.Execute(buf, c.props)
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
	cmp.propsFn = propsFn
	cmp.componentID = uuid.New().String()

	return cmp, nil
}
