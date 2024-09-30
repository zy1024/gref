package gref

import (
	"errors"
	"github.com/zy1024/gref/utils"
	"reflect"
)

// CopyFields copies fields from src to dst.
// It only copies fields with the same name and type.
// Notice : src and dst must be pointers to struct.
func CopyFields(src, dst interface{}) error {

	// check if src and dst are not nil
	if src == nil {
		return errors.New("CopyFields: src is nil")
	}
	if dst == nil {
		return errors.New("CopyFields: dst is nil")
	}

	// check if src and dst are pointers to struct
	err := utils.CheckPointerToStruct(src, dst)
	if err != nil {
		return err
	}

	// get the pointer to struct of src and dst
	srcStruct := reflect.ValueOf(src).Elem()
	dstStruct := reflect.ValueOf(dst).Elem()

	for i := 0; i < srcStruct.NumField(); i++ {
		srcField := srcStruct.Field(i)
		srcFieldName := srcStruct.Type().Field(i).Name

		// get the field of dst with the same name of src
		dstField := dstStruct.FieldByName(srcFieldName)
		// check if the dst field exists, the dst field can be set and the dst field type is the same as the src field type
		if dstField.IsValid() && dstField.CanSet() && dstField.Type() == srcField.Type() {
			// check if the dst field is zero value, if so, set it to the src field
			if utils.IsZero(dstField) {
				dstField.Set(srcField)
			}
		}
	}

	return nil
}
