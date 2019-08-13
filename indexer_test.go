package ngrams

import (
	"testing"

	"github.com/stretchr/testify/require"

	stores "github.com/mochi-co/trigrams-test/stores"
	tk "github.com/mochi-co/trigrams-test/tokenizers"
)

func TestNewIndex(t *testing.T) {

	// Force all defaults
	i := NewIndex(0, Options{})
	require.NotNil(t, i)
	require.IsType(t, new(stores.MemoryStore), i.Store)
	require.IsType(t, new(tk.DefaultWord), i.Tokenizer)
	require.Equal(t, i.N, defaultN)

	// Custom n value
	i = NewIndex(2, Options{})
	require.Equal(t, i.N, 2)

	// Custom Tokenizer
	// ...

	// Custom Store
	// ...
}

func TestIndexParse(t *testing.T) {

}

/*

// Tests on struct methods take the name of TestStructMethod to keep them unique.
func TestIndexParse(t *testing.T) {

	b1 := "first second third fourth fifth sixth seventh 「characters」"

	s := New(3, NewMemoryStore())
	require.NotNil(t, s)
	require.IsType(t, new(MemoryStore), s.Store)
	require.Equal(t, s.N, 3)

	err := s.Parse(b1)
	require.NoError(t, err)

}
*/
