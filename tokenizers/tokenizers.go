package tokenizers

// Tokenizer is a string tokenizer which splits a string into discrete tokens.
type Tokenizer interface {

	// Tokenize tokenizes a string.
	Tokenize(string) []string

	// Scanner tokenizes a byte array and can be used by bufio.Scanner.
	Scanner(data []byte, atEOF bool) (advance int, token []byte, err error)

	// Format joins a slice of tokens by the tokenizer rules.
	Format([]string) string
}

// runeInSlice returns true if the rune was found in the slice of runes.
func runeInSlice(c rune, r []rune) bool {
	for i := 0; i < len(r); i++ {
		if r[i] == c {
			return true
		}
	}

	return false
}
