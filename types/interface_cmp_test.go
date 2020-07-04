package types

import (
	"testing"
)

type MyErr struct {
}

func (e *MyErr) Error() string {
	return "My Error"
}

func getError() error {
	var err *MyErr
	return err
}

func getMyErr() *MyErr {
	var err *MyErr
	return err
}

func TestErrorCompare(t *testing.T) {
	println(getError() == nil)          // false, error (interface) compare with nil
	println(getError().(*MyErr) == nil) // true, assert error (interface) as *MyErr, then compare with nil
	println(getMyErr() == nil)          // true, *MyErr compare with nil
}

func TestInterfaceCompare(t *testing.T) {
	var x, y interface{}
	println(x == y) // true, both nil
	x = int32(1)
	y = int64(1)
	println(x == y) // false, type int32 != int64
	x = int64(1)
	println(x == y) // true, both are int64 (type) and 1 (value)
}
