package main

/*
// require Go 1.17

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// Translator calls google translate api.
// Dev account: https://cloud.google.com/docs/authentication/getting-started
// Limitation: https://cloud.google.com/translate/quotas
type Translator struct {
	targetLanguage language.Tag
	googleClient   *translate.Client
}

func NewTranslator(targetLanguage string) (*Translator, error) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS",
		"/home/tungdt/go/src/github.com/daominah/google_application_credentials.json")
	targetLangTag, err := language.Parse(targetLanguage)
	if err != nil {
		return nil, fmt.Errorf("error when Parse %v: %v", targetLanguage, err)
	}
	googleClient, err := translate.NewClient(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error when NewClient: %v", err)
	}
	return &Translator{targetLanguage: targetLangTag, googleClient: googleClient}, nil
}

func (t Translator) TranslateMany(originTexts []string) ([]string, error) {
	targets, err := t.googleClient.Translate(
		context.Background(), originTexts, t.targetLanguage, nil)
	if err != nil {
		return nil, fmt.Errorf("error when Translate: %v", err)
	}
	ret := make([]string, len(targets))
	for i, targetText := range targets {
		ret[i] = targetText.Text
	}
	return ret, nil
}



func main() {
	//t, err := NewTranslator("vi")
	t, err := NewTranslator("en")
	if err != nil {
		log.Fatal(err)
	}
	//translateds, err := t.TranslateMany([]string{"fuck you"})
	translateds, err := t.TranslateMany([]string{"con cặc"})
	if err != nil || len(translateds) < 1 {
		log.Fatal(err)
	}
	log.Println("translated:", translateds[0])
}

*/
