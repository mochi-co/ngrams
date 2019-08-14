package tokenizers

// Tokenizer is a string tokenizer which splits a string into discrete tokens.
type Tokenizer interface {

	// Tokenize tokenizes a string.
	Tokenize(string) []string

	// Formatter joins a slice of tokens by the tokenizer rules.
	Formatter([]string) string
}
