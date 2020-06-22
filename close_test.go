package safe

import (
	"fmt"
	"io"
	"testing"
)

type stub struct {
	closer func() error
}

func TestClose(t *testing.T) {
	Printf = func(format string, v ...interface{}) {}
	tests := []struct {
		name   string
		closer func() io.Closer
	}{
		{
			name: "nil",
			closer: func() io.Closer {
				return nil
			},
		},
		{
			name: "success",
			closer: func() io.Closer {
				return &stub{func() error {
					return nil
				}}
			},
		},
		{
			name: "error",
			closer: func() io.Closer {
				return &stub{func() error {
					return fmt.Errorf("error")
				}}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Close(test.closer(), test.name)
		})
	}
}

func (s *stub) Close() error {
	return s.closer()
}
