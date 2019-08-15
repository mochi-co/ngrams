package ngrams

import (
	"errors"
	"strings"

	stores "github.com/mochi-co/ngrams/stores"
	tk "github.com/mochi-co/ngrams/tokenizers"
)

const (

	// defaultN is the default number of grams per record.
	defaultN int = 3
)

var (
	// ErrEmptyIndex indicates that the index has not yet learned any ngrams.
	ErrEmptyIndex = errors.New("index is empty")
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
func NewIndex(n int, o *Options) *Index {

	i := &Index{
		N:         n,
		Store:     stores.NewMemoryStore(),
		Tokenizer: tk.NewDefaultWordTokenizer(true),
	}

	// Ensure n is never 0.
	if i.N == 0 {
		i.N = defaultN
	}

	// Use custom options if available.
	if o != nil {
		if o.Store != nil {
			i.Store = o.Store
		}
		if o.Tokenizer != nil {
			i.Tokenizer = o.Tokenizer
		}
	}

	return i
}

// Close attempts to gracefully close and shutdown the index and any stores.
func (i *Index) Close() error {
	err := i.Store.Close()
	if err != nil {
		return err
	}

	return nil
}

// Parse parses a string into ngrams and adds them to the index.
func (i *Index) Parse(str string) (tokens []string, err error) {

	// Tokenize the string using whichever tokenizer was selected.
	tokens = i.Tokenizer.Tokenize(str)

	// Iterate through the tokens creating n-grams of n length.
	for j := 0; j < len(tokens); j++ {
		k, f := i.extractNgram(j, tokens)
		if k == "" {
			break
		}
		err = i.Store.Add(k, f)
		if err != nil {
			return
		}
	}

	return
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
	if j+n > len(tokens) {
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

	// Only add a future if there's one more token to support it.
	if j+n < len(tokens) {
		future = tokens[j+n]
	}

	return

}

// Result contains the result of a ngram lookup.
type Result struct {

	// Prefix is the last token of the key that was matched. It is added
	// to a key from the variations to make the next key, eg. Prefix+" "+VKey.
	Prefix string

	// Next contains the future variations of the ngram and the number of times
	// they were indexed (probabilty score).
	Next stores.Variations
}

// Seek returns potential ngrams from the store matching the seed string.
func (i *Index) Seek(key string) (ok bool, result *Result) {
	var v stores.Variations
	ok, v = i.Store.Get(key)
	if !ok {
		return
	}

	// Use the tokenizer to split the key, and return the last token as
	// part of the result to make next lookups more convenient.
	parts := i.Tokenizer.Tokenize(key)
	if len(parts) == 0 {
		return
	}

	result = &Result{
		Prefix: parts[len(parts)-1],
		Next:   v,
	}

	return
}

// Babble generates a random sequence of up to n ngrams. The future ngrams will be
// selected based on their probability. The n value total will include discrete
// punctuation depending on the tokenizer in use.
func (i *Index) Babble(start string, n int) (b string, err error) {

	// We need the start string as the first tokens in the selected output,
	// otherwise they'll be skipped. If the string is blank, this will start
	// with an empty slice and immediately seek any ngram.
	o := i.Tokenizer.Tokenize(start)

	// For however many tokens we want to use, we'll range through looking
	// for matching ngram keys.
	for j := 0; j < n; j++ {
		ok, r := i.Seek(start)
		if !ok { // If nothing was found for the ngram, pick a new ngram at random.
			k, _, err := i.Store.Any()
			if err != nil {
				return "", err
			}

			if k == "" {
				return "", ErrEmptyIndex
			}

			ok, r = i.Seek(k)
		}

		// Get the next ngram using a weighted random selection from the variations.
		next := r.Next.NextWeightedRand()
		start = r.Prefix + " " + next
		if next != "" {
			o = append(o, next)
		}
	}

	// Output the string using the token formatter for the in-use tokenizer.
	b = i.Tokenizer.Format(o)
	return
}
