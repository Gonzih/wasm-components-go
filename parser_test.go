package main

import (
	"io"
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// TextToken
// A StartTagToken looks like <a>.
// StartTagToken
// An EndTagToken looks like </a>.
// EndTagToken
// A SelfClosingTagToken tag looks like <br/>.
// SelfClosingTagToken
// A CommentToken looks like <!--x-->.
// CommentToken
// A DoctypeToken looks like <!DOCTYPE x>
// DoctypeToken

func ConstructAnElement(token *html.Token, z *html.Tokenizer) {

}

// walk a tree
// contsructanelement should be called recursively

func TestBasicHTMLParsing(t *testing.T) {
	s := `<div><p><a href='https://some-link' @click='methodName' v:bind='data'>some text</a></p></div>`
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
			log.Print(t)

			log.Println(t.Data)
			log.Println(t.Attr)
		case tt == html.EndTagToken:
		case tt == html.SelfClosingTagToken:
		case tt == html.CommentToken:
		case tt == html.DoctypeToken:
		}
		break
	}
}
