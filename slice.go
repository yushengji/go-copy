package gocp

import (
	"reflect"
)

type sliceCopier struct{}

func (a sliceCopier) Check(src *ReflectEntity) bool {
	return src.tpe().Kind() == reflect.Slice
}

func (a sliceCopier) Cp(src, dst *ReflectEntity) {
	srcV := src.elemVal()
	dstV := dst.val().Elem()
	for i := 0; i < srcV.Len(); i++ {
		item := srcV.Index(i)
		appendItem := reflect.New(dst.elemTpe().Elem())

		// cp
		doCp(&ReflectEntity{t: item.Type(), v: item},
			&ReflectEntity{t: appendItem.Type(), v: appendItem})

		dstV.Set(reflect.Append(dstV, appendItem.Elem()))
	}
}

func (a sliceCopier) Kd() []reflect.Kind {
	return []reflect.Kind{
		reflect.Slice,
	}
}
