package main

type DomHelperer interface {
	SetInnerHTMLByID(string, string) error
	GetInnerHTMLByID(string) (string, error)
}

var domHelper DomHelperer
