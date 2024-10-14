package gref

import (
	"errors"
	"fmt"
	"github.com/zy1024/gref/utils"
	"reflect"
)

// CopyStructFields copies fields from src to dst.
// It only copies fields with the same name and type.
// Notice : src and dst must be pointers to struct.
func CopyStructFields(src, dst interface{}) error {
	// check if src and dst are not nil
	if src == nil {
		return errors.New("CopyStructFields: src is nil")
	}
	if dst == nil {
		return errors.New("CopyStructFields: dst is nil")
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
		// check if the dst field exists and can be set
		if dstField.IsValid() && dstField.CanSet() {
			// check if the dst field is not zero value, if so, skip it
			if !utils.IsZero(dstField) {
				continue
			}

			// base on the type of dst field, copy the value from src to dst
			switch {
			case dstField.Type() == srcField.Type():
				// check if the dst field is zero value, if so, set it to the src field
				dstField.Set(srcField)

			case dstField.Kind() == reflect.Slice && srcField.Kind() == reflect.Slice:
				// Handle slice fields
				switch {
				case dstField.Type().Elem() == srcField.Type().Elem():
					dstField.Set(srcField)

				case dstField.Type().Elem().Kind() == reflect.Struct && srcField.Type().Elem().Kind() == reflect.Struct:
					// Handle nested struct slices
					for j := 0; j < srcField.Len(); j++ {
						srcElem := srcField.Index(j)
						dstElem := reflect.New(dstField.Type().Elem()).Elem()
						if err := CopyStructFields(srcElem.Addr().Interface(), dstElem.Addr().Interface()); err != nil {
							return fmt.Errorf("CopyStructFields: error copying basic type field %s: %w", srcFieldName, err)
						}
						dstField.Set(reflect.Append(dstField, dstElem))
					}

				case dstField.Type().Elem().Kind() == reflect.Ptr && srcField.Type().Elem().Kind() == reflect.Ptr:
					// Handle nested pointer slices
					for j := 0; j < srcField.Len(); j++ {
						srcElem := srcField.Index(j)
						if srcElem.IsNil() {
							dstField.Set(reflect.Append(dstField, reflect.Zero(dstField.Type().Elem())))
						} else {
							dstElem := reflect.New(dstField.Type().Elem().Elem()).Elem()
							if err := CopyStructFields(srcElem.Elem().Addr().Interface(), dstElem.Addr().Interface()); err != nil {
								return fmt.Errorf("CopyStructFields: error copying basic type field %s: %w", srcFieldName, err)
							}
							dstField.Set(reflect.Append(dstField, dstElem.Addr()))
						}
					}

				case utils.IsBasicType(srcField.Type().Elem().Kind()) && utils.IsBasicType(dstField.Type().Elem().Kind()):
					// Handle basic type slices
					if err := CopyBasicValue(srcField.Addr().Interface(), dstField.Addr().Interface()); err != nil {
						return fmt.Errorf("CopyStructFields: error copying basic type field %s: %w", srcFieldName, err)
					}

				}

			case dstField.Kind() == reflect.Struct && srcField.Kind() == reflect.Struct:
				// Handle nested struct fields
				if err := CopyStructFields(srcField.Addr().Interface(), dstField.Addr().Interface()); err != nil {
					return fmt.Errorf("CopyStructFields: error copying basic type field %s: %w", srcFieldName, err)
				}

			case dstField.Kind() == reflect.Ptr && srcField.Kind() == reflect.Ptr:
				// Handle pointer fields
				switch {
				case dstField.Type().Elem() == srcField.Type().Elem():
					dstField.Set(srcField)

				case dstField.Type().Elem().Kind() == reflect.Struct && srcField.Type().Elem().Kind() == reflect.Struct:
					// Handle nested struct pointers
					if srcField.IsNil() {
						dstField.Set(reflect.Zero(dstField.Type()))
					} else {
						dstElem := reflect.New(dstField.Type().Elem()).Elem()
						if err := CopyStructFields(srcField.Elem().Addr().Interface(), dstElem.Addr().Interface()); err != nil {
							return fmt.Errorf("CopyStructFields: error copying basic type field %s: %w", srcFieldName, err)
						}
						dstField.Set(dstElem.Addr())
					}

				case utils.IsBasicType(srcField.Type().Elem().Kind()) && utils.IsBasicType(dstField.Type().Elem().Kind()):
					// Handle basic type slices
					if err := CopyBasicValue(srcField.Addr().Interface(), dstField.Addr().Interface()); err != nil {
						return fmt.Errorf("CopyStructFields: error copying basic type field %s: %w", srcFieldName, err)
					}
				}

			case utils.IsBasicType(srcField.Kind()) && utils.IsBasicType(dstField.Kind()):
				// Handle basic type slices
				if err := CopyBasicValue(srcField.Addr().Interface(), dstField.Addr().Interface()); err != nil {
					return fmt.Errorf("CopyStructFields: error copying basic type field %s: %w", srcFieldName, err)
				}
			}
		}
	}

	return nil
}

// CopyBasicValue copies the value from src to dst.
// It supports conversion between different basic types.
// Notice: src and dst must be pointers to basic types.
func CopyBasicValue(src, dst interface{}) error {
	// check if src and dst are not nil
	if src == nil {
		return errors.New("CopyValue: src is nil")
	}
	if dst == nil {
		return errors.New("CopyValue: dst is nil")
	}

	// check if src and dst are pointers
	if err := utils.IsPointer(src, dst); err != nil {
		return errors.New("CopyValue: src and dst must be pointers")
	}

	// check if src and dst are pointers
	srcElem := reflect.ValueOf(src).Elem()
	dstElem := reflect.ValueOf(dst).Elem()

	// check if src and dst are basic types
	if !utils.IsBasicType(srcElem.Kind()) {
		return errors.New("CopyValue: src must be basic type")
	}
	if !utils.IsBasicType(dstElem.Kind()) {
		return errors.New("CopyValue: dst must be basic types")
	}

	// handle basic types
	return utils.CopyBasicValue(srcElem, dstElem)
}
