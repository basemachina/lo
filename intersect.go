package lo

import "slices"

// copied from: https://github.com/samber/lo/blob/v1.38.1/intersect.go

// Every returns true if all elements of a subset are contained into a collection or if the subset is empty.
func Every[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if !slices.Contains(collection, elem) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained into a collection.
// If the subset is empty Some returns false.
func Some[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if slices.Contains(collection, elem) {
			return true
		}
	}

	return false
}

// Intersect returns the intersection between two collections.
func Intersect[T comparable, Slice ~[]T](list1 Slice, list2 Slice) Slice {
	result := Slice{}
	seen := map[T]struct{}{}

	for i := range list1 {
		seen[list1[i]] = struct{}{}
	}

	for i := range list2 {
		if _, ok := seen[list2[i]]; ok {
			result = append(result, list2[i])
		}
	}

	return result
}
