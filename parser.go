package main

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

func ConstructAnElement(tt html.TokenType, z *html.Tokenizer) *El {
	token := z.Token()

	parent := &El{}

	parent.Type = token.Data
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
				child := &El{Type: "text", NodeValue: t.Data}
				parent.Children = append(parent.Children, child)
			case tt == html.EndTagToken:
				return parent
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
