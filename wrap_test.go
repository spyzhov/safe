package safe

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrap(t *testing.T) {
	simple := fmt.Errorf("error")
	type args struct {
		err   error
		scope string
	}
	tests := []struct {
		name   string
		args   args
		result error
	}{
		{
			name:   "nil",
			args:   args{nil, ""},
			result: nil,
		},
		{
			name:   "nil/scope",
			args:   args{nil, "scope"},
			result: nil,
		},
		{
			name:   "error/scope",
			args:   args{simple, "scope"},
			result: simple,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrap(tt.args.err, tt.args.scope); err != tt.result && !errors.Is(err, tt.result) {
				t.Errorf("Wrap() error = %v, wanted %v", err, tt.result)
			}
		})
	}
}
