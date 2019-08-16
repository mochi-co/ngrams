package stores

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
	s.internal[key][future]++

	return nil
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

// Any returns a random ngram from the store.
func (s *MemoryStore) Any() (k string, v Variations, err error) {
	s.RLock()
	defer s.RUnlock()

	// Maps are unordered, so far these purposes we can just take
	// whatever is the first in the range. In a more sensitive
	// environment it would be better to use a randomizer.
	for k, v = range s.internal {

		// If you wanted to get a bit clever here, you could implement
		// some kind of check that only selected starter ngrams, but this
		// would require storing metadata with the ngram.
		break
	}

	return
}

// Close gracefully disconnects the store. Because this is just in-memory,
// it will do nothing and return no errors.
func (s *MemoryStore) Close() error {
	return nil
}
