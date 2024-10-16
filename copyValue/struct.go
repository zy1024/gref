package copyValue

import (
	"github.com/zy1024/gref/utils"
	"reflect"
)

func StructValue(srcStruct, dstStruct reflect.Value) error {

	for i := 0; i < srcStruct.NumField(); i++ {
		srcField := srcStruct.Field(i)
		srcFieldName := srcStruct.Type().Field(i).Name

		// get the field of dst with the same name of src
		dstField := dstStruct.FieldByName(srcFieldName)
		// check if the dst field exists and can be set
		if dstField.IsValid() && dstField.CanSet() {
			// check if the dst field is zero value, if so, set it to the src field
			if !utils.IsZero(dstField) && dstField.Kind() != reflect.Ptr {
				continue
			}

			// base on the type of dst field, copy the value from src to dst
			switch {
			case dstField.Type() == srcField.Type():
				dstField.Set(srcField)

			case dstField.Kind() == reflect.Slice && srcField.Kind() == reflect.Slice:
				// Handle slice fields
				err := SliceValue(srcField, dstField)
				if err != nil {
					return err
				}

			case dstField.Kind() == reflect.Struct && srcField.Kind() == reflect.Struct:
				// Handle nested struct fields
				err := StructValue(srcField, dstField)
				if err != nil {
					return err
				}

			case dstField.Kind() == reflect.Ptr && srcField.Kind() == reflect.Ptr:
				if !utils.IsZero(dstField.Elem()) {
					continue
				}
				// Handle pointer fields
				err := PointerValue(srcField, dstField)
				if err != nil {
					return err
				}

			case utils.IsBasicType(srcField.Kind()) && utils.IsBasicType(dstField.Kind()):
				// Handle basic type slices
				err := BasicValue(srcField, dstField)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
