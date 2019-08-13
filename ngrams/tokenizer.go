package ngrams

import (
	"log"
)

// Tokenize splits a string into word tokens. Each instance of standard punctuation
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
	symbols := []rune{
		46,    // .
		44,    // ,
		63,    // ?
		33,    // !
		8253,  // ‽
		58,    // :
		59,    // ;
		12290, // 。 (cjk)
		12289, // 、 (cjk)
		65281, // ！ (cjk)
		65311, // ？ (cjk)
		65306, // ： (cjk)
		65307, // ； (cjk)
	}

	for j := 0; j < len(symbols); j++ {
		if symbols[j] == c {
			return true
		}
	}

	return false

}

// isSkippableRune returns true if the rune is a variety of whitespace or newline.
func isSkippableRune(c rune) bool {
	skippable := []rune{
		9,     // \t tab
		10,    // \n newline
		13,    // \r return
		32,    // \s space
		12288, // space (cjk)
	}

	for j := 0; j < len(skippable); j++ {
		if skippable[j] == c {
			return true
		}
	}

	return false
}
