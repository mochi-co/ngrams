package stores

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMemoryStore(t *testing.T) {
	m := NewMemoryStore()
	require.NotNil(t, m)
	require.IsType(t, new(MemoryStore), m)
}

func TestAdd(t *testing.T) {

	//s1 := "first second third fourth fifth sixth seventh"

}

func TestGet(t *testing.T) {

}

func TestRemove(t *testing.T) {

}
