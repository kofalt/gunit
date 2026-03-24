package should

import (
	"errors"
	"reflect"
)

// BeZeroValue verifies that actual is the zero value for its type.
func BeZeroValue(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}
	if actual == nil {
		return nil
	}
	if reflect.ValueOf(actual).IsZero() {
		return nil
	}
	return failure("got %#v, want zero value", actual)
}

// BeZeroValue negated!
func (negated) BeZeroValue(actual any, expected ...any) error {
	err := BeZeroValue(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}
	if err != nil {
		return err
	}
	return failure("got zero value, want non-zero value")
}
