package stores

// Store is a data storage mechanism for ngrams.
type Store interface {

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
