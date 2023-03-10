package gocp

import (
	"reflect"
)

type sliceCopier struct{}

func (s sliceCopier) Check(src *ReflectEntity) bool {
	return src.tpe().Kind() == reflect.Slice
}

func (s sliceCopier) Cp(src, dst *ReflectEntity) {
	s.doCp(dst.elemTpe().Elem(), src.elemVal(), dst.val().Elem())
}

func (s sliceCopier) Kd() []reflect.Kind {
	return []reflect.Kind{
		reflect.Slice,
	}
}

// doCp fill type fields
func (s sliceCopier) doCp(dstT reflect.Type, srcV, dstV reflect.Value) {
	for i := 0; i < srcV.Len(); i++ {
		item := srcV.Index(i)
		appendItem := reflect.New(dstT)

		// cp
		doCp(&ReflectEntity{t: item.Type(), v: item},
			&ReflectEntity{t: appendItem.Type(), v: appendItem})

		dstV.Set(reflect.Append(dstV, appendItem.Elem()))
	}
}
