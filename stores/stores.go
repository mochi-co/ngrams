package stores

// Store is a data storage mechanism for ngrams.
type Store interface {

	// Add adds a new ngram and key-variation pair to the index.
	Add(string, []string) error

	// Get returns the potential ngram variations for a key.
	Get(string) (bool, map[string]int64)

	// Delete removes an ngram key and all variations from the index.
	Delete(string) error
}

// NGram is a map of Potential ngrams keyed on gram-key (eg. "to be").
// This is primarily used by the in-memory store, but can also be used to
// structure data for other storage engines.
type NGram map[string]Potential

// Potential contains the future-sequenced ngrams and the number of times
// they were indexed, keyed on gram-variation (eg. {"be or":3})
type Potential map[string]int64
