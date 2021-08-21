package pointer

func To[T any](t T) *T {
	return &t
}

func Get[T any](t *T) T {
	if t == nil {
		var zero T
		return zero
	}
	return *t
}

// No `func ToOrNil[T any](t T) *T` due to time.Time.Equal
