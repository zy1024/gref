package utils

import "reflect"

// CanBeCopy check if srcValue can be copied to dstValue
func CanBeCopy(srcValue, dstValue reflect.Type) bool {

	switch {
	case IsBasicType(srcValue.Kind(), dstValue.Kind()):
		return true

	case srcValue.Kind() == reflect.Ptr && dstValue.Kind() == reflect.Ptr:
		return CanBeCopy(srcValue.Elem(), dstValue.Elem())

	case srcValue.Kind() == reflect.Slice && dstValue.Kind() == reflect.Slice:
		return CanBeCopy(srcValue.Elem(), dstValue.Elem())

	case srcValue.Kind() == reflect.Struct && dstValue.Kind() == reflect.Struct:
		return true
	}

	return false
}
