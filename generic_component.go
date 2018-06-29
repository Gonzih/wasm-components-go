package main

import (
	"bytes"
	"html/template"
	"log"
	"syscall/js"

	"github.com/albrow/vdom"
	"github.com/google/uuid"
)

type GenericComponent struct {
	props       interface{}
	propsFn     func(*GenericComponent) error
	template    *template.Template
	tree        *vdom.Tree
	dirty       bool
	targetID    string
	componentID string
}

func (c *GenericComponent) Notify() {
	c.dirty = true
}

func (c *GenericComponent) RenderToString() (string, error) {
	return string(c.tree.HTML()), nil
}

func (c *GenericComponent) Render() error {
	if c.dirty {
		globalObserver.SetContext(c.Notify)

		log.Println("Regenerating dom tree")
		err := c.propsFn(c)

		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)

		err = c.template.Execute(buf, c.props)
		if err != nil {
			return err
		}

		newTree, err := vdom.Parse(buf.Bytes())
		if err != nil {
			return err
		}

		if c.tree != nil && len(c.targetID) > 0 {
			// Calculate the diff between this render and the last render
			// patches, err := vdom.Diff(c.tree, newTree)
			// if err != nil {
			// 	return err
			// }

			// Effeciently apply changes to the actual DOM
			// root := js.Global().Get("document").Call("getElementById", c.targetID)
			// if err := patches.Patch(root); err != nil {
			// 	return err
			// }
		}

		c.tree = newTree
		c.dirty = false
	}

	root := js.Global().Get("document").Call("getElementById", c.targetID)
	html, err := c.RenderToString()
	if err != nil {
		return err
	}
	root.Set("innerHTML", html)

	return nil
}

func NewComponent(templateID, targetID string, propsFn func(*GenericComponent) error) (Component, error) {
	c := &GenericComponent{}
	markup := js.Global().Get("document").Call("getElementById", templateID).Get("innerHTML").String()

	tmpl, err := template.New(templateID).Parse(markup)

	if err != nil {
		return c, err
	}

	c.template = tmpl
	c.targetID = targetID
	c.propsFn = propsFn
	c.componentID = uuid.New().String()
	c.dirty = true

	return c, nil
}
