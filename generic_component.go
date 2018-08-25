package main

import (
	"bytes"
	"html/template"
	"log"

	"github.com/google/uuid"
)

type GenericComponent struct {
	props       interface{}
	propsFn     func(*GenericComponent) error
	template    *template.Template
	dirty       bool
	targetID    string
	componentID string
}

func (c *GenericComponent) Notify() {
	c.dirty = true
}

func (c *GenericComponent) RenderToString() (string, error) {
	buf := new(bytes.Buffer)

	err := c.propsFn(c)

	if err != nil {
		return buf.String(), err
	}

	err = c.template.Execute(buf, c.props)
	if err != nil {
		return buf.String(), err
	}

	return buf.String(), nil
}

func (c *GenericComponent) Render() error {
	if c.dirty {

		globalObserver.SetContext(c.Notify)

		log.Println("Regenerating dom tree")

		html, err := c.RenderToString()
		if err != nil {
			return err
		}
		domHelper.SetInnerHTMLByID(c.targetID, html)
		c.dirty = false
	}

	return nil
}

func NewComponent(templateID, targetID string, propsFn func(*GenericComponent) error) (Component, error) {
	c := &GenericComponent{}
	markup, _ := domHelper.GetInnerHTMLByID(templateID)

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
