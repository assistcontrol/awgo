package awgo

func Coalesce[T comparable](a, b T) T {
	var empty T
	if a == empty {
		return b
	}
	return a
}
