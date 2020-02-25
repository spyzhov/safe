package safe

import "fmt"

// Wrap will wrap current error with the scope value, or return nil if error wasn't set.
func Wrap(err error, scope string) error {
	if !IsNil(err) {
		return fmt.Errorf("%s: %w", scope, err)
	}
	return nil
}
