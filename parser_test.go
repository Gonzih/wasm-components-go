package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

type El struct {
	Tag         string
	Attr        []html.Attribute
	Content     string
	Children    []*El
	SelfClosing bool
}

func (el *El) Print() {
	el.printElement(0)
}

func (el *El) printElement(level int) {
	prefix := strings.Repeat("  ", level)
	attrs := ""
	for _, attr := range el.Attr {
		attrs += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Val)
	}

	if el.SelfClosing {
		log.Printf("%s<%s%s/>", prefix, el.Tag, attrs)
	} else {
		log.Printf("%s<%s%s>", prefix, el.Tag, attrs)
		if len(el.Content) > 0 {
			log.Printf("%s  %s", prefix, el.Content)
		}
		for _, child := range el.Children {
			child.printElement(level + 1)
		}
		log.Printf("%s</%s>", prefix, el.Tag)
	}
}

func ConstructAnElement(tt html.TokenType, z *html.Tokenizer) *El {
	token := z.Token()

	parent := &El{}

	parent.Tag = token.Data
	parent.Attr = token.Attr

	if tt != html.SelfClosingTagToken {
		for {
			tt := z.Next()
			switch {
			case tt == html.ErrorToken:
				err := z.Err()
				if err == io.EOF {
					return parent
				}
				log.Printf("Error: %s", err)
			case tt == html.StartTagToken:
				child := ConstructAnElement(tt, z)
				parent.Children = append(parent.Children, child)
			case tt == html.TextToken:
				t := z.Token()
				parent.Content = t.Data
			case tt == html.EndTagToken:
				break
			case tt == html.SelfClosingTagToken:
				child := ConstructAnElement(tt, z)
				parent.Children = append(parent.Children, child)
			case tt == html.CommentToken:
				break
			case tt == html.DoctypeToken:
				break
			}
		}
	} else {
		parent.SelfClosing = true
	}

	return parent
}

func ParseHTML(z *html.Tokenizer) *El {
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			err := z.Err()
			if err == io.EOF {
				return nil
			}
			log.Fatal(err)
		case tt == html.StartTagToken:
			return ConstructAnElement(tt, z)
		default:
			log.Fatalf("Wrong token type %v", tt)
		}
		break
	}

	return nil
}

// walk a tree
// contsructanelement should be called recursively

func TestBasicHTMLParsing(t *testing.T) {
	s := `<div><p><a href='https://some-link' @click='methodName' v:bind='data' data-test=bebebe>some text</a></p><br/></div>`
	r := strings.NewReader(s)
	z := html.NewTokenizer(r)

	el := ParseHTML(z)

	el.Print()
}
