package set

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubNotInSet(t *testing.T) {
	got := FindSubNotInSet([]string{"a", "b", "c", "d", "e"}, []string{"a", "b", "c"})
	got1 := FindSubNotInSet([]uint{1, 2, 3, 4, 5}, []uint{1, 2, 3})
	assert.Equal(t, []string{"d", "e"}, got)
	assert.Equal(t, []uint{4, 5}, got1)
}

func TestRemoveDuplicatedItem(t *testing.T) {
	got, duplicated := RemoveDuplicatedItem([]string{"a", "a", "b", "b", "b", "c"})
	assert.Equal(t, []string{"a", "b", "c"}, got)
	assert.Equal(t, []string{"a", "b", "b"}, duplicated)
}

func TestContainsExplicit(t *testing.T) {
	got := ContainsExplicit([]string{"a", "b"}, "a")
	got1 := ContainsExplicit([]string{"a", "b"}, "c")
	assert.Equal(t, true, got)
	assert.Equal(t, false, got1)
}

func TestContainsWithFunc(t *testing.T) {
	got := ContainsWithFunc([]string{"abandon", "bell"}, "ab", func(a, b string) bool {
		return strings.HasPrefix(a, b)
	})
	got1 := ContainsWithFunc([]string{"abandon", "bell"}, "cell", func(a, b string) bool {
		return strings.HasPrefix(a, b)
	})
	assert.Equal(t, true, got)
	assert.Equal(t, false, got1)
}
