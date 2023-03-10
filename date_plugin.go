package gocp

import (
	"reflect"
	"time"
)

// DateStringPlugin Suitable for date and string conversion.
// When the string field of time.Time/*time.Time to be copied to
// the target structure appears in the structure, using this plugin
// will automatically convert.
// For example if we have struct A:
//
//	type A struct {
//			A time.Time
//	}
//
// and we have struct B:
//
//	type B struct {
//			B string
//	}
//
// If this plug-in has been registered, the format value will be automatically
// assigned to the B structure variable when copying.
var DateStringPlugin = &dateStringPlugin{}

type dateStringPlugin struct{}

func (d dateStringPlugin) Check(src, dst reflect.StructField) bool {
	sTypeName := src.Type.Name()
	if src.Type.Kind() == reflect.Ptr {
		sTypeName = src.Type.Elem().Name()
	}

	dTypeName := dst.Type.Name()
	if dst.Type.Kind() == reflect.Ptr {
		dTypeName = dst.Type.Elem().Name()
	}

	return src.Type != dst.Type && (sTypeName == "Time" || src.Type.Kind() == reflect.String) &&
		(dTypeName == "Time" || dst.Type.Kind() == reflect.String)
}

func (d dateStringPlugin) To(srcT, _ reflect.StructField, srcV, dstV reflect.Value) reflect.Value {
	isDstStr := dstV.Type().Kind() == reflect.String

	var srcObj time.Time
	if srcT.Type.Kind() == reflect.String {
		var err error
		srcObj, err = time.ParseInLocation("2006-01-02", srcV.String(), time.Local)
		if err != nil {
			panic(err)
		}
	} else {
		var ok bool
		srcObj, ok = srcV.Interface().(time.Time)
		if !ok {
			srcObj = *srcV.Interface().(*time.Time)
		}
	}

	if isDstStr {
		return reflect.ValueOf(srcObj.Format("2006-01-02"))
	}

	return reflect.ValueOf(srcObj)
}
