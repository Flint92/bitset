package bitset

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBitSet2(t *testing.T) {
	var x, y BitSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	require.Equal(t, x.String(), "{1 9 144}")
	require.Equal(t, x.Len(), 3)
	require.Equal(t, x.Has(9), true)
	require.Equal(t, x.Has(123), false)

	y.Add(9)
	y.Add(42)
	require.Equal(t, y.String(), "{9 42}")

	x.UnionWith(&y)
	require.Equal(t, x.String(), "{1 9 42 144}")
	require.Equal(t, x.Has(9), true)
	require.Equal(t, x.Has(123), false)

	remove1 := x.Remove(5)
	require.Equal(t, remove1, false)

	remove2 := x.Remove(9)
	require.Equal(t, remove2, true)
	require.Equal(t, x.String(), "{1 42 144}")
	require.Equal(t, x.Len(), 3)

	x.Clear()
	require.Equal(t, x.String(), "{}")
	require.Equal(t, x.Len(), 0)
}

func TestBitSet_Copy(t *testing.T) {
	var x BitSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y := x.Copy()
	require.Equal(t, x.String(), y.String())
	require.NotEqual(t, &x, &y)

}


func TestBitSet_Elements(t *testing.T) {
	var x BitSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	elements := x.Elements()
	require.Equal(t, elements[0], 1)
	require.Equal(t, elements[1], 9)
	require.Equal(t, elements[2], 144)
}
