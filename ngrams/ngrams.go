package ngrams

import (
	"log"
	//	"strings"
	//	"unicode"
)

const (

	// defaultN is the default quantity of n-grams to index per key.
	defaultN int = 3
)

// Store is an data storage mechanism for the ngrams. All the methods
// here are exported in case we wanted to add a store from a different
// package.
type Store interface {

	// Add adds a new ngram and key-variation pair to the index.
	Add(string, []string) error

	// Get returns the future ngram variations for a key.
	Get(string) (bool, Futures)

	// Delete removes an ngram key and all variations from the index.
	Delete(string) error
}

// Grams is a map of Gram keyed on gram-key (eg. "to be"). This is
// primarily used by the in-memory store, but can also be used to
// structure data for other storage engines.
type Grams map[string]Futures

// Futures contains the future-sequenced ngrams and the number of times
// they were indexed, keyed on gram-variation (eg. {"be or":3})
type Futures map[string]int64

// Index parses and stores ngrams.
type Index struct {

	// N is the number of grams to store per key.
	N int

	// Store contains the ngrams which have been indexed, and provides
	// methods for setting/getting.
	Store Store
}

// New returns a pointer to an ngrams Index. It can be initialized
// with a custom store, otherwise in-memory store will be used.
func New(n int, store Store) *Index {
	if store == nil {
		store = NewMemoryStore()
	}

	return &Index{
		N:     n,
		Store: store,
	}
}

// Parse parses a string of words into tokens and indexes the n-grams.
func (i *Index) Parse(str string) error {

	log.Println(str)

	return nil
}

// sanitize
func (i *Index) sanitize(r []rune) []rune {

	return []rune{}
}
