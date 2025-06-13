package lo

// copied from: https://github.com/samber/lo/blob/v1.38.1/type_manipulation.go

// ToPtr returns a pointer copy of value.
func ToPtr[T any](x T) *T {
	return &x
}

// FromPtrOr returns the pointer value or the fallback value.
func FromPtrOr[T any](x *T, fallback T) T {
	if x == nil {
		return fallback
	}

	return *x
}
