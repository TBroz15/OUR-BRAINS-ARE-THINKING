package helpers

import (
	"strings"
)

// Returns true If the string has the word "brain" and "think"
// regardless if it is plural, has prefixes, or has capitals.
func HasTheWords(commentStr string) bool {
	hasBrain := false
	hasThinking := false

	words := strings.SplitSeq(commentStr, " ")

	for word := range words {
		rootWord := RootWordGetter9000(word)

		if strings.Contains(rootWord, "brain") {
			hasBrain = true
		}

		if strings.Contains(rootWord, "think") {
			hasThinking = true
		}

		if hasBrain && hasThinking {
			break
		}
	}

	return hasBrain && hasThinking
}
