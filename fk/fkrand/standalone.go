package fkrand

func Digit() int {
	return digit()
}

func Int(length uint32) int {
	return intWithLength(length)
}

func IntInRange(min, max int) int {
	return intInRange(min, max)
}
