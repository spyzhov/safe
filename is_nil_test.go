package safe

import (
	"fmt"
	"testing"
)

func TestIsNil(t *testing.T) {
	var p *tp
	var v tp
	tests := []struct {
		name  string
		value fmt.Stringer
		nil   bool
	}{
		{
			name:  "nil",
			value: nil,
			nil:   true,
		},
		{
			name:  "pointer",
			value: p,
			nil:   true,
		},
		{
			name:  "value",
			value: &v,
			nil:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if IsNil(test.value) != test.nil {
				t.Error("wrong")
			}
		})
	}
}
