package gocp

import (
	"reflect"
	"time"
)

// TimeStringPlugin Suitable for date and string conversion.
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
//			B string `gocp:"time"`
//	}
//
// If this plug-in has been registered, the format value will be automatically
// assigned to the B structure variable when copying.
var TimeStringPlugin = &timeStringPlugin{}

type timeStringPlugin struct{}

func (d timeStringPlugin) Check(dstF reflect.StructField) bool {
	return dstF.Tag.Get("gocp") == "time"
}

func (d timeStringPlugin) Match(srcT reflect.Type, dstF reflect.StructField) (reflect.StructField, bool) {
	return srcT.FieldByName(dstF.Name)
}

func (d timeStringPlugin) Verify(srcF, dstF reflect.StructField) bool {
	sT := srcF.Type
	if srcF.Type.Kind() == reflect.Ptr {
		sT = srcF.Type.Elem()
	}

	dT := dstF.Type
	if dstF.Type.Kind() == reflect.Ptr {
		dT = dstF.Type.Elem()
	}

	return (sT.Kind() == reflect.String || sT.Name() == "Time") &&
		(dT.Kind() == reflect.String || dT.Name() == "Time")
}

func (d timeStringPlugin) Transform(srcV, dstV reflect.Value) reflect.Value {
	var srcObj time.Time
	if srcV.Type().Kind() == reflect.String {
		var err error
		srcObj, err = time.ParseInLocation("2006-01-02 15:04:05", srcV.String(), time.Local)
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

	if dstV.Type().Kind() == reflect.String {
		return reflect.ValueOf(srcObj.Format("2006-01-02 15:04:05"))
	}

	return reflect.ValueOf(srcObj)
}

func (d timeStringPlugin) Order() int {
	return 10
}

func (d timeStringPlugin) Find(src reflect.Type, dstF reflect.StructField) (reflect.StructField, bool) {
	if dstF.Tag.Get("gocp") != "time" {
		return reflect.StructField{}, false
	}

	return src.FieldByName(dstF.Name)
}
