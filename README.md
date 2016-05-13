# pointer [![GoDoc](https://godoc.org/github.com/AlekSi/pointer?status.svg)](https://godoc.org/github.com/AlekSi/pointer) [![Build Status](https://travis-ci.org/AlekSi/pointer.svg)](https://travis-ci.org/AlekSi/pointer)

[Documentation](http://godoc.org/github.com/AlekSi/pointer).


```go
package motivationalexample

import (
	"encoding/json"

	"github.com/AlekSi/pointer"
)

const (
	defaultName = "some name"
)

// Stuff contains optional fields.
type Stuff struct {
	Name    *string
	Comment *string
	Value   *int64
}

// SomeStuff makes some JSON-encoded stuff.
func SomeStuff() (data []byte, err error) {
	return json.Marshal(&Stuff{
		Name:    pointer.ToString(defaultName), // can't say &defaultName
		Comment: pointer.ToString("not yet"),   // can't say &"not yet"
		Value:   pointer.ToInt64(42),           // can't say &42 or &int64(42)
	})
}
```
