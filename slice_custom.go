package lo

// Almost all functions are based on: https://github.com/samber/lo/blob/v1.38.1/slice.go
// MapWithError and MapWithIndexError are based on PR by @NilsJPWerner: https://github.com/samber/lo/pull/43

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T) R) []R {
	return MapWithIndex(collection, func(item T, _ int) R {
		return iteratee(item)
	})
}

// MapWithIndex manipulates a slice and transforms it to a slice of another type.
func MapWithIndex[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

// MapWithError manipulates a slice and transforms it to a slice of another type. If iteratee
// returns an error the function shortcircuits and the error is returned alongside nil.
func MapWithError[T any, R any](collection []T, iteratee func(item T) (R, error)) ([]R, error) {
	return MapWithIndexError(collection, func(item T, _ int) (R, error) {
		return iteratee(item)
	})
}

// MapWithIndexError manipulates a slice and transforms it to a slice of another type. If iteratee
// returns an error the function shortcircuits and the error is returned alongside nil.
func MapWithIndexError[T any, R any](collection []T, iteratee func(item T, index int) (R, error)) ([]R, error) {
	result := make([]R, len(collection))
	var err error

	for i, item := range collection {
		result[i], err = iteratee(item, i)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// HasDuplicates checks whether the given slice contains duplicate values and returns a boolean result.
func HasDuplicates[T comparable](collection []T) bool {
	isDupl := make(map[T]bool, len(collection))

	for _, item := range collection {
		_, ok := isDupl[item]
		if ok {
			return true
		}
		isDupl[item] = true
	}

	return false
}

// HasDuplicatesBy determines whether the provided slice contains duplicate values by examining the return value from the passed function and returns a boolean result.
func HasDuplicatesBy[T any, U comparable](collection []T, iteratee func(item T) U) bool {
	isDupl := make(map[U]bool, len(collection))

	for _, item := range collection {
		key := iteratee(item)

		_, ok := isDupl[key]
		if ok {
			return true
		}
		isDupl[key] = true
	}

	return false
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
// The transform function can either return a slice or a `nil`, and in the `nil` case
// no value is added to the final slice.
func FlatMap[T any, R any](collection []T, iteratee func(item T) []R) []R {
	return FlatMapWithIndex(collection, func(item T, _ int) []R {
		return iteratee(item)
	})
}

// FlatMapWithIndex manipulates a slice and transforms and flattens it to a slice of another type.
// The transform function can either return a slice or a `nil`, and in the `nil` case
// no value is added to the final slice.
func FlatMapWithIndex[T any, R any](collection []T, iteratee func(item T, index int) []R) []R {
	result := make([]R, 0, len(collection))

	for i, item := range collection {
		result = append(result, iteratee(item, i)...)
	}

	return result
}

// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[V any](collection []V, predicate func(item V) bool) []V {
	return FilterWithIndex(collection, func(item V, _ int) bool {
		return predicate(item)
	})
}

// FilterWithIndex iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func FilterWithIndex[V any](collection []V, predicate func(item V, index int) bool) []V {
	result := make([]V, 0, len(collection))

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T) R, initial R) R {
	return ReduceWithIndex(collection, func(agg R, item T, _ int) R {
		return accumulator(agg, item)
	}, initial)
}

// ReduceWithIndex reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
// The accumulator function receives the index of each element.
func ReduceWithIndex[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i := range collection {
		initial = accumulator(initial, collection[i], i)
	}

	return initial
}
