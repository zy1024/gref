package copyValue

import (
	"github.com/zy1024/gref/utils"
	"reflect"
)

func SliceValue(srcSlice, dstSlice reflect.Value) error {

	// Check whether src and dst can be assigned to each other
	if !utils.CanBeCopy(srcSlice.Type().Elem(), dstSlice.Type().Elem()) {
		return nil
	}

	// set dstSlice to a new slice
	newDst := reflect.MakeSlice(dstSlice.Type(), srcSlice.Len(), srcSlice.Len())
	dstSlice.Set(newDst)

	// copy each element
	for i := 0; i < srcSlice.Len(); i++ {
		srcElemData := srcSlice.Index(i)
		dstElemData := dstSlice.Index(i)

		switch {
		case utils.IsBasicType(srcElemData.Kind()) && utils.IsBasicType(srcElemData.Kind()):
			err := BasicValue(srcElemData, dstElemData)
			if err != nil {
				return err
			}

		case srcElemData.Kind() == reflect.Struct && dstElemData.Kind() == reflect.Struct:
			err := StructValue(srcElemData, dstElemData)
			if err != nil {
				return err
			}

		case srcElemData.Kind() == reflect.Slice && dstElemData.Kind() == reflect.Slice:
			err := SliceValue(srcElemData, dstElemData)
			if err != nil {
				return err
			}

		case srcElemData.Kind() == reflect.Ptr && dstElemData.Kind() == reflect.Ptr:
			err := PointerValue(srcElemData, dstElemData)
			if err != nil {
				return err
			}

		}

	}

	return nil
}
