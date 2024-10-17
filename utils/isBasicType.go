package utils

import "reflect"

// IsBasicType checks if all given kinds are basic types.
func IsBasicType(kinds ...reflect.Kind) bool {
	for _, kind := range kinds {
		switch kind {
		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128,
			reflect.String:
			continue
		default:
			return false
		}
	}
	return true
}
