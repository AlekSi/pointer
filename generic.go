//go:build go1.18
// +build go1.18

package pointer

// To returns a pointer to the passed value.
func To[T any](t T) *T {
	return &t
}

// ToOrNil returns a pointer to the passed value, or nil, if the passed value is a zero value.
// If the passed value has `IsZero() bool` method (for example, time.Time instance),
// it is used to determine if the value is zero.
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

// Get returns the value from the passed pointer or the zero value if the pointer is nil.
func Get[T any](t *T) T {
	if t == nil {
		var zero T
		return zero
	}
	return *t
}

// Deref returns the value from the passed pointer, or nil, if the passed non-nil pointer have a nil value.
func Deref[T any](t *T) any {
	if t == (*T)(nil) {
		return nil
	}
	return *t
}

// GetValue returns the value from the passed pointer, or nil, if the passed non-nil pointer have a nil value.
// Only supports builtin primitive types.
func GetValue(t any) any {
	var x any
	switch v := t.(type) {
	case *bool:
		x = Deref(v)
	case *byte:
		x = Deref(v)
	case *complex64:
		x = Deref(v)
	case *complex128:
		x = Deref(v)
	case *float32:
		x = Deref(v)
	case *float64:
		x = Deref(v)
	case *int:
		x = Deref(v)
	case *int8:
		x = Deref(v)
	case *int16:
		x = Deref(v)
	case *int32:
		x = Deref(v)
	case *int64:
		x = Deref(v)
	case *string:
		x = Deref(v)
	case *uint16:
		x = Deref(v)
	case *uint32:
		x = Deref(v)
	case *uint64:
		x = Deref(v)
	case *uintptr:
		x = Deref(v)
	default:
		x = v
	}
	return x
}
