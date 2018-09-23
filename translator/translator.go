package translator

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
)

type TranslationResult struct {
	word string
	desc string
}

type ITranslator interface {
	Translate(source, dest, text string) (TranslationResult, error)
}

type GoogleTranslator struct {
	urlFormat string
}

func DefaultTranslator() ITranslator {
	return &GoogleTranslator{
		urlFormat: "https://translate.google.com/#%s/%s/%s",
	}
}

func (g *GoogleTranslator) Translate(source, dest, text string) (result TranslationResult, err error) {
	escapedText := url.PathEscape(text)
	url := fmt.Sprintf(g.urlFormat, source, dest, escapedText)

	res, err := http.Get(url)
	if err != nil {
		return
	}

	node, err := html.Parse(res.Body)
	fmt.Println(node.Data)
	return
}
