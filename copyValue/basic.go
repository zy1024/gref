package copyValue

import (
	"fmt"
	"reflect"
	"strconv"
)

// BasicValue copies a basic value from src to dst, converting types if necessary.
func BasicValue(src, dst reflect.Value) error {
	switch src.Kind() {
	case reflect.Bool:
		switch dst.Kind() {
		case reflect.Bool:
			dst.SetBool(src.Bool())
		case reflect.String:
			dst.SetString(strconv.FormatBool(src.Bool()))
		default:
			return fmt.Errorf("unsupported type conversion from bool to %s", dst.Kind())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch dst.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			dst.SetInt(src.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			dst.SetUint(uint64(src.Int()))
		case reflect.Float32, reflect.Float64:
			dst.SetFloat(float64(src.Int()))
		case reflect.String:
			dst.SetString(strconv.FormatInt(src.Int(), 10))
		default:
			return fmt.Errorf("unsupported type conversion from int to %s", dst.Kind())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		switch dst.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			dst.SetInt(int64(src.Uint()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			dst.SetUint(src.Uint())
		case reflect.Float32, reflect.Float64:
			dst.SetFloat(float64(src.Uint()))
		case reflect.String:
			dst.SetString(strconv.FormatUint(src.Uint(), 10))
		default:
			return fmt.Errorf("unsupported type conversion from uint to %s", dst.Kind())
		}
	case reflect.Float32, reflect.Float64:
		switch dst.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			dst.SetInt(int64(src.Float()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			dst.SetUint(uint64(src.Float()))
		case reflect.Float32, reflect.Float64:
			dst.SetFloat(src.Float())
		case reflect.String:
			dst.SetString(strconv.FormatFloat(src.Float(), 'f', -1, 64))
		default:
			return fmt.Errorf("unsupported type conversion from float to %s", dst.Kind())
		}
	case reflect.Complex64, reflect.Complex128:
		switch dst.Kind() {
		case reflect.Complex64, reflect.Complex128:
			dst.SetComplex(src.Complex())
		case reflect.String:
			dst.SetString(strconv.FormatComplex(src.Complex(), 'f', -1, 128))
		default:
			return fmt.Errorf("unsupported type conversion from complex to %s", dst.Kind())
		}
	case reflect.String:
		switch dst.Kind() {
		case reflect.Bool:
			val, err := strconv.ParseBool(src.String())
			if err != nil {
				return err
			}
			dst.SetBool(val)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, err := strconv.ParseInt(src.String(), 10, 64)
			if err != nil {
				return err
			}
			dst.SetInt(val)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			val, err := strconv.ParseUint(src.String(), 10, 64)
			if err != nil {
				return err
			}
			dst.SetUint(val)
		case reflect.Float32, reflect.Float64:
			val, err := strconv.ParseFloat(src.String(), 64)
			if err != nil {
				return err
			}
			dst.SetFloat(val)
		case reflect.Complex64, reflect.Complex128:
			val, err := strconv.ParseComplex(src.String(), 128)
			if err != nil {
				return err
			}
			dst.SetComplex(val)
		case reflect.String:
			dst.SetString(src.String())
		default:
			return fmt.Errorf("unsupported type conversion from string to %s", dst.Kind())
		}
	default:
		return fmt.Errorf("unsupported type %s", src.Kind())
	}

	return nil
}
