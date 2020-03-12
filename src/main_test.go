package main

import (
	"encoding/xml"
	"io"
	"strings"
	"testing"
)

// Poor man's HTML validator, courtesy of https://stackoverflow.com/a/52410528
func validateHtml(r io.Reader) bool {
	d := xml.NewDecoder(r)

	// Configure the decoder for HTML; leave off strict and autoclose for XHTML
	d.Strict = true
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		_, err := d.Token()
		switch err {
		case io.EOF:
			return true // We're done, it's valid!
		case nil:
		default:
			return false // Oops, something wasn't right
		}
	}
}

// Very contrived test for a very basic application.
func TestHtml(t *testing.T) {
	path := "/"
	ok := validateHtml(strings.NewReader(helloHtml(path)))
	if !ok {
		t.Fatal("HTML was not valid")
	}
}

func TestGreeting(t *testing.T) {
	path := "/"
	text := helloHtml(path)
	if strings.Contains(text, "Salutations") {
		t.Fatal("Wrong greeting in text")
	}
}

func TestJavaScriptInjection(t *testing.T) {
	path := "/<script>alert('hello')</script>"
	text := helloHtml(path)
	if strings.Contains(text, "<script>") {
		t.Fatal("JavaScript injection attack")
	}
}

func TestConvertFoo(t *testing.T) {
	path := "/foo/bar"
	text := convertFoo(path)
	if strings.Contains(text, "foo") {
		t.Fatal("'foo' in path")
	}
}
