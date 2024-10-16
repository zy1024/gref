package copyValue

import (
	"github.com/zy1024/gref/utils"
	"reflect"
)

func SliceValue(srcElem, dstElem reflect.Value) error {

	// set dstElem to a new slice
	newDst := reflect.MakeSlice(dstElem.Type(), srcElem.Len(), srcElem.Len())
	dstElem.Set(newDst)

	// copy each element
	for i := 0; i < srcElem.Len(); i++ {
		srcElemData := srcElem.Index(i)
		dstElemData := dstElem.Index(i)

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
