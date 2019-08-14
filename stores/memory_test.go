package stores

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMemoryStore(t *testing.T) {
	m := NewMemoryStore()
	require.NotNil(t, m)
	require.IsType(t, new(MemoryStore), m)
}

func TestAdd(t *testing.T) {
	m := NewMemoryStore()
	require.NotNil(t, m)

	err := m.Add("to be", "or")
	require.NoError(t, err)
	require.NotNil(t, m.(*MemoryStore).internal["to be"])
	require.NotNil(t, m.(*MemoryStore).internal["to be"]["or"])
	require.Equal(t, int64(1), m.(*MemoryStore).internal["to be"]["or"])

	log.Printf("%#v\n", m.(*MemoryStore).internal)
}

func TestGet(t *testing.T) {
	f := Variations{
		"or": 3,
		"is": 2,
		"a":  1,
	}

	m := &MemoryStore{
		internal: Grams{
			"to be": f,
		},
	}

	ok, v := m.Get("to be")
	require.Equal(t, true, ok)
	require.Equal(t, f, v)

}

func TestRemove(t *testing.T) {

}
