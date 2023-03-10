package gocp

import "reflect"

type valCopier struct{}

func (v valCopier) Check(src *ReflectEntity) bool {
	k := src.tpe().Kind()
	return k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 ||
		k == reflect.Int64 || k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 ||
		k == reflect.Uint32 || k == reflect.Uint64 || k == reflect.Float64 || k == reflect.Float32 ||
		k == reflect.Bool || k == reflect.Complex64 || k == reflect.Complex128 || k == reflect.String ||
		k == reflect.Map
}

func (v valCopier) Cp(src, dst *ReflectEntity) {
	v.doCp(src.elemVal(), dst.elemVal())
}

func (v valCopier) doCp(src, dst reflect.Value) {
	dst.Set(src)
}

func (v valCopier) Kd() []reflect.Kind {
	return []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float64,
		reflect.Float32,
		reflect.Bool,
		reflect.Complex64,
		reflect.Complex128,
		reflect.String,
		reflect.Map,
	}
}
