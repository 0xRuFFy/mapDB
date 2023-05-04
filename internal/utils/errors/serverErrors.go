package errors

import "fmt"

func NilListener() error {
	return fmt.Errorf("listener is nil")
}
