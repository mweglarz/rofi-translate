package translator

import (
	"fmt"
	"net/http"
	"net/url"
)

type ITranslator interface {
	Translate(source, dest, text string) ([]TranslationResult, error)
}

type TranslationResult struct {
	word string
	desc string
}

func DefaultTranslator() ITranslator {
	// return &GoogleTranslator{
	// urlFormat: "https://translate.google.com/#%s/%s/%s",
	// }

	return &GoogleTranslator{
		urlFormat: "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s",
	}
}

// Google translator

type GoogleTranslator struct {
	urlFormat string
}

func (g *GoogleTranslator) Translate(source, dest, text string) (result []TranslationResult, err error) {
	escapedText := url.PathEscape(text)
	url := fmt.Sprintf(g.urlFormat, source, dest, escapedText)

	res, err := http.Get(url)
	if err != nil {
		return
	}
	handler := NewResponseHandler()
	return handler.HandleResponse(res.Body)
}
