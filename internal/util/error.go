package util

import (
	"errors"
	"fmt"
)

func HasError(err error) bool {
	if err != nil {
		Logger("has error:", err)

	}
	return err != nil
}

func Logger(any ...any) {
	fmt.Println("\n" + fmt.Sprint(any...) + "\n")
}

// Recover from panic and return error
func Recover() error {
	if err := recover(); err != nil {
		return errors.New(fmt.Sprint(err))
	}
	return nil
}
