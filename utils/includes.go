package utils

func Includes[T comparable](list []T, item T) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return true
		}
	}
	return false
}
