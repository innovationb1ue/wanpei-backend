package utils

func Contains[T comparable](elems []T, ele T) bool {
	for _, e := range elems {
		if e == ele {
			return true
		}
	}
	return false
}
