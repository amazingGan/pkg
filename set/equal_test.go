package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqual(t *testing.T) {
	arr1 := []uint32{1, 2, 3, 5, 4}
	arr2 := []uint32{2, 5, 4, 1, 3}

	assert.Equal(t, true,
		IsEqual[uint32](arr1, arr2, func(u1, u2 uint32) bool {
			return u1 < u2
		}),
	)

	arr3 := []string{"a", "d", "e", "b"}
	arr4 := []string{"c", "b", "e", "a"}

	assert.Equal(t, false,
		IsEqual[string](arr3, arr4, func(u1, u2 string) bool {
			return u1 < u2
		}),
	)

	arr5 := []string{"a", "d", "e", "b"}
	arr6 := []string{"d", "b", "e", "a"}

	assert.Equal(t, true,
		IsEqual[string](arr5, arr6, func(u1, u2 string) bool {
			return u1 < u2
		}),
	)

	arr7 := []int{1, 2, 3, 4, 5}
	arr8 := []int{1, 2, 3}

	assert.Equal(t, false,
		IsEqual[int](arr7, arr8, func(u1, u2 int) bool {
			return u1 < u2
		}),
	)

}
