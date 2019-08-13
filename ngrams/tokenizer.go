package ngrams

import (
	"strings"
)

var (

	// skippableSpaces is a slice of whitespace-like runes which can be skipped.
	skippableSpaces = []rune{
		9,     // \t tab
		10,    // \n newline
		13,    // \r return
		32,    // \s space
		12288, // space (cjk)
	}

	// validPunctuation is a slice of punctuation which can be formed into tokens.
	// This list is not all-encompassing, but takes the most common punctuation
	// points found in literature. Quotemarks and the like are sanitized out (for now).
	validPunctuation = []rune{
		46,    // .
		44,    // ,
		63,    // ?
		33,    // !
		8253,  // ‽
		58,    // :
		59,    // ;
		38,    // &
		12290, // 。 (cjk)
		12289, // 、 (cjk)
		65281, // ！ (cjk)
		65311, // ？ (cjk)
		65306, // ： (cjk)
		65307, // ； (cjk)
	}

	// invalidChars is a slice of invalid characters that must be stripped.
	// These are virtually all parenthesis and quote marks.
	invalidChars = []rune{
		40,    // (
		41,    // )
		91,    // [
		93,    // ]
		123,   // {
		125,   // }
		34,    // "
		38,    // &
		8220,  // “
		8221,  // ”
		8216,  // ‘
		8217,  // ’
		171,   // «
		187,   // »
		8222,  // „
		12302, // 『
		12303, // 』
		12300, // 「
		12301, // 」
	}
)

// Tokenize splits a string into tokens. Each instance of standard punctuation
// is also considered to be token in order to preserve expected grammar.
func Tokenize(str string) []string {

	r := []rune(str)

	// If we were only taking words, we could use strings.FieldsFunc to extract
	// everything between the spaces. However, in order to preserve the expected
	// grammar of the sentences, we need to take the punctation too as discrete
	// tokens. The easiest way to do this is range through the string as runes,
	// extracting tokens as we go.

	start := 0           // start indicates the starting index of a new token.
	tokens := []string{} // tokens is the slice of tokens we're collecting.
	var j int            // j is the index, and it's up a scope so we can append the final word.
	var skipping bool    // skipping indicates that we're scanning over skippable runes.

	for ; j < len(r); j++ {

		// If the rune is skippable, note that we're not tracking skippable
		// runes, add any previous token to the slice, and continue.
		if isSkippableRune(r[j]) {
			skipping = true
			tokens = append(tokens, string(r[start:j]))
			continue
		}

		// If we're still going it means the current rune is not skippable. If
		// we were previously tracking skippables, that means we've now started a new token.
		if skipping {
			skipping = false
			start = j
		}

		// If the rune is a punctation mark, add it to the tokens and immediately
		// set a new start position. Punctuation is only valid if the next character
		// is whitespace, otherwise it will be treated as part of the previous token.
		// This allows us to preserve hyphenated words (eg. thirty-two, far-flung).
		// We must also capture any trailing punctuation.
		if isPunctuation(r[j]) && (j+1 < len(r) && r[j+1] == 32 || j == len(r)-1) {
			tokens = append(tokens, string(r[start:j]))
			start = j
		}
	}

	// If there are more characters left since the last split, take whatever is
	// left and create the last token.
	if j > start {
		tokens = append(tokens, string(r[start:]))
	}

	//for _, v := range tokens {
	//	log.Printf("`%s`\n", v)
	//}

	return tokens
}

// isPunctuation returns true if the rune is a variety of whitelisted punctuation.
func isPunctuation(c rune) bool {
	for j := 0; j < len(validPunctuation); j++ {
		if validPunctuation[j] == c {
			return true
		}
	}

	return false
}

// isSkippableRune returns true if the rune is a variety of whitespace or newline.
func isSkippableRune(c rune) bool {
	for j := 0; j < len(skippableSpaces); j++ {
		if skippableSpaces[j] == c {
			return true
		}
	}

	return false
}

// sanitize removes all unsupported characters from a string.
func sanitize(str string) string {
	return strings.Map(func(c rune) rune {
		for j := 0; j < len(invalidChars); j++ {
			if invalidChars[j] == c {
				return -1
			}
		}
		return c
	}, str)
}
