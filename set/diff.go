package set

// 找出2个数组对象中不同的元素，并且返回这两个数组的差异元素
func Diff[T comparable](a, b []T) (diffa []T, diffb []T) {
	// 创建一个映射来存储第一个数组的元素
	arr1Map := make(map[T]struct{})

	// 创建一个映射来存储第二个数组的元素
	arr2Map := make(map[T]struct{})

	// 将第一个数组的元素存储到第一个映射中
	for _, v := range a {
		arr1Map[v] = struct{}{}
	}

	// 将第二个数组的元素存储到第二个映射中
	for _, v := range b {
		arr2Map[v] = struct{}{}
	}

	// 遍历第一个数组，如果元素不在第二个数组中，将其添加到第一个差集中
	for _, v := range a {
		if _, exists := arr2Map[v]; !exists {
			diffa = append(diffa, v)
		}
	}

	// 遍历第二个数组，如果元素不在第一个数组中，将其添加到第二个差集中
	for _, v := range b {
		if _, exists := arr1Map[v]; !exists {
			diffb = append(diffb, v)
		}
	}

	return
}

// 找出数组中重复的元素
func FindExistedElment[T comparable](arr []T) (res []T) {
	set := make(map[T]struct{})

	for _, v := range arr {
		if _, has := set[v]; has {
			res = append(res, v)
		}
		set[v] = struct{}{}
	}
	return
}
