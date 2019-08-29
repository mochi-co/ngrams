package ngrams

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	stores "github.com/mochi-co/ngrams/stores"
	tk "github.com/mochi-co/ngrams/tokenizers"
)

type MockTokenizer struct {
	tokens    []string
	formatted string
}

func (m *MockTokenizer) Tokenize(str string) []string {
	return m.tokens
}

func (m *MockTokenizer) Scanner(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return
}

func (m *MockTokenizer) Format(tokens []string) string {
	return m.formatted
}

type MockStore struct {
	added    []string
	errAdd   bool
	errClose bool
	errGet   bool
	errAny   bool
}

func (m *MockStore) Add(key, future string) error {
	if m.errAdd {
		return errors.New("test")
	}

	m.added = append(m.added, key+" "+future)

	return nil
}

func (m *MockStore) Get(key string) (bool, stores.Variations) {
	if m.errGet {
		return false, stores.Variations{}
	}
	return true, stores.Variations{"test": 100}
}

func (m *MockStore) Delete(key string) error {
	return nil
}

func (m *MockStore) Any() (string, stores.Variations, error) {
	if m.errAny {
		return "", stores.Variations{}, errors.New("test")
	}
	return "", stores.Variations{}, nil
}

func (m *MockStore) Close() error {
	if m.errClose {
		return errors.New("test")
	}
	return nil
}

func TestNewIndex(t *testing.T) {

	// Force all defaults
	i := NewIndex(0, nil)
	require.NotNil(t, i)
	require.IsType(t, new(stores.MemoryStore), i.Store)
	require.IsType(t, new(tk.DefaultWord), i.Tokenizer)
	require.Equal(t, i.N, defaultN)

	// Custom n value
	i = NewIndex(2, nil)
	require.Equal(t, i.N, 2)

	// Custom Tokenizer and Store
	i = NewIndex(0, &Options{
		Tokenizer: &MockTokenizer{
			formatted: "test",
		},
		Store: new(MockStore),
	})
	require.NotNil(t, i)
	require.Equal(t, "test", i.Tokenizer.Format([]string{}))

	ok, v := i.Store.Get("any")
	require.Equal(t, true, ok)
	require.NotNil(t, v["test"])
	require.Equal(t, int64(100), v["test"])
}

func TestParse(t *testing.T) {
	i := NewIndex(3, nil)
	tokens, err := i.Parse("to be or not to be that is the question")
	require.NoError(t, err)
	require.Equal(t, 10, len(tokens))
	require.Equal(t, "question", tokens[9])

	i = NewIndex(3, &Options{
		Store: &MockStore{
			errAdd: true,
		},
	})
	tokens, err = i.Parse("to be or not to be that is the question")
	require.Error(t, err)

}

func TestRead(t *testing.T) {
	i := NewIndex(3, &Options{
		Store: new(MockStore),
	})

	text := "to be or not to be that is the question"
	r := strings.NewReader(text)
	err := i.Read(r)
	require.NoError(t, err)
	ex := []string{
		"to be or",
		"be or not",
		"or not to",
		"not to be",
		"to be that",
		"be that is",
		"that is the",
		"is the question",
	}
	require.Equal(t, ex, i.Store.(*MockStore).added)

	i.Store = new(MockStore)

	t2 := "to be, or not to be, is that the question?"
	r = strings.NewReader(t2)
	err = i.Read(r)
	require.NoError(t, err)
	ex2 := []string{
		"to be ,",
		"be , or",
		", or not",
		"or not to",
		"not to be",
		"to be ,",
		"be , is",
		", is that",
		"is that the",
		"that the question",
		"the question ?",
	}

	require.Equal(t, ex2, i.Store.(*MockStore).added)

}

func TestClose(t *testing.T) {
	i := NewIndex(3, nil)
	err := i.Close()
	require.NoError(t, err)

	i = NewIndex(3, &Options{
		Store: &MockStore{
			errClose: true,
		},
	})
	err = i.Close()
	require.Error(t, err)
}

func TestExtractNgram(t *testing.T) {
	tokens := []string{"to", "be", "or", "not", "to", "be", ",", "that", "is", "the", "question", "."}

	i := NewIndex(3, nil)
	key, future := i.extractNgram(0, tokens)
	require.Equal(t, "to be", key)
	require.Equal(t, "or", future)

	key, future = i.extractNgram(9, tokens)
	require.Equal(t, "the question", key)
	require.Equal(t, ".", future)

	i = NewIndex(3, nil)
	key, future = i.extractNgram(0, tokens)
	require.Equal(t, "to be", key)
	require.Equal(t, "or", future)

	key, future = i.extractNgram(9, tokens)
	require.Equal(t, "the question", key)
	require.Equal(t, ".", future)

	key, future = i.extractNgram(11, tokens)
	require.Equal(t, "", key)    // blank key for out of range.
	require.Equal(t, "", future) // blank future for out of range.

	i = NewIndex(4, nil)
	key, future = i.extractNgram(0, tokens)
	require.Equal(t, "to be or", key)
	require.Equal(t, "not", future)

	key, future = i.extractNgram(8, tokens)
	require.Equal(t, "is the question", key)
	require.Equal(t, ".", future)

	i = NewIndex(2, nil)
	key, future = i.extractNgram(0, tokens)
	require.Equal(t, "to", key)
	require.Equal(t, "be", future)

	i = NewIndex(1, nil)
	key, future = i.extractNgram(0, tokens)
	require.Equal(t, "to", key)
	require.Equal(t, "to", future) // I don't know why you'd use this for monograms but here we are.

}

func TestExtractAndStore(t *testing.T) {
	tokens := []string{"to", "be", ",", "or", "not", "to", "be", ",", "is", "that", "the", "question", "?"}

	i := NewIndex(3, &Options{
		Store: new(MockStore),
	})
	for j := 0; j < len(tokens); j++ {
		err := i.extractAndStore(j, tokens)
		require.NoError(t, err)
	}

	ex2 := []string{
		"to be ,",
		"be , or",
		", or not",
		"or not to",
		"not to be",
		"to be ,",
		"be , is",
		", is that",
		"is that the",
		"that the question",
		"the question ?",
	}
	require.Equal(t, ex2, i.Store.(*MockStore).added)

}

func TestSeek(t *testing.T) {
	i := NewIndex(2, nil)
	i.Parse("to be or not to be that is the question")

	ok, result := i.Seek("to")
	require.Equal(t, true, ok)
	require.Equal(t, "to", result.Prefix) // this is unusued with bigrams
	require.Equal(t, stores.Variations{"be": 2}, result.Next)

	i = NewIndex(3, nil)
	i.Parse("to be or not to be that is the question")

	ok, result = i.Seek("to be")
	require.Equal(t, true, ok)
	require.Equal(t, "be", result.Prefix)
	require.Equal(t, stores.Variations{"that": 1, "or": 1}, result.Next)

	ok, result = i.Seek("be that")
	require.Equal(t, true, ok)
	require.Equal(t, "that", result.Prefix)
	require.Equal(t, stores.Variations{"is": 1}, result.Next)

	i = NewIndex(4, nil)
	i.Parse("to be or not to be that is the question")

	ok, result = i.Seek("to be or")
	require.Equal(t, true, ok)
	require.Equal(t, "or", result.Prefix)
	require.Equal(t, stores.Variations{"not": 1}, result.Next)

	i = NewIndex(3, &Options{
		Store:     new(MockStore),
		Tokenizer: new(MockTokenizer),
	})
	ok, result = i.Seek("to be or")
	require.Equal(t, true, ok)
	require.Nil(t, result)

}

func TestBabble(t *testing.T) {
	i := NewIndex(3, nil)

	file, err := os.Open("training/pride-prejudice.txt")
	if err != nil {
		require.NoError(t, err)
	}
	defer file.Close()

	d, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	i.Parse(string(d))

	start, _, _ := i.Store.Any()
	b, err := i.Babble(start, 200)
	require.NoError(t, err)
	require.NotEmpty(t, b)
}

func TestShortBabble(t *testing.T) {

	i := NewIndex(3, nil)
	i.Parse("to be or not to be, that is the question.")
	i.Parse("be or not to be something, what is the question?")
	i.Parse("what can we be, or not be, if we ask the question of ourselves.")
	//i.Parse("Mr. Bingley was good-looking and gentlemanlike; he had a pleasant countenance, and easy, unaffected manners.")
	//	i.Parse("To think it more than commonly anxious to get round to the preference of one, and offended by the other as politely and more cheerfully.")
	//	i.Parse("Their visit afforded was produced by the lady with whom she almost looked up to the stables. They were to set out with such a woman.")

	b, err := i.Babble("be something", 200)
	require.NoError(t, err)
	require.NotEmpty(t, b)

	i = NewIndex(3, &Options{
		Store: &MockStore{
			errGet: true,
			errAny: true,
		},
	})
	_, err = i.Babble("", 10)
	require.Error(t, err)

	i = NewIndex(3, &Options{
		Store: &MockStore{
			errGet: true,
		},
	})
	_, err = i.Babble("", 10)
	require.Error(t, err)
	require.Equal(t, ErrEmptyIndex, err)

	i = NewIndex(3, &Options{
		Store: &MockStore{},
	})
	_, err = i.Babble("", 10)
	require.Error(t, err)
	require.Equal(t, ErrNoResult, err)
}
