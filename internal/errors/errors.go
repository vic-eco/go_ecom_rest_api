package errors

import "errors"

var ErrNotFound = errors.New("resource not found")
var ErrNoStock = errors.New("resource out of stock")

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
func IsNoStock(err error) bool {
	return errors.Is(err, ErrNoStock)
}
