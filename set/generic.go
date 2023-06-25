package set

// FindSubNotInSet 找出a不在集合b中的子集合
func FindSubNotInSet[T comparable](a, b []T) []T {
	dp := make(map[any]uint64, 0)

	for _, item := range b {
		if v, ok := dp[item]; ok {
			v++
			dp[item] = v
		} else {
			dp[item] = 1
		}
	}

	var res []T

	for _, item := range a {
		if _, ok := dp[item]; !ok {
			res = append(res, item)
		}
	}
	return res
}

// RemoveDuplicatedItem 去除集合中重复的项 返回去重后的集合以及重复的对象集合
func RemoveDuplicatedItem[T comparable](a []T) (res []T, duplicated []T) {
	dp := make(map[T]struct{})

	for _, item := range a {
		if _, ok := dp[item]; ok {
			// 重复项记录
			duplicated = append(duplicated, item)
		} else {
			dp[item] = struct{}{}
		}
	}

	for k := range dp {
		res = append(res, k)
	}

	return
}

// 定义条件函数 可以自定义包含判断逻辑，例如：两个对象相等，或者正则匹配模糊匹配等等
type ConditionFunc[T comparable] func(T, T) bool

func ContainsExplicit[T comparable](set []T, item T) bool {
	return ContainsWithFunc(set, item, func(a, b T) bool {
		return a == b
	})
}

func ContainsWithFunc[T comparable](set []T, item T, fn ConditionFunc[T]) bool {
	for _, element := range set {
		if fn(element, item) {
			return true
		}
	}
	return false
}
