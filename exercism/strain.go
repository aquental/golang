package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](list []T, filter func(T) bool) []T {
	var result []T
	for _, value := range list {
		if filter(value) {
			result = append(result, value)
		}
	}
	return result
}

func Discard[T any](list []T, filter func(T) bool) []T {
	return Keep(list, func(value T) bool { return !filter(value) })
}

