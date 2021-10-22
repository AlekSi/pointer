//go:build go1.18
// +build go1.18

package pointer

import (
	"testing"
	"time"
)

func TestGeneric(t *testing.T) {
	var x time.Time
	if *To(x) != x {
		t.Errorf("*To(%v)", x)
	}
	if ToOrNil(x) != nil {
		t.Errorf("ToOrNil(%v)", x)
	}
	if Get((*time.Time)(nil)) != x {
		t.Errorf("Time(%v)", nil)
	}

	x = time.Date(2014, 6, 25, 12, 24, 40, 0, time.UTC)
	if *To(x) != x {
		t.Errorf("*To(%v)", x)
	}
	if *ToOrNil(x) != x {
		t.Errorf("*ToOrNil(%v)", x)
	}
	if Get(&x) != x {
		t.Errorf("Get(%v)", &x)
	}
}
