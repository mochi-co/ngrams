package ngrams

import (
	"sync"
)

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
