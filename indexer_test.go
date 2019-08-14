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

func TestParse(t *testing.T) {

	i := NewIndex(3, Options{})
	i.Parse("to be or not to be that is the question")

}

func TestExtractNgram(t *testing.T) {
	tokens := []string{"to", "be", "or", "not", "to", "be", ",", "that", "is", "the", "question", "."}

	i := NewIndex(3, Options{})
	key, future := i.extractNgram(0, tokens)
	require.Equal(t, "to be", key)
	require.Equal(t, "or", future)

	key, future = i.extractNgram(9, tokens)
	require.Equal(t, "the question", key)
	require.Equal(t, ".", future)

	key, future = i.extractNgram(10, tokens)
	require.Equal(t, "", key)    // blank key for out of range.
	require.Equal(t, "", future) // blank future for out of range.

	i = NewIndex(4, Options{})
	key, future = i.extractNgram(0, tokens)
	require.Equal(t, "to be or", key)
	require.Equal(t, "not", future)

	key, future = i.extractNgram(8, tokens)
	require.Equal(t, "is the question", key)
	require.Equal(t, ".", future)

	i = NewIndex(2, Options{})
	key, future = i.extractNgram(0, tokens)
	require.Equal(t, "to", key)
	require.Equal(t, "be", future)

	i = NewIndex(1, Options{})
	key, future = i.extractNgram(0, tokens)
	require.Equal(t, "to", key)
	require.Equal(t, "to", future) // I don't know why you'd use this for monograms but here we are.

}
