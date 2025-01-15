package randomizer

func Element[K any](slice []K) K {
	return slice[InRange(0, len(slice)-1)]
}

func MultipleElmenets[K any](count int, slice []K) []K {
	if count >= len(slice) {
		return slice
	}

	selected := make([]K, count)
	for i := 0; i < count; i++ {
		selected[i] = Element(slice)
	}

	return selected
}
