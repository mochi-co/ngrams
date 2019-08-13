package tokenizers

import (
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
}

func TestIsSkippableRune(t *testing.T) {
	tk := NewDefaultWordTokenizer()

	require.Equal(t, true, tk.isSkippableRune(9))     // tab
	require.Equal(t, true, tk.isSkippableRune(10))    // newline
	require.Equal(t, true, tk.isSkippableRune(13))    // return
	require.Equal(t, true, tk.isSkippableRune(32))    // space
	require.Equal(t, true, tk.isSkippableRune(12288)) // space (cjk)
	require.Equal(t, false, tk.isSkippableRune(44))   // ,
	require.Equal(t, false, tk.isSkippableRune(45))   // -
	require.Equal(t, false, tk.isSkippableRune(46))   // .
	require.Equal(t, false, tk.isSkippableRune(92))   // \
	require.Equal(t, false, tk.isSkippableRune(49))   // 1
	require.Equal(t, false, tk.isSkippableRune(50))   // 2
	require.Equal(t, false, tk.isSkippableRune(97))   // a
	require.Equal(t, false, tk.isSkippableRune(98))   // b
}

func TestIsPunctuation(t *testing.T) {
	tk := NewDefaultWordTokenizer()

	require.Equal(t, true, tk.isPunctuation(46))   // .
	require.Equal(t, true, tk.isPunctuation(44))   // ,
	require.Equal(t, true, tk.isPunctuation(63))   // ?
	require.Equal(t, true, tk.isPunctuation(33))   // !
	require.Equal(t, true, tk.isPunctuation(8253)) // ‽
	require.Equal(t, true, tk.isPunctuation(58))   // :
	require.Equal(t, true, tk.isPunctuation(59))   // ;
	require.Equal(t, true, tk.isPunctuation(38))   // &

	require.Equal(t, true, tk.isPunctuation(12290)) // 。 (cjk)
	require.Equal(t, true, tk.isPunctuation(12289)) //  、 (cjk)
	require.Equal(t, true, tk.isPunctuation(65281)) // ！ (cjk)
	require.Equal(t, true, tk.isPunctuation(65311)) // ？ (cj
	require.Equal(t, true, tk.isPunctuation(65306)) // ： (cjk)
	require.Equal(t, true, tk.isPunctuation(65307)) // ； (cjk)

	require.Equal(t, false, tk.isPunctuation(9))     // tab
	require.Equal(t, false, tk.isPunctuation(10))    // newline
	require.Equal(t, false, tk.isPunctuation(13))    // return
	require.Equal(t, false, tk.isPunctuation(32))    // space
	require.Equal(t, false, tk.isPunctuation(12288)) // space (cjk)
	require.Equal(t, false, tk.isPunctuation(45))    // -
	require.Equal(t, false, tk.isPunctuation(95))    // _
}

func TestSanitize(t *testing.T) {
	tk := NewDefaultWordTokenizer()

	// In the "real world" you might do something a bit more comprehensive than this,
	// maybe a test table (to allow for easier maintenance), but this is simple and serves a purpose.
	str := "  («T[his 『is』 a “stri]n”g) \"int‘e{rspe’rsed wit}h „removable“ 「characters」.»  "
	require.Equal(t, "This is a string interspersed with removable characters.", tk.sanitize(str))

}
