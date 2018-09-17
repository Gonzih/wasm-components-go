package main

import (
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

func (el *El) Print(level int) {
	prefix := strings.Repeat("\t", level)
	if el.SelfClosing {
		log.Printf("%s<%s %v/>", prefix, el.Tag, el.Attr)
	} else {
		log.Printf("%s<%s %v>", prefix, el.Tag, el.Attr)
		if len(el.Content) > 0 {
			log.Printf("%s\t%s", prefix, el.Content)
		}
		for _, child := range el.Children {
			child.Print(level + 1)
		}
		log.Printf("%s</%s>", prefix, el.Tag)
	}
}

func ConstructAnElement(token *html.Token, tt html.TokenType, z *html.Tokenizer) *El {
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
				t := z.Token()
				child := ConstructAnElement(&t, tt, z)
				parent.Children = append(parent.Children, child)
			case tt == html.TextToken:
				t := z.Token()
				parent.Content = t.Data
			case tt == html.EndTagToken:
				break
			case tt == html.SelfClosingTagToken:
				t := z.Token()
				child := ConstructAnElement(&t, tt, z)
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

// walk a tree
// contsructanelement should be called recursively

func TestBasicHTMLParsing(t *testing.T) {
	s := `<div><p><a href='https://some-link' @click='methodName' v:bind='data' data-test=bebebe>some text</a></p><br/></div>`
	r := strings.NewReader(s)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			err := z.Err()
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		case tt == html.StartTagToken:
			t := z.Token()
			el := ConstructAnElement(&t, tt, z)
			el.Print(0)
		}
		break
	}
}
