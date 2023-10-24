package set

import (
	"reflect"
	"sort"
)

func IsEqual[T comparable](arr1, arr2 []T, less func(T, T) bool) bool {
	// 如果数组长度不相等，它们一定不相同
	if len(arr1) != len(arr2) {
		return false
	}
	// 排序后判断是否一致
	sort.SliceStable(arr1, func(i, j int) bool {
		return less(arr1[i], arr1[j])
	})
	sort.SliceStable(arr2, func(i, j int) bool {
		return less(arr2[i], arr2[j])
	})
	return reflect.DeepEqual(arr1, arr2)
}
