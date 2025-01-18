package randomizer

func Element[K any](slice []K) K {
	return slice[InRange(0, len(slice)-1)]
}

func Elements[K any](count int, slice []K) []K {
	if count > len(slice) {
		panic("count cannot be greater than the length of the slice")
	}

	result := make([]K, 0, count)
	usedIndexes := make(map[int]bool)

	for len(result) < count {
		index := InRange(0, len(slice))
		if !usedIndexes[index] {
			result = append(result, slice[index])
			usedIndexes[index] = true
		}
	}

	return result
}
