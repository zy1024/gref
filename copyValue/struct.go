package copyValue

import (
	"github.com/zy1024/gref/utils"
	"reflect"
)

func StructValue(srcStruct, dstStruct reflect.Value) error {

	// check if the src struct is zero value
	if utils.IsZero(srcStruct) {
		return nil
	}

	// copy the value of src struct to dst struct
	for i := 0; i < dstStruct.NumField(); i++ {
		dstField := dstStruct.Field(i)
		dstFieldName := dstStruct.Type().Field(i).Name

		// get the field of src with the same name of dst
		srcField := srcStruct.FieldByName(dstFieldName)
		// check if the src field exists and can be set
		if srcField.IsValid() && dstField.CanSet() {

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
