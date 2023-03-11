package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertURLs(t *testing.T, expected string, actual []string) {
	if assert.Len(t, actual, 1) {
		assert.Equal(t, actual[0], expected)
	}
}

func TestHttp(t *testing.T) {
	url := "http://google.com"
	actual := ParseURLs(fmt.Sprintf("My http URL %s is here.", url))
	assertURLs(t, url, actual)
}

func TestHttps(t *testing.T) {
	url := "https://google.com"
	actual := ParseURLs(fmt.Sprintf("My https URL %s is here.", url))
	assertURLs(t, url, actual)
}

func TestFtp(t *testing.T) {
	url := "ftp://ftp.vim.org/pub/vim/README"
	actual := ParseURLs(fmt.Sprintf("My ftp URL %s is here.", url))
	assertURLs(t, url, actual)
}

func TestPathParams(t *testing.T) {
	url := "https://foo.com/bar/baz"
	actual := ParseURLs(fmt.Sprintf("My paths %s are here.", url))
	assertURLs(t, url, actual)
}

func TestQueryParams(t *testing.T) {
	url := "https://www.example.com?foo=bar&baz=qux"
	actual := ParseURLs(fmt.Sprintf("My query params: %s", url))
	assertURLs(t, url, actual)
}

func TestLocalhost(t *testing.T) {
	url := "http://localhost"
	actual := ParseURLs(fmt.Sprintf("This is my url %s.", url))
	assertURLs(t, url, actual)
}

func TestLocalhostWithPort(t *testing.T) {
	url := "http://localhost:3000"
	actual := ParseURLs(fmt.Sprintf("This is my url %s.", url))
	assertURLs(t, url, actual)
}

func TestURLWithPort(t *testing.T) {
	url := "http://google.com:3000/mark"
	actual := ParseURLs(fmt.Sprintf("This is my url %s.", url))
	assertURLs(t, url, actual)
}

func TestMarkdown(t *testing.T) {
	url := "https://google.com?foo=bar"
	actual := ParseURLs(fmt.Sprintf("This is [my link](%s).", url))
	assertURLs(t, url, actual)
}

func TestHyphens(t *testing.T) {
	url := "https://devblogs.microsoft.com/typescript/announcing-typescript-5-0-beta/"
	actual := ParseURLs(fmt.Sprintf("Description %s", url))
	assertURLs(t, url, actual)
}

func TestFileExtension(t *testing.T) {
	url := "http://google.com/index.js"
	actual := ParseURLs(fmt.Sprintf("Description %s", url))
	assertURLs(t, url, actual)
}

func TestSpaces(t *testing.T) {
	url := "http://google.com/foo+bar/spam%20eggs?foo=bar+baz&green=spam%20eggs"
	actual := ParseURLs(fmt.Sprintf("Description %s", url))
	assertURLs(t, url, actual)
}

func TestPeriodsInQueryParams(t *testing.T) {
	url := "https://google.com?foo=bar.baz"
	actual := ParseURLs(fmt.Sprintf("Description %s", url))
	assertURLs(t, url, actual)
}

func TestCommasInQueryParams(t *testing.T) {
	url := "https://google.com?foo=bar,baz"
	actual := ParseURLs(fmt.Sprintf("Description %s", url))
	assertURLs(t, url, actual)
}

func TestHash(t *testing.T) {
	url := "https://github.github.com/gfm/#html-blocks"
	actual := ParseURLs(fmt.Sprintf("This is %s my url", url))
	assertURLs(t, url, actual)
}

func TestMultiple(t *testing.T) {
	a := "http://google.com?foo=bar"
	b := "https://github.github.com/gfm/#html-blocks"
	actual := ParseURLs(fmt.Sprintf("This is one %s url, with %s another.", a, b))

	if assert.Len(t, actual, 2) {
		assert.Equal(t, actual[0], a)
		assert.Equal(t, actual[1], b)
	}
}

func TestUnpkgUrl(t *testing.T) {
	url := "https://unpkg.com/@mskelton/observer@1.0.1/dist/index.js"
	actual := ParseURLs(fmt.Sprintf("This is %s my url", url))
	assertURLs(t, url, actual)
}
