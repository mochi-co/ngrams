package ngrams

/*
import (
	"sync"
)

// Grams is a map of Gram keyed on gram-key (eg. "to be"). This is
// primarily used by the in-memory store, but can also be used to
// structure data for other storage engines.
type Grams map[string]Futures

// Futures contains the future-sequenced ngrams and the number of times
// they were indexed, keyed on gram-variation (eg. {"be or":3})
type Futures map[string]int64

// NewMemoryStore returns an in-memory ngram store. Ngrams added to the store
// are not persisted when the service restarts.
func NewMemoryStore() Store {

	return &MemoryStore{
		internal: make(Grams),
	}

}

// MemoryStore is an in-memory ngram store. It complies with Store interface.
type MemoryStore struct {

	// sync implements a mutex for concurrent read/write.
	sync.RWMutex

	// internal contains the indexed grams.
	internal Grams
}

// Add adds an ngram to the store.
func (s *MemoryStore) Add(key string, grams []string) error {
	s.RLock()
	defer s.RUnlock()

	return nil
}

// Get gets an ngram from the store.
func (s *MemoryStore) Get(key string) (ok bool, f Futures) {
	s.Lock()
	defer s.Unlock()

	f, ok = s.internal[key]

	return
}

// Delete removes an ngram to the store.
func (s *MemoryStore) Delete(key string) error {
	s.Lock()
	defer s.Unlock()

	delete(s.internal, key)

	return nil
}

*/
