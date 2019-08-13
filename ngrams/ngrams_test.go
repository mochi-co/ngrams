package ngrams

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewIndex(t *testing.T) {

	// Test with in-memory store.
	s := New(3, nil)
	require.NotNil(t, s)
	require.IsType(t, new(MemoryStore), s.Store)
	require.Equal(t, s.N, 3)

	// Test with persistent store.
	// ...

}

// Tests on struct methods take the name of TestStructMethod to keep them unique.
func TestIndexParse(t *testing.T) {

	b1 := "first second third fourth fifth sixth seventh"

	s := New(3, NewMemoryStore())
	require.NotNil(t, s)
	require.IsType(t, new(MemoryStore), s.Store)
	require.Equal(t, s.N, 3)

	err := s.Parse(b1)
	require.NoError(t, err)

}
