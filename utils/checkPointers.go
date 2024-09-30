package utils

import (
	"errors"
	"reflect"
)

// CheckPointerToStruct checks if the given parameters are pointers to structs.
// It returns an error if any of the parameters are not pointers or not pointers to structs.
func CheckPointerToStruct(params ...interface{}) error {
	for _, param := range params {
		pointer := reflect.ValueOf(param)

		// check if the parameter is a pointer
		if pointer.Kind() != reflect.Ptr {
			return errors.New("CheckPointers: parameter must be a pointer")
		}

		// get the value that the pointer is pointing to
		Elem := pointer.Elem()

		// check if the value is a struct
		if Elem.Kind() != reflect.Struct {
			return errors.New("CheckPointers: parameter must be a pointer to a struct")
		}
	}
	return nil
}
