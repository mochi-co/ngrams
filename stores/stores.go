package stores

import (
	"math/rand"
	"time"
)

// Store is a data storage mechanism for ngrams.
type Store interface {

	// Connect would be used to setup any connections to an external
	// data store. However, since there are no implementations of that
	// yet, it's noted here as a stub.
	// Connect(options interface{}) error

	// Add adds a new ngram and key-variation pair to the index. It
	// should take the key, which will be used to index the gram, and
	// the future, which is a slice of consequent gram tokens.
	// I like to add labels to my interface method inputs because it
	// makes the context a little easier to understand for whoever
	// follows my footsteps (sometimes that's future-me).
	Add(key, future string) error

	// Get returns the potential ngram variations for a key.
	Get(key string) (bool, Variations)

	// Delete removes an ngram key and all variations from the index.
	Delete(key string) error

	// Any returns a random ngram from the store.
	Any() (string, Variations, error)

	// Len returns the number of tokens in the store.

	// Close is used to gracefully shutdown any connections.
	Close() error
}

// Grams is a map of Variations keyed on gram-key (eg. "to be").
// This is primarily used by the in-memory store, but can also be used to
// structure data for other storage engines.
// Fun fact; the `gram` in n-gram derives from the Latin `gramma`, which
// is something that is written or drawn (such as a letter or number).
type Grams map[string]Variations

// Variations contains the future-sequenced ngrams and the number of times
// they were indexed, keyed on gram-variation (eg. {"be or":3})
type Variations map[string]int64

// NextWeightedRand returns a random variation, probability-weighted by the
// number of times it was indexed.
// Using a linear scan ala https://blog.bruce-hill.com/a-faster-weighted-random-choice
func (v *Variations) NextWeightedRand() string {
	rand.Seed(time.Now().UnixNano())

	// Get a sum total of the probabities of all variations.
	var total int64
	for _, i := range *v {
		total += i
	}

	// Pick a random int between 0 and the total sum.
	r := rand.Int63n(total)

	// Range through the possible variations and subtract the probability
	// weight from the random number. If r goes below zero, select the key.
	var k string
	var i int64
	for k, i = range *v {
		r -= i
		if r < 0 {
			break
		}
	}

	return k
}
