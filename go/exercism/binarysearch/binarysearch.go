package binarysearch

func searchFromHalf(list []int, key int, idx int, seen *map[int]bool) int {
	if idx >= len(list) || (*seen)[idx] == true {
		return -1
	}

	(*seen)[idx] = true
	halfValue := list[idx]

	if key > halfValue {
		newIdx := idx + len(list[idx:])/2
		return searchFromHalf(list, key, newIdx, seen)
	}

	if key < halfValue {
		newIdx := len(list[:idx]) / 2
		return searchFromHalf(list, key, newIdx, seen)
	}

	return idx
}

func SearchInts(list []int, key int) int {
	seen := make(map[int]bool)
	return searchFromHalf(list, key, (len(list)-1)/2, &seen)
}
