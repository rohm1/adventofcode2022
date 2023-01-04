package utils

func Remove[T comparable](list []T, item T) []T {
	if len(list) == 0 {
		return make([]T, 0)
	}

	newList := make([]T, 0)
	for i := 0; i < len(list); i++ {
		if list[i] != item {
			newList = append(newList, list[i])
		}
	}
	return newList
}
