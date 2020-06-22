package safe

import (
	"io"
)

// Close will close any io.Closer object if it's not nil, any error will be only logged.
func Close(closer io.Closer, scope string) {
	if !IsNil(closer) {
		if err := closer.Close(); !IsNil(err) {
			Printf("error on close '%s': %s", scope, err)
		}
	}
}
