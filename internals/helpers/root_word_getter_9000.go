package helpers

import (
	"strings"

	"github.com/kljensen/snowball"
)

// aka. a lemmatizer? idk...
// it returns the root word of the english language
func RootWordGetter9000(word string) string {
	word = toEnglishOnlyLetters(word)

	stemmed, err := snowball.Stem(word, "english", true)
	if err != nil {
		return ""
	}

	return stemmed
}

// ENGLISH MF, DO YOU SPEAK IT???
func toEnglishOnlyLetters(str string) string {
	b := strings.Builder{}

	for _, rune := range str {
		if (rune >= 'A' && rune <= 'Z') || (rune >= 'a' && rune <= 'z') {
			b.WriteRune(rune)
		}
	}

	return b.String()
}
