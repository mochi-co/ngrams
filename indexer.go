package ngrams

import (
	stores "github.com/mochi-co/trigrams-test/stores"
	tk "github.com/mochi-co/trigrams-test/tokenizers"
)

const (

	// defaultN is the default number of grams per record.
	defaultN int = 3
)

// Index indexes ngrams and provides meachnisms for ngram retrieval and
// generation.
type Index struct {

	// N is the number of grams to store per key.
	N int

	// Store contains the ngrams which have been indexed.
	Store stores.Store

	// Tokenizer is the tokenizer to use to split strings into tokens.
	Tokenizer tk.Tokenizer
}

// Options contains parameters for the ngram indexer.
type Options struct {

	// Store contains the ngrams which have been indexed.
	Store stores.Store

	// Tokenizer is the tokenizer to use to split strings into tokens.
	Tokenizer tk.Tokenizer
}

// NewIndex returns a pointer to an Ngrams Index. It can be initialized
// with a custom store and tokenizer, otherwise the default in-memory store
// and latin word tokenizer will be used.
func NewIndex(n int, o Options) *Index {
	i := &Index{
		N:         n,
		Store:     o.Store,
		Tokenizer: o.Tokenizer,
	}

	if i.N == 0 {
		i.N = defaultN
	}

	if i.Store == nil {
		i.Store = stores.NewMemoryStore()
	}

	if i.Tokenizer == nil {
		i.Tokenizer = tk.NewDefaultWordTokenizer()
	}

	return i
}

// Parse parses a string into ngrams and adds them to the index.
func (i *Index) Parse(str string) error {

}

// Next returns the next potential ngrams from the store.
func (i *Index) Next(key string) stores.Potential {

}
