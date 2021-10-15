//go:build go1.18
// +build go1.18

package pointer

func To[T any](t T) *T {
	return &t
}

func ToOrNil[T comparable](t T) *T {
	if z, ok := any(t).(interface{ IsZero() bool }); ok {
		if z.IsZero() {
			return nil
		}
		return &t
	}

	var zero T
	if t == zero {
		return nil
	}
	return &t
}

func Get[T any](t *T) T {
	if t == nil {
		var zero T
		return zero
	}
	return *t
}
