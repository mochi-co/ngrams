package ngrams

import (
	"log"
	"strings"

	stores "github.com/mochi-co/trigrams-test/stores"
	tk "github.com/mochi-co/trigrams-test/tokenizers"
)

const (

	// defaultN is the default number of grams per record.
	defaultN int = 3
)

// Options contains parameters for the ngram indexer.
type Options struct {

	// Store contains the ngrams which have been indexed.
	Store stores.Store

	// Tokenizer is the tokenizer to use to split strings into tokens.
	Tokenizer tk.Tokenizer
}

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

	// Tokenize the string using whichever tokenizer was selected.
	tokens := i.Tokenizer.Tokenize(str)

	// Iterate through the tokens creating n-grams of n length.
	for j := 0; j < len(tokens); j++ {
		k, f := i.extractNgram(j, tokens)
		if k == "" {
			break
		}
		i.Store.Add(k, f)
		log.Println(k, f)
	}

	i.Store.(*stores.MemoryStore).Print()

	return nil
}

// extractNgram extracts the maximum possible length ngram from a slice of
// tokens, starting at index and continuing until either n or len(tokens)
// has been met.
func (i *Index) extractNgram(j int, tokens []string) (key, future string) {

	// Because we're using the n value for index lookups,
	// it needs to start at 0 not 1.
	n := i.N - 1

	// An n-gram must have exactly as many tokens as (n), otherwise it would be
	// called a sort-of-n-but-actually-sometimes-not-gram.
	if j+n >= len(tokens) {
		return
	}

	// Using whitespace as token joiners. In most cases whitespace is stripped
	// in the tokenizer, so it will be unique. If you were using this in a
	// different scenario (say, biological analysis) and you observed
	// whitespace as a valid character, you could run into problems if there
	// were competing entries: "a", "b", and "a b".
	// That's a bit out of scope, but worth noting nonetheless.
	if n > 0 {
		key = strings.Join(tokens[j:j+n], " ")
	} else { // Handle monograms in case anyone wants to do that (n=1, -1, n==0).
		key = tokens[j]
	}

	future = tokens[j+n]

	return

}

// Seek returns potential ngrams from the store matching the seed string.
func (i *Index) Seek(seed string) Result {

	return Result{}
}

// Result contains the result of a ngram lookup.
type Result struct {
	Next string
}
