package client

import (
	"log"

	"github.com/st3v/translator/google"
)

const (
	GoogleApiKey = ""
)

func Translate(text string) string {
	translator := google.NewTranslator(GoogleApiKey)

	translation, err := translator.Translate(text, "en", "pt")
	if err != nil {
		log.Panicf("Error during translation: %s", err.Error())
	}

	return translation
}
