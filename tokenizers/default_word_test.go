package tokenizers

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	tk := NewDefaultWordTokenizer()

	tokens := tk.Tokenize("to be or not to be.")
	require.Equal(t, 7, len(tokens))
	require.Equal(t, "to", tokens[0])
	require.Equal(t, ".", tokens[6])

	tokens = tk.Tokenize("first second third fourth fifth sixth seventh")
	require.NotEmpty(t, tokens)
	require.Equal(t, 7, len(tokens))

	tokens = tk.Tokenize("first, second. thi,rd fourth fifth sixth seventh")
	require.Equal(t, 9, len(tokens))
	require.Equal(t, "thi,rd", tokens[4])

	tokens = tk.Tokenize("Mr. Bingley was good-looking and gentlemanlike; he had a pleasant countenance, and easy, unaffected manners.")
	require.Equal(t, 20, len(tokens))
	require.Equal(t, "Bingley", tokens[2])
	require.Equal(t, "good-looking", tokens[4])
	require.Equal(t, ".", tokens[19])

	tokens = tk.Tokenize("I am sick of Mr. Bingley, cried his wife.")
	require.Equal(t, 12, len(tokens))

	tokens = tk.Tokenize("first")
	require.NotEmpty(t, tokens)
	require.Equal(t, 1, len(tokens))

	tokens = tk.Tokenize("")
	require.Empty(t, tokens)
	require.Equal(t, 0, len(tokens))

	tokens = tk.Tokenize(". ")
	require.NotEmpty(t, tokens)
	require.Equal(t, 1, len(tokens))

	tokens = tk.Tokenize(`“Are you quite sure, ma'am?--is not there a little mistake?” said Jane.
			“I certainly saw Mr. Darcy speaking to her.”`)
	require.NotEmpty(t, tokens)

	tokens = tk.Tokenize(`said Jane.
	Darcy.`)
	require.NotEmpty(t, tokens)
	require.Equal(t, 5, len(tokens))
	require.Equal(t, []string{"said", "Jane", ".", "Darcy", "."}, tokens)

}

func TestSanitize(t *testing.T) {
	tk := NewDefaultWordTokenizer()

	// In the "real world" you might do something a bit more comprehensive than this,
	// maybe a test table (to allow for easier maintenance), but this is simple and serves a purpose.
	str := "  («T[his 『is』 a “stri]n”g) \"int‘e{rspe’rsed wit}h „removable“ 「characters」.»  "
	require.Equal(t, "This is a string interspersed with removable characters.", tk.sanitize(str))

}

func TestFormat(t *testing.T) {
	tk := NewDefaultWordTokenizer()

	tokens := []string{"i", "am", "sick", "of", "Mr", ".", "Bingley", ",", "cried", "his", "wife", ".", "he's", "like", "a", "character", "from", "a", "Jane", "Austen", "novel", "!"}

	require.Equal(t, "I am sick of Mr. Bingley, cried his wife. He's like a character from a Jane Austen novel!", tk.Format(tokens))

}
