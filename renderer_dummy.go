package main

import (
	"fmt"
	"log"
	"strings"
)

type DummyRenderer struct {
}

func (_ *DummyRenderer) Render(el *El) error {
	return printElement(el, 0)
}

func printElement(el *El, level int) error {
	prefix := strings.Repeat("  ", level)
	attrs := ""
	for _, attr := range el.Attr {
		attrs += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Val)
	}

	switch el.Type {
	case "text":
		log.Printf("%s%s", prefix, el.NodeValue)
	case "br":
		log.Printf("%s<%s%s/>", prefix, el.Type, attrs)
	default:
		log.Printf("%s<%s%s>", prefix, el.Type, attrs)
		for _, child := range el.Children {
			err := printElement(child, level+1)
			if err != nil {
				return err
			}
		}
		log.Printf("%s</%s>", prefix, el.Type)
	}

	return nil
}
