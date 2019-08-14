package tokenizers

import (
	"log"
	"strings"
)

// DefaultWord is the default tokenizer, designed to be used with
// bodies of text in english and other latin-based languages.
type DefaultWord struct {

	// skippable is a slice of whitespace-like runes which can be skipped.
	skippable []rune

	// stoppers is a slice of punctuation that indicate the end of a sentence.
	stoppers []rune

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
		stoppers: []rune{
			46,    // .
			63,    // ?
			33,    // !
			8253,  // ‽
			12290, // 。 (cjk)
			65311, // ？ (cjk)
			65281, // ！ (cjk)
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
		if runeInSlice(r[j], tk.skippable) {
			log.Println("skipping", r[start:j], string(r[start:j]))
			if !skipping {
				tokens = append(tokens, string(r[start:j]))
				skipping = true
			}

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
		if runeInSlice(r[j], tk.punctuation) && (j+1 < len(r) && r[j+1] == 32 || j == len(r)-1) {
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

// sanitize removes all unsupported characters from a string.
func (tk *DefaultWord) sanitize(str string) string {
	return strings.Map(func(c rune) rune {
		if runeInSlice(c, tk.invalidChars) {
			return -1
		}
		return c
	}, strings.TrimSpace(str))
}

// Format joins a slice of tokens by the tokenizer rules.
// If we want to output the ngrams again, they're not always going to
// go back together in the same manner they were consumed, especially if
// the tokenizers are wildly different. A formatter allows us to ensure the
// resulting text is appropriate.
func (tk *DefaultWord) Format(tokens []string) string {
	if len(tokens) == 0 {
		return ""
	}

	var o string
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "" { // Defensive coding. Continue on blank tokens.
			continue
		}
		log.Printf("%d `%s`", i, tokens[i])

		// If the token is punctuation, just append it and move on.
		if runeInSlice([]rune(tokens[i])[0], tk.punctuation) {
			o += tokens[i]
			continue
		}

		// If the word is not the first, add a space beforehand.
		if i > 0 {
			o += " "
		}

		// If the word follows a stopper, or is at the start of the sentence,
		// then it should be capitalized.
		if i == 0 || runeInSlice([]rune(tokens[i-1])[0], tk.stoppers) {
			o += strings.ToUpper(string(tokens[i][0])) + string(tokens[i][1:])
		} else {
			o += tokens[i]
		}

		// If it's the last word, and it's not punctuation, add a . to the end.
		if i == len(tokens)-1 && !runeInSlice([]rune(tokens[i])[0], tk.punctuation) {
			o += "."
		}
	}

	return o

}
