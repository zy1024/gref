package utils

import "reflect"

// IsZero checks if a value is zero
func IsZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr, reflect.Func, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.Struct:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	default:
		return false
	}
}
