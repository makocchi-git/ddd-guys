package util

import (
	"net/http"
)

// 独自の Error を定義
// TODO: i18n 化とか余裕があれば・・
type ErrorContentTypeMustBeSet struct{}

func (e *ErrorContentTypeMustBeSet) Error() string {
	return "Content-Type header must be set"
}

type ErrorContentTypeIsNotJSON struct{}

func (e *ErrorContentTypeIsNotJSON) Error() string {
	return "Content-Type header is not application/json"
}

// ちょっと func 名が長くなりすぎた感があります・・
func ValidateContentTypeApplicationJSON(h http.Header) error {
	contentType := h.Get("Content-TYpe")
	if len(contentType) == 0 {
		// Content-Type is not set
		return &ErrorContentTypeMustBeSet{}
	}

	// check Content-Type
	if contentType != "application/json" {
		return &ErrorContentTypeIsNotJSON{}
	}

	// ok :-)
	return nil
}
