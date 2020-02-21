package safe

import (
	"fmt"
	"testing"
)

func TestMust(t *testing.T) {
	tests := []struct {
		name  string
		arg   interface{}
		err   error
		panic bool
	}{
		{
			name:  "nil",
			arg:   nil,
			err:   nil,
			panic: false,
		},
		{
			name:  "err",
			arg:   nil,
			err:   fmt.Errorf("error"),
			panic: true,
		},
		{
			name:  "value",
			arg:   123,
			err:   nil,
			panic: false,
		},
		{
			name:  "value/err",
			arg:   123,
			err:   fmt.Errorf("error"),
			panic: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if test.panic && r == nil {
					t.Error("panic not found")
				} else if !test.panic && r != nil {
					t.Error("panic was found")
				}
			}()

			if Must(test.arg, test.err) != test.arg {
				t.Error("result was changed")
			}
		})
	}
}
