package tokenizers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRuneInSlice(t *testing.T) {

	r := []rune{
		46,    // .
		65281, // ！ (cjk)
	}

	require.Equal(t, true, runeInSlice(46, r))
	require.Equal(t, false, runeInSlice(47, r))
	require.Equal(t, true, runeInSlice(65281, r))

}
