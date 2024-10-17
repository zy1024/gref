package utils

import (
	"errors"
	"reflect"
)

// CheckSrcAndDst checks if the given src and dst are pointers and returns the values that they are pointing to.
func CheckSrcAndDst(src interface{}, dst interface{}) (reflect.Value, reflect.Value, error) {
	// check if src and dst are not nil
	if src == nil {
		return reflect.Value{}, reflect.Value{}, errors.New("src is nil")
	}
	if dst == nil {
		return reflect.Value{}, reflect.Value{}, errors.New("dst is nil")
	}

	// check if src and dst are pointer
	if reflect.ValueOf(src).Kind() != reflect.Ptr {
		return reflect.Value{}, reflect.Value{}, errors.New("src must be pointer")
	}
	if reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return reflect.Value{}, reflect.Value{}, errors.New("dst must be pointer")
	}

	srcElem := reflect.ValueOf(src).Elem()
	dstElem := reflect.ValueOf(dst).Elem()

	return srcElem, dstElem, nil
}
