package ngrams

import (
	"io/ioutil"
	"log"
	"os"
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

	i = NewIndex(3, Options{})
	key, future = i.extractNgram(0, tokens)
	require.Equal(t, "to be", key)
	require.Equal(t, "or", future)

	key, future = i.extractNgram(9, tokens)
	require.Equal(t, "the question", key)
	require.Equal(t, ".", future)

	key, future = i.extractNgram(11, tokens)
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

func TestSeek(t *testing.T) {
	i := NewIndex(2, Options{})
	i.Parse("to be or not to be that is the question")

	ok, result := i.Seek("to")
	require.Equal(t, true, ok)
	require.Equal(t, "to", result.Prefix) // this is unusued with bigrams
	require.Equal(t, stores.Variations{"be": 2}, result.Next)

	i = NewIndex(3, Options{})
	i.Parse("to be or not to be that is the question")

	ok, result = i.Seek("to be")
	require.Equal(t, true, ok)
	require.Equal(t, "be", result.Prefix)
	require.Equal(t, stores.Variations{"that": 1, "or": 1}, result.Next)

	ok, result = i.Seek("be that")
	require.Equal(t, true, ok)
	require.Equal(t, "that", result.Prefix)
	require.Equal(t, stores.Variations{"is": 1}, result.Next)

	i = NewIndex(4, Options{})
	i.Parse("to be or not to be that is the question")

	ok, result = i.Seek("to be or")
	require.Equal(t, true, ok)
	require.Equal(t, "or", result.Prefix)
	require.Equal(t, stores.Variations{"not": 1}, result.Next)
}

func TestBabble(t *testing.T) {
	i := NewIndex(3, Options{})

	/*
		i := NewIndex(3, Options{})
		i.Parse("to be or not to be, that is the question.")
		i.Parse("be or not to be something, what is the question?")
		i.Parse("what can we be, or not be, if we ask the question of ourselves.")
		//i.Parse("Mr. Bingley was good-looking and gentlemanlike; he had a pleasant countenance, and easy, unaffected manners.")
		//	i.Parse("To think it more than commonly anxious to get round to the preference of one, and offended by the other as politely and more cheerfully.")
		//	i.Parse("Their visit afforded was produced by the lady with whom she almost looked up to the stables. They were to set out with such a woman.")

		//i.Store.(*stores.MemoryStore).Print()

		b := i.Babble("to be", 20)
		log.Println("##", b)
	*/

	// Read Austen
	file, err := os.Open("training/pride_prejudice.txt")
	if err != nil {
		require.NoError(t, err)
	}
	defer file.Close()

	d, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	i.Parse(string(d))

	start, _ := i.Store.Any()
	b := i.Babble(start, 60)
	log.Println("##", b)
}
