//+build !js

package main

import "log"

type DomHelper struct {
}

func (d *DomHelper) SetInnerHTMLByID(id, content string) error {
	log.Println("Trying to set content of #%s to \"%s\"", id, content)

	return nil
}

func (d *DomHelper) GetInnerHTMLByID(id string) (string, error) {
	return "<p>dummy template</p>", nil
}

func init() {
	domHelper = &DomHelper{}
}
