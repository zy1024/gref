package copyValue

import (
	"github.com/zy1024/gref/utils"
	"reflect"
)

func SliceValue(srcSlice, dstSlice reflect.Value) error {

	// if srcSlice is nil, not need to copy
	if utils.IsZero(srcSlice) {
		return nil
	}

	// Check whether src and dst can be assigned to each other
	if !utils.CanBeCopy(srcSlice.Type().Elem(), dstSlice.Type().Elem()) {
		return nil
	}

	// set dstSlice to a new slice
	newDst := reflect.MakeSlice(dstSlice.Type(), srcSlice.Len(), srcSlice.Len())
	dstSlice.Set(newDst)

	// copy each element
	for i := 0; i < srcSlice.Len(); i++ {
		srcSliceIndex := srcSlice.Index(i)
		dstSliceIndex := dstSlice.Index(i)

		switch {
		case srcSliceIndex.Type() == dstSliceIndex.Type():
			dstSliceIndex.Set(srcSliceIndex)

		case utils.IsBasicType(srcSliceIndex.Kind(), srcSliceIndex.Kind()):
			err := BasicValue(srcSliceIndex, dstSliceIndex)
			if err != nil {
				return err
			}

		case srcSliceIndex.Kind() == reflect.Struct && dstSliceIndex.Kind() == reflect.Struct:
			err := StructValue(srcSliceIndex, dstSliceIndex)
			if err != nil {
				return err
			}

		case srcSliceIndex.Kind() == reflect.Slice && dstSliceIndex.Kind() == reflect.Slice:
			err := SliceValue(srcSliceIndex, dstSliceIndex)
			if err != nil {
				return err
			}

		case srcSliceIndex.Kind() == reflect.Ptr && dstSliceIndex.Kind() == reflect.Ptr:
			err := PointerValue(srcSliceIndex, dstSliceIndex)
			if err != nil {
				return err
			}

		}

	}

	return nil
}
