package gref

import (
	"errors"
	"github.com/zy1024/gref/copyValue"
	"github.com/zy1024/gref/utils"
	"reflect"
)

// Copy Assignment of two different types, from src to dst
// Notice: src and dst must be pointers.
func Copy(src, dst interface{}) error {
	srcElem, dstElem, err := utils.CheckSrcAndDst(src, dst)
	if err != nil {
		return errors.New("Copy: " + err.Error())
	}

	switch {
	case srcElem.Kind() == reflect.Struct && dstElem.Kind() == reflect.Struct:
		err = copyValue.StructValue(srcElem, dstElem)
		if err != nil {
			return errors.New("gref.Copy: " + err.Error())
		}

	case srcElem.Kind() == reflect.Slice && dstElem.Kind() == reflect.Slice:
		err = copyValue.SliceValue(srcElem, dstElem)
		if err != nil {
			return errors.New("gref.Copy: " + err.Error())
		}

	case srcElem.Kind() == reflect.Ptr && dstElem.Kind() == reflect.Ptr:
		err = copyValue.PointerValue(srcElem, dstElem)
		if err != nil {
			return errors.New("gref.Copy: " + err.Error())
		}

	case utils.IsBasicType(srcElem.Kind()) && utils.IsBasicType(dstElem.Kind()):
		err = copyValue.BasicValue(srcElem, dstElem)
		if err != nil {
			return errors.New("gref.Copy: " + err.Error())
		}

	}

	return nil
}
