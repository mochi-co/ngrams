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

func TestMemoryAdd(t *testing.T) {
	m := NewMemoryStore()
	require.NotNil(t, m)

	err := m.Add("to be", "or")
	require.NoError(t, err)
	require.NotNil(t, m.(*MemoryStore).internal["to be"])
	require.NotNil(t, m.(*MemoryStore).internal["to be"]["or"])
	require.Equal(t, int64(1), m.(*MemoryStore).internal["to be"]["or"])

	// new variation
	err = m.Add("to be", "can")
	require.NoError(t, err)
	require.Equal(t, int64(1), m.(*MemoryStore).internal["to be"]["or"])
	require.Equal(t, int64(1), m.(*MemoryStore).internal["to be"]["can"])

	// increment existing
	err = m.Add("to be", "or")
	require.NoError(t, err)
	require.Equal(t, int64(2), m.(*MemoryStore).internal["to be"]["or"])
	require.Equal(t, int64(1), m.(*MemoryStore).internal["to be"]["can"])

}

func TestMemoryGet(t *testing.T) {
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

func TestMemoryAny(t *testing.T) {
	m := NewMemoryStore()
	require.NotNil(t, m)

	err := m.Add("to be", "or")
	require.NoError(t, err)
	err = m.Add("be or", "not")
	require.NoError(t, err)

	k, v, err := m.Any()
	require.NoError(t, err)
	require.NotEmpty(t, k)
	require.NotEmpty(t, v)
	require.Equal(t, true, k == "to be" || k == "be or")
}

func TestMemoryRemove(t *testing.T) {
	m := &MemoryStore{
		internal: Grams{
			"to be": Variations{
				"or": 3,
				"is": 2,
				"a":  1,
			},
			"be or": Variations{
				"thing": 1,
			},
		},
	}

	require.NotNil(t, m.internal["to be"])
	require.NotNil(t, m.internal["be or"])

	err := m.Delete("to be")
	require.NoError(t, err)

	require.Nil(t, m.internal["to be"])
	require.NotNil(t, m.internal["be or"])
}

func TestMemoryClose(t *testing.T) {
	m := NewMemoryStore()
	err := m.Close()
	require.NoError(t, err)
}
