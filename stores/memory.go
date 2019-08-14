package stores

import (
	"log"
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
func (s *MemoryStore) Add(key, future string) error {
	s.RLock()
	defer s.RUnlock()

	// If this particular key doesn't exist at all, we can add it with
	// the provided future, and a starting quantity of 1.
	if _, ok := s.internal[key]; !ok {
		s.internal[key] = Variations{
			future: 1,
		}
		return nil
	}

	// If the gram _does_ exist, then we need to add the variation if it
	// doesn't exist, and then ensure the variation quantity is incremented.
	if _, ok := s.internal[key][future]; !ok {
		s.internal[key][future] = 0
	}
	s.internal[key][future] += 1

	return nil
}

func (s *MemoryStore) Print() {
	log.Printf("%+v\n", s.internal)
}

// Get gets an ngram variation from the store.
func (s *MemoryStore) Get(key string) (ok bool, v Variations) {
	s.Lock()
	defer s.Unlock()

	v, ok = s.internal[key]

	return
}

// Delete removes an ngram to the store.
func (s *MemoryStore) Delete(key string) error {
	s.Lock()
	defer s.Unlock()

	delete(s.internal, key)

	return nil
}
