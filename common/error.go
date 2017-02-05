package common

import "github.com/pkg/errors"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Error(msg string) error {
	return errors.New(msg)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	panic(Sprintf(format, args...))
}
