package randomizer

func Element[K any](slice []K) K {
	return slice[InRange(0, len(slice))]
}

// Elements returns count distinct elements picked at random from the slice.
// If count is greater than the length of the slice it is clamped to
// len(slice), so the function never panics.
func Elements[K any](count int, slice []K) []K {
	if count > len(slice) {
		count = len(slice)
	}
	if count < 0 {
		count = 0
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
