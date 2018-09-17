package main

import (
	"golang.org/x/net/html"
)

type El struct {
	Type        string
	Attr        []html.Attribute
	NodeValue   string
	Children    []*El
	SelfClosing bool
}
