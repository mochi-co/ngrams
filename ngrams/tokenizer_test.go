package ngrams

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	tokens := Tokenize("to be or not to be.")
	require.Equal(t, 7, len(tokens))
	require.Equal(t, "to", tokens[0])
	require.Equal(t, ".", tokens[6])

	tokens = Tokenize("first second third fourth fifth sixth seventh")
	require.NotEmpty(t, tokens)
	require.Equal(t, 7, len(tokens))

	tokens = Tokenize("first, second. thi,rd fourth fifth sixth seventh")
	require.Equal(t, 9, len(tokens))
	require.Equal(t, "thi,rd", tokens[4])

	tokens = Tokenize("Mr. Bingley was good-looking and gentlemanlike; he had a pleasant countenance, and easy, unaffected manners.")
	require.Equal(t, 20, len(tokens))
	require.Equal(t, "Bingley", tokens[2])
	require.Equal(t, "good-looking", tokens[4])
	require.Equal(t, ".", tokens[19])

	tokens = Tokenize("I am sick of Mr. Bingley, cried his wife.")
	require.Equal(t, 12, len(tokens))
}

func TestIsSkippableRune(t *testing.T) {
	require.Equal(t, true, isSkippableRune(9))     // tab
	require.Equal(t, true, isSkippableRune(10))    // newline
	require.Equal(t, true, isSkippableRune(13))    // return
	require.Equal(t, true, isSkippableRune(32))    // space
	require.Equal(t, true, isSkippableRune(12288)) // space (cjk)
	require.Equal(t, false, isSkippableRune(44))   // ,
	require.Equal(t, false, isSkippableRune(45))   // -
	require.Equal(t, false, isSkippableRune(46))   // .
	require.Equal(t, false, isSkippableRune(92))   // \
	require.Equal(t, false, isSkippableRune(49))   // 1
	require.Equal(t, false, isSkippableRune(50))   // 2
	require.Equal(t, false, isSkippableRune(97))   // a
	require.Equal(t, false, isSkippableRune(98))   // b
}

func TestIsPunctuation(t *testing.T) {
	require.Equal(t, true, isPunctuation(46))   // .
	require.Equal(t, true, isPunctuation(44))   // ,
	require.Equal(t, true, isPunctuation(63))   // ?
	require.Equal(t, true, isPunctuation(33))   // !
	require.Equal(t, true, isPunctuation(8253)) // ‽
	require.Equal(t, true, isPunctuation(58))   // :
	require.Equal(t, true, isPunctuation(59))   // ;

	require.Equal(t, true, isPunctuation(12290)) // 。 (cjk)
	require.Equal(t, true, isPunctuation(12289)) //  、 (cjk)
	require.Equal(t, true, isPunctuation(65281)) // ！ (cjk)
	require.Equal(t, true, isPunctuation(65311)) // ？ (cj
	require.Equal(t, true, isPunctuation(65306)) // ： (cjk)
	require.Equal(t, true, isPunctuation(65307)) // ； (cjk)

	require.Equal(t, false, isPunctuation(9))     // tab
	require.Equal(t, false, isPunctuation(10))    // newline
	require.Equal(t, false, isPunctuation(13))    // return
	require.Equal(t, false, isPunctuation(32))    // space
	require.Equal(t, false, isPunctuation(12288)) // space (cjk)
	require.Equal(t, false, isPunctuation(45))    // -
	require.Equal(t, false, isPunctuation(95))    // _
}
