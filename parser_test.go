package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestBasicHTMLParsing(t *testing.T) {
	s := `<div><p><a href='https://some-link' @click='methodName' v:bind='data' data-test=bebebe>some text<br/>more text</a></p><br/></div>`
	r := strings.NewReader(s)
	z := html.NewTokenizer(r)

	el := ParseHTML(z)

	renderer := &DummyRenderer{}

	err := renderer.Render(el)
	assert.Nil(t, err)

	assert.Equal(t, el.Type, "div")
	assert.Equal(t, el.Children[0].Children[0].Type, "a")
	assert.Equal(t, el.Children[0].Children[0].Children[0].NodeValue, "some text")
	assert.Equal(t, el.Children[0].Children[0].Children[1].Type, "br")
	assert.Equal(t, el.Children[0].Children[0].Children[2].NodeValue, "more text")
}
