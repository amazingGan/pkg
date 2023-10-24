package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	args1 := []uint32{1, 2, 3, 4, 5}
	args2 := []uint32{2, 3, 5, 6, 7}
	wantDiffA := []uint32{1, 4}
	wantDiffB := []uint32{6, 7}
	diffa, diffb := Diff[uint32](args1, args2)
	assert.Equal(t, wantDiffA, diffa)
	assert.Equal(t, wantDiffB, diffb)

	args3 := []string{"A", "b", "C", "D", "e"}
	args4 := []string{"D", "c", "C", "D", "A"}
	wantDiffC := []string{"b", "e"}
	wantDiffD := []string{"c"}
	diffc, diffd := Diff[string](args3, args4)
	assert.Equal(t, wantDiffC, diffc)
	assert.Equal(t, wantDiffD, diffd)
}

func TestFindExistedElment(t *testing.T) {
	got := FindExistedElment([]uint32{1, 3, 3, 2, 5, 5})
	assert.Equal(t, []uint32{3, 5}, got)

	got1 := FindExistedElment([]string{"A", "b", "B", "b", "A", "C"})
	assert.Equal(t, []string{"b", "A"}, got1)
}
