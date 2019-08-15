package stores

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVariationsNextWeightedRand(t *testing.T) {

	v := &Variations{
		"or": 2,
		"to": 3,
		"be": 5,
	}

	results := map[string]int{
		"or": 0,
		"to": 0,
		"be": 0,
	}
	for i := 0; i < 10000; i++ {
		next := v.NextWeightedRand()
		require.NotEmpty(t, next)
		results[next] += 1
	}

	// Fuzzy distribution check. Tolerance +/- 4%
	require.Equal(t, true, (results["or"] > 1800 && results["or"] < 2200))
	require.Equal(t, true, (results["to"] > 2800 && results["to"] < 3200))
	require.Equal(t, true, (results["be"] > 4800 && results["be"] < 5200))

}
