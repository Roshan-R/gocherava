package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

func gethtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("Could not get html")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body), nil
}

func scrape(url string, selector string) (string, error) {
	site_html, err := gethtml(url)

	if err != nil {
		return "", err
	}

	sel, err := css.Parse(selector)
	if err != nil {
		return "", errors.New("Could not create selector")
	}
	node, err := html.Parse(strings.NewReader(site_html))
	if err != nil {
		return "", errors.New("Could parse html")
	}

	var buf bytes.Buffer
	w := io.Writer(&buf)

	for _, ele := range sel.Select(node) {
		html.Render(w, ele.FirstChild)
	}

	return buf.String(), nil

}
