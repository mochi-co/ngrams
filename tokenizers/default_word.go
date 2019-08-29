package tokenizers

import (
	"bufio"
	"bytes"
	//"log"
	"strings"
	"unicode/utf8"
)

// DefaultWord is the default tokenizer, designed to be used with
// bodies of text in english and other latin-based languages.
type DefaultWord struct {

	// skippable is a slice of whitespace-like runes which can be skipped.
	skippable []rune

	// unskippable is a slice of characters which are not skippable, but should
	// indicate the end of a token. This is used when line breaks are preserved
	// in order to indicate that punctuation should be a stopper.
	unskippable []rune

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
func NewDefaultWordTokenizer(stripLinebreaks bool) *DefaultWord {
	d := &DefaultWord{
		skippable: []rune{
			9,     // \t tab
			32,    // \s space
			12288, // space (cjk)
		},

		// Invalid characters which should be ignored.
		invalidChars: []rune{
			40,   // (
			41,   // )
			91,   // [
			93,   // ]
			123,  // {
			125,  // }
			34,   // "
			38,   // &
			8220, // “
			8221, // ”
			8216, // ‘
			//8217,  // ’
			171,   // «
			187,   // »
			8222,  // „
			12302, // 『
			12303, // 』
			12300, // 「
			12301, // 」
			42,    // *
		},

		// stoppers are characters which indicate the end of a sentence (for formatting).
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
	}

	if stripLinebreaks {
		d.skippable = append(d.skippable, 10) // \n newline
		d.skippable = append(d.skippable, 13) // \r return
	} else {

		// If we keep line breaks, we have to treat them as punctuation and
		// ensure they are not skippable in the punctation splitter.
		d.punctuation = append(d.punctuation, 10) // \n newline
		d.punctuation = append(d.punctuation, 13) // \r return
		d.unskippable = append(d.unskippable, 10) // \n newline
		d.unskippable = append(d.unskippable, 13) // \r return
	}

	return d
}

// Tokenize splits a string into tokens. Each instance of standard punctuation
// is also considered to be token in order to preserve expected grammar.
func (tk *DefaultWord) Tokenize(str string) []string {

	// Parse the string into the reader so it can be scanned using the tokenizer's
	// own scanner.
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(tk.Scanner)

	tokens := []string{}
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	return tokens

}

// Scanner is the core tokenizer which splits a slice of bytes into tokens. It can be
// used with bufio.scanner to tokenizer a stream of data.
func (tk *DefaultWord) Scanner(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Skip leading skippable characters.
	var start int
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !runeInSlice(r, tk.skippable) {
			break
		}
	}

	// Found a non-skippable, so continue searching for token end.
	// i is the byte index, width is rune width in bytes, start is byte start pos.
	for width, i := 0, start; i < len(data); i += width {

		// Get the first rune in the advanced slice...
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		// If the rune is a space or skippable character, return everything so far.
		if runeInSlice(r, tk.skippable) {
			return i + width, tk.sanitize(data[start:i]), nil
		}

		// If the rune is punctuation and there's other runes since the start,
		// return everything up to now, but make sure to start on this rune next time.
		if runeInSlice(r, tk.punctuation) {
			if i+width-utf8.RuneLen(r) > 0 {

				// If the next rune is skippable, this is trailing punctuation
				// and can be split.
				nr, nw := utf8.DecodeRune(data[i+width:]) // Get next rune
				if (runeInSlice(nr, tk.skippable) || runeInSlice(nr, tk.invalidChars) || runeInSlice(nr, tk.unskippable)) || nw == 0 {
					return i, tk.sanitize(data[start:i]), nil
				}
			}
		}
	}

	// If at EOF, return whatever is left.
	if atEOF && len(data) > start {
		return len(data), tk.sanitize(data[start:]), nil
	}

	// Return for next data.
	return start, nil, nil

}

// sanitize will remove any invalid characters from a byte string.
func (tk *DefaultWord) sanitize(data []byte) []byte {
	rs := bytes.Runes(data)
	buf := make([]byte, 0, len(data))
	for _, r := range rs {
		if !runeInSlice(r, tk.invalidChars) {
			buf = append(buf, []byte(string(r))...)
		}
	}
	return buf
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
		if i == 0 || (len(tokens[i]) > 1 && runeInSlice([]rune(tokens[i-1])[0], tk.stoppers)) {
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
