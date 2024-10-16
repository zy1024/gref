package copyValue

import (
	"github.com/zy1024/gref/utils"
	"reflect"
)

// PointerValue copy src pointer value to dst
func PointerValue(srcPtr, dstPtr reflect.Value) error {

	// if src pointers is nil, return nil
	if srcPtr.IsNil() {
		return nil
	}

	// if dst pointers is nil, set dst pointers to zero value
	if dstPtr.IsNil() {
		dstPtr.Set(reflect.New(dstPtr.Type().Elem()))
	}

	srcElem := srcPtr.Elem()
	dstElem := dstPtr.Elem()

	switch {
	case srcElem.Kind() == reflect.Struct && dstElem.Kind() == reflect.Struct:
		err := StructValue(srcElem, dstElem)
		if err != nil {
			return err
		}

	case srcElem.Kind() == reflect.Slice && dstElem.Kind() == reflect.Slice:
		err := SliceValue(srcElem, dstElem)
		if err != nil {
			return err
		}

	case srcElem.Kind() == reflect.Ptr && dstElem.Kind() == reflect.Ptr:
		err := PointerValue(srcElem, dstElem)
		if err != nil {
			return err
		}

	case utils.IsBasicType(srcElem.Kind()) && utils.IsBasicType(dstElem.Kind()):
		err := BasicValue(srcElem, dstElem)
		if err != nil {
			return err
		}

	}

	return nil
}
