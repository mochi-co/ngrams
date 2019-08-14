package stores

import (
	"log"
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
	log.Println(results)
	log.Println(v)

	// Fuzzy distribution check. Tolerance +/- 2%
	require.Equal(t, true, (results["or"] > 1900 && results["or"] < 2100))
	require.Equal(t, true, (results["to"] > 2900 && results["to"] < 3100))
	require.Equal(t, true, (results["be"] > 4900 && results["be"] < 5100))

}
