package tokenizers

import (
	"strings"
)

// DefaultWord is the default tokenizer, designed to be used with
// bodies of text in english and other latin-based languages.
type DefaultWord struct {

	// skippable is a slice of whitespace-like runes which can be skipped.
	skippable []rune

	// punctuation is a slice of punctuation which can be formed into tokens.
	// This list is not all-encompassing, but takes the most common punctuation
	// points found in literature. Quotemarks and the like are sanitized out (for now).
	punctuation []rune

	// invalidChars is a slice of invalid characters that must be stripped.
	// These are virtually all parenthesis and quote marks.
	invalidChars []rune
}

// NewDefaultWordTokenizer returns a new default word tokenizer.
func NewDefaultWordTokenizer() *DefaultWord {
	return &DefaultWord{
		skippable: []rune{
			9,     // \t tab
			10,    // \n newline
			13,    // \r return
			32,    // \s space
			12288, // space (cjk)
		},
		punctuation: []rune{
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
		},
		invalidChars: []rune{
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
		},
	}
}

// Tokenize splits a string into tokens. Each instance of standard punctuation
// is also considered to be token in order to preserve expected grammar.
func (tk *DefaultWord) Tokenize(str string) []string {

	// Sanitize the input string.
	//str = strings.ToLower(str)
	str = tk.sanitize(str)

	r := []rune(str)
	tokens := make([]string, 0, len(str))

	// If we were only taking words, we could use strings.FieldsFunc to extract
	// everything between the spaces. However, in order to preserve the expected
	// grammar of the sentences, we need to also extract the punctation as discrete
	// tokens. The easiest way to do this is range through the string as runes,
	// extracting tokens as we go.

	var start int     // start indicates the starting index of a new token.
	var j int         // j is the index, and it's up a scope so we can append the final word.
	var skipping bool // skipping indicates that we're scanning over skippable runes.

	for ; j < len(r); j++ {

		// If the rune is skippable, note that we're not tracking skippable
		// runes, add any previous token to the slice, and continue.
		if tk.isSkippableRune(r[j]) {
			skipping = true
			tokens = append(tokens, string(r[start:j]))
			start = j // bring the start forward.
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
		if tk.isPunctuation(r[j]) && (j+1 < len(r) && r[j+1] == 32 || j == len(r)-1) {
			token := string(r[0])
			if j > 0 {
				token = string(r[start:j])
				tokens = append(tokens, token)
				start = j
			}
		}
	}

	// If there are more characters left since the last split, take whatever is
	// left and create the last token.
	if j > start {
		if token := strings.TrimSpace(string(r[start:])); token != "" {
			tokens = append(tokens, token)
		}
	}

	return tokens
}

// isPunctuation returns true if the rune is a variety of whitelisted punctuation.
func (tk *DefaultWord) isPunctuation(c rune) bool {
	for j := 0; j < len(tk.punctuation); j++ {
		if tk.punctuation[j] == c {
			return true
		}
	}

	return false
}

// isSkippableRune returns true if the rune is a variety of whitespace or newline.
func (tk *DefaultWord) isSkippableRune(c rune) bool {
	for j := 0; j < len(tk.skippable); j++ {
		if tk.skippable[j] == c {
			return true
		}
	}

	return false
}

// sanitize removes all unsupported characters from a string.
func (tk *DefaultWord) sanitize(str string) string {
	return strings.Map(func(c rune) rune {
		for j := 0; j < len(tk.invalidChars); j++ {
			if tk.invalidChars[j] == c {
				return -1
			}
		}
		return c
	}, strings.TrimSpace(str))
}

// Formatter joins a slice of tokens by the tokenizer rules.
func (tk *DefaultWord) Formatter(tokens []string) string {

	if len(tokens) == 0 {
		return ""
	}

	var o string
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "" { // Defensive coding. Continue on blank tokens.
			continue
		}

		if i > 0 && !tk.isPunctuation([]rune(tokens[i])[0]) {
			o += " "
		}
		o += tokens[i]
	}

	return o

}
