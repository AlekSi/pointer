package pointer

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

/*
Order as in spec:
	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr
	time.Duration time.Time
*/

func TestBool(t *testing.T) {
	var x bool
	if *ToBool(x) != x {
		t.Errorf("*ToBool(%v)", x)
	}
	if ToBoolOrNil(x) != nil {
		t.Errorf("ToBoolOrNil(%v)", x)
	}
	if GetBool(nil) != x {
		t.Errorf("GetBool(%v)", nil)
	}

	x = true
	if *ToBool(x) != x {
		t.Errorf("*ToBool(%v)", x)
	}
	if *ToBoolOrNil(x) != x {
		t.Errorf("*ToBoolOrNil(%v)", x)
	}
	if GetBool(&x) != x {
		t.Errorf("GetBool(%v)", &x)
	}
}

func TestByte(t *testing.T) {
	var x byte
	if *ToByte(x) != x {
		t.Errorf("*ToByte(%v)", x)
	}
	if ToByteOrNil(x) != nil {
		t.Errorf("ToByteOrNil(%v)", x)
	}
	if GetByte(nil) != x {
		t.Errorf("GetByte(%v)", nil)
	}

	x = 42
	if *ToByte(x) != x {
		t.Errorf("*ToByte(%v)", x)
	}
	if *ToByteOrNil(x) != x {
		t.Errorf("*ToByteOrNil(%v)", x)
	}
	if GetByte(&x) != x {
		t.Errorf("GetByte(%v)", &x)
	}
}

func TestComplex64(t *testing.T) {
	var x complex64
	if *ToComplex64(x) != x {
		t.Errorf("*ToComplex64(%v)", x)
	}
	if ToComplex64OrNil(x) != nil {
		t.Errorf("ToComplex64OrNil(%v)", x)
	}
	if GetComplex64(nil) != x {
		t.Errorf("GetComplex64(%v)", nil)
	}

	x = 42
	if *ToComplex64(x) != x {
		t.Errorf("*ToComplex64(%v)", x)
	}
	if *ToComplex64OrNil(x) != x {
		t.Errorf("*ToComplex64OrNil(%v)", x)
	}
	if GetComplex64(&x) != x {
		t.Errorf("GetComplex64(%v)", &x)
	}
}

func TestComplex128(t *testing.T) {
	var x complex128
	if *ToComplex128(x) != x {
		t.Errorf("*ToComplex128(%v)", x)
	}
	if ToComplex128OrNil(x) != nil {
		t.Errorf("ToComplex128OrNil(%v)", x)
	}
	if GetComplex128(nil) != x {
		t.Errorf("GetComplex128(%v)", nil)
	}

	x = 42
	if *ToComplex128(x) != x {
		t.Errorf("*ToComplex128(%v)", x)
	}
	if *ToComplex128OrNil(x) != x {
		t.Errorf("*ToComplex128OrNil(%v)", x)
	}
	if GetComplex128(&x) != x {
		t.Errorf("GetComplex128(%v)", &x)
	}
}

func TestError(t *testing.T) {
	var x error
	if *ToError(x) != x {
		t.Errorf("*ToError(%v)", x)
	}
	if ToErrorOrNil(x) != nil {
		t.Errorf("ToErrorOrNil(%v)", x)
	}
	if GetError(nil) != x {
		t.Errorf("GetError(%v)", nil)
	}

	x = errors.New("error")
	if *ToError(x) != x {
		t.Errorf("*ToError(%v)", x)
	}
	if *ToErrorOrNil(x) != x {
		t.Errorf("*ToErrorOrNil(%v)", x)
	}
	if GetError(&x) != x {
		t.Errorf("GetError(%v)", &x)
	}
}

func TestFloat32(t *testing.T) {
	var x float32
	if *ToFloat32(x) != x {
		t.Errorf("*ToFloat32(%v)", x)
	}
	if ToFloat32OrNil(x) != nil {
		t.Errorf("ToFloat32OrNil(%v)", x)
	}
	if GetFloat32(nil) != x {
		t.Errorf("GetFloat32(%v)", nil)
	}

	x = 42
	if *ToFloat32(x) != x {
		t.Errorf("*ToFloat32(%v)", x)
	}
	if *ToFloat32OrNil(x) != x {
		t.Errorf("*ToFloat32OrNil(%v)", x)
	}
	if GetFloat32(&x) != x {
		t.Errorf("GetFloat32(%v)", &x)
	}
}

func TestFloat64(t *testing.T) {
	var x float64
	if *ToFloat64(x) != x {
		t.Errorf("*ToFloat64(%v)", x)
	}
	if ToFloat64OrNil(x) != nil {
		t.Errorf("ToFloat64OrNil(%v)", x)
	}
	if GetFloat64(nil) != x {
		t.Errorf("GetFloat64(%v)", nil)
	}

	x = 42
	if *ToFloat64(x) != x {
		t.Errorf("*ToFloat64(%v)", x)
	}
	if *ToFloat64OrNil(x) != x {
		t.Errorf("*ToFloat64OrNil(%v)", x)
	}
	if GetFloat64(&x) != x {
		t.Errorf("GetFloat64(%v)", &x)
	}
}

func TestInt(t *testing.T) {
	var x int
	if *ToInt(x) != x {
		t.Errorf("*ToInt(%v)", x)
	}
	if ToIntOrNil(x) != nil {
		t.Errorf("ToIntOrNil(%v)", x)
	}
	if GetInt(nil) != x {
		t.Errorf("GetInt(%v)", nil)
	}

	x = 42
	if *ToInt(x) != x {
		t.Errorf("*ToInt(%v)", x)
	}
	if *ToIntOrNil(x) != x {
		t.Errorf("*ToIntOrNil(%v)", x)
	}
	if GetInt(&x) != x {
		t.Errorf("GetInt(%v)", &x)
	}
}

func TestInt8(t *testing.T) {
	var x int8
	if *ToInt8(x) != x {
		t.Errorf("*ToInt8(%v)", x)
	}
	if ToInt8OrNil(x) != nil {
		t.Errorf("ToInt8OrNil(%v)", x)
	}
	if GetInt8(nil) != x {
		t.Errorf("GetInt8(%v)", nil)
	}

	x = 42
	if *ToInt8(x) != x {
		t.Errorf("*ToInt8(%v)", x)
	}
	if *ToInt8OrNil(x) != x {
		t.Errorf("*ToInt8OrNil(%v)", x)
	}
	if GetInt8(&x) != x {
		t.Errorf("GetInt8(%v)", &x)
	}
}

func TestInt16(t *testing.T) {
	var x int16
	if *ToInt16(x) != x {
		t.Errorf("*ToInt16(%v)", x)
	}
	if ToInt16OrNil(x) != nil {
		t.Errorf("ToInt16OrNil(%v)", x)
	}
	if GetInt16(nil) != x {
		t.Errorf("GetInt16(%v)", nil)
	}

	x = 42
	if *ToInt16(x) != x {
		t.Errorf("*ToInt16(%v)", x)
	}
	if *ToInt16OrNil(x) != x {
		t.Errorf("*ToInt16OrNil(%v)", x)
	}
	if GetInt16(&x) != x {
		t.Errorf("GetInt16(%v)", &x)
	}
}

func TestInt32(t *testing.T) {
	var x int32
	if *ToInt32(x) != x {
		t.Errorf("*ToInt32(%v)", x)
	}
	if ToInt32OrNil(x) != nil {
		t.Errorf("ToInt32OrNil(%v)", x)
	}
	if GetInt32(nil) != x {
		t.Errorf("GetInt32(%v)", nil)
	}

	x = 42
	if *ToInt32(x) != x {
		t.Errorf("*ToInt32(%v)", x)
	}
	if *ToInt32OrNil(x) != x {
		t.Errorf("*ToInt32OrNil(%v)", x)
	}
	if GetInt32(&x) != x {
		t.Errorf("GetInt32(%v)", &x)
	}
}

func TestInt64(t *testing.T) {
	var x int64
	if *ToInt64(x) != x {
		t.Errorf("*ToInt64(%v)", x)
	}
	if ToInt64OrNil(x) != nil {
		t.Errorf("ToInt64OrNil(%v)", x)
	}
	if GetInt64(nil) != x {
		t.Errorf("GetInt64(%v)", nil)
	}

	x = 42
	if *ToInt64(x) != x {
		t.Errorf("*ToInt64(%v)", x)
	}
	if *ToInt64OrNil(x) != x {
		t.Errorf("*ToInt64OrNil(%v)", x)
	}
	if GetInt64(&x) != x {
		t.Errorf("GetInt64(%v)", &x)
	}
}

func TestRune(t *testing.T) {
	var x rune
	if *ToRune(x) != x {
		t.Errorf("*ToRune(%v)", x)
	}
	if ToRuneOrNil(x) != nil {
		t.Errorf("ToRuneOrNil(%v)", x)
	}
	if GetRune(nil) != x {
		t.Errorf("GetRune(%v)", nil)
	}

	x = 'x'
	if *ToRune(x) != x {
		t.Errorf("*ToRune(%v)", x)
	}
	if *ToRuneOrNil(x) != x {
		t.Errorf("*ToRuneOrNil(%v)", x)
	}
	if GetRune(&x) != x {
		t.Errorf("GetRune(%v)", &x)
	}
}

func TestString(t *testing.T) {
	var x string
	if *ToString(x) != x {
		t.Errorf("*ToString(%v)", x)
	}
	if ToStringOrNil(x) != nil {
		t.Errorf("ToStringOrNil(%v)", x)
	}
	if GetString(nil) != x {
		t.Errorf("GetString(%v)", nil)
	}

	x = "x"
	if *ToString(x) != x {
		t.Errorf("*ToString(%v)", x)
	}
	if *ToStringOrNil(x) != x {
		t.Errorf("*ToStringOrNil(%v)", x)
	}
	if GetString(&x) != x {
		t.Errorf("GetString(%v)", &x)
	}
}

func TestUint(t *testing.T) {
	var x uint
	if *ToUint(x) != x {
		t.Errorf("*ToUint(%v)", x)
	}
	if ToUintOrNil(x) != nil {
		t.Errorf("ToUintOrNil(%v)", x)
	}
	if GetUint(nil) != x {
		t.Errorf("GetUint(%v)", nil)
	}

	x = 42
	if *ToUint(x) != x {
		t.Errorf("*ToUint(%v)", x)
	}
	if *ToUintOrNil(x) != x {
		t.Errorf("*ToUintOrNil(%v)", x)
	}
	if GetUint(&x) != x {
		t.Errorf("GetUint(%v)", &x)
	}
}

func TestUint8(t *testing.T) {
	var x uint8
	if *ToUint8(x) != x {
		t.Errorf("*ToUint8(%v)", x)
	}
	if ToUint8OrNil(x) != nil {
		t.Errorf("ToUint8OrNil(%v)", x)
	}
	if GetUint8(nil) != x {
		t.Errorf("GetUint8(%v)", nil)
	}

	x = 42
	if *ToUint8(x) != x {
		t.Errorf("*ToUint8(%v)", x)
	}
	if *ToUint8OrNil(x) != x {
		t.Errorf("*ToUint8OrNil(%v)", x)
	}
	if GetUint8(&x) != x {
		t.Errorf("GetUint8(%v)", &x)
	}
}

func TestUint16(t *testing.T) {
	var x uint16
	if *ToUint16(x) != x {
		t.Errorf("*ToUint16(%v)", x)
	}
	if ToUint16OrNil(x) != nil {
		t.Errorf("ToUint16OrNil(%v)", x)
	}
	if GetUint16(nil) != x {
		t.Errorf("GetUint16(%v)", nil)
	}

	x = 42
	if *ToUint16(x) != x {
		t.Errorf("*ToUint16(%v)", x)
	}
	if *ToUint16OrNil(x) != x {
		t.Errorf("*ToUint16OrNil(%v)", x)
	}
	if GetUint16(&x) != x {
		t.Errorf("GetUint16(%v)", &x)
	}
}

func TestUint32(t *testing.T) {
	var x uint32
	if *ToUint32(x) != x {
		t.Errorf("*ToUint32(%v)", x)
	}
	if ToUint32OrNil(x) != nil {
		t.Errorf("ToUint32OrNil(%v)", x)
	}
	if GetUint32(nil) != x {
		t.Errorf("GetUint32(%v)", nil)
	}

	x = 42
	if *ToUint32(x) != x {
		t.Errorf("*ToUint32(%v)", x)
	}
	if *ToUint32OrNil(x) != x {
		t.Errorf("*ToUint32OrNil(%v)", x)
	}
	if GetUint32(&x) != x {
		t.Errorf("GetUint32(%v)", &x)
	}
}

func TestUint64(t *testing.T) {
	var x uint64
	if *ToUint64(x) != x {
		t.Errorf("*ToUint64(%v)", x)
	}
	if ToUint64OrNil(x) != nil {
		t.Errorf("ToUint64OrNil(%v)", x)
	}
	if GetUint64(nil) != x {
		t.Errorf("GetUint64(%v)", nil)
	}

	x = 42
	if *ToUint64(x) != x {
		t.Errorf("*ToUint64(%v)", x)
	}
	if *ToUint64OrNil(x) != x {
		t.Errorf("*ToUint64OrNil(%v)", x)
	}
	if GetUint64(&x) != x {
		t.Errorf("GetUint64(%v)", &x)
	}
}

func TestUintptr(t *testing.T) {
	var x uintptr
	if *ToUintptr(x) != x {
		t.Errorf("*ToUintptr(%v)", x)
	}
	if ToUintptrOrNil(x) != nil {
		t.Errorf("ToUintptrOrNil(%v)", x)
	}
	if GetUintptr(nil) != x {
		t.Errorf("GetUintptr(%v)", nil)
	}

	x = 42
	if *ToUintptr(x) != x {
		t.Errorf("*ToUintptr(%v)", x)
	}
	if *ToUintptrOrNil(x) != x {
		t.Errorf("*ToUintptrOrNil(%v)", x)
	}
	if GetUintptr(&x) != x {
		t.Errorf("GetUintptr(%v)", &x)
	}
}

func TestDuration(t *testing.T) {
	var x time.Duration
	if *ToDuration(x) != x {
		t.Errorf("*ToDuration(%v)", x)
	}
	if ToDurationOrNil(x) != nil {
		t.Errorf("ToDurationOrNil(%v)", x)
	}
	if GetDuration(nil) != x {
		t.Errorf("GetDuration(%v)", nil)
	}

	x = time.Second
	if *ToDuration(x) != x {
		t.Errorf("*ToDuration(%v)", x)
	}
	if *ToDurationOrNil(x) != x {
		t.Errorf("*ToDurationOrNil(%v)", x)
	}
	if GetDuration(&x) != x {
		t.Errorf("GetDuration(%v)", &x)
	}
}

func TestTime(t *testing.T) {
	var x time.Time
	if *ToTime(x) != x {
		t.Errorf("*ToTime(%v)", x)
	}
	if ToTimeOrNil(x) != nil {
		t.Errorf("ToTimeOrNil(%v)", x)
	}
	if GetTime(nil) != x {
		t.Errorf("GetTime(%v)", nil)
	}

	x = time.Date(2019, 9, 22, 7, 34, 0, 0, time.UTC)
	if *ToTime(x) != x {
		t.Errorf("*ToTime(%v)", x)
	}
	if *ToTimeOrNil(x) != x {
		t.Errorf("*ToTimeOrNil(%v)", x)
	}
	if GetTime(&x) != x {
		t.Errorf("GetTime(%v)", &x)
	}
}

func Example() {
	const (
		defaultName = "some name"
	)

	// Stuff contains optional fields.
	type Stuff struct {
		Name    *string
		Comment *string
		Value   *int64
		Time    *time.Time
	}

	b, _ := json.Marshal(&Stuff{
		Name:    ToString(defaultName),                                   // can't say &defaultName
		Comment: ToString("not yet"),                                     // can't say &"not yet"
		Value:   ToInt64(42),                                             // can't say &42 or &int64(42)
		Time:    ToTime(time.Date(2014, 6, 25, 12, 24, 40, 0, time.UTC)), // can't say &time.Date(…)
	})

	fmt.Printf("%s", b)

	// Output: {"Name":"some name","Comment":"not yet","Value":42,"Time":"2014-06-25T12:24:40Z"}
}
