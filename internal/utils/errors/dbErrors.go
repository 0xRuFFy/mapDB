package errors

import "errors"

func InvalidItemValueTypeError() error {
	return errors.New("invalid item value type")
}
