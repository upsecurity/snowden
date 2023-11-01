package client

import (
	"log"
	"os"

	"github.com/st3v/translator/google"
)

func Translate(text string) string {
	translator := google.NewTranslator(os.Getenv("GOOGLE_API_KEY"))

	translation, err := translator.Translate(text, "en", "pt")
	if err != nil {
		log.Panicf("Error during translation: %s", err.Error())
	}

	return translation
}
