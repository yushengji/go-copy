package gocp

import (
	"reflect"
)

// Cp copy value from src to dst
func Cp(src, dst interface{}) {
	if src == nil || dst == nil {
		return
	}

	cp(&entity{o: src, t: reflect.TypeOf(src)},
		&entity{o: dst, t: reflect.TypeOf(dst)})
}

func cp(src, dst *entity) {
	// dst must ptr
	if dst.tpe().Kind() != reflect.Ptr {
		return
	}

	// match kind
	if src.elemTpe().Kind() != dst.elemTpe().Kind() {
		return
	}

	// src ptr
	if src.tpe().Kind() == reflect.Ptr {
		ptr(src, dst)
		return
	}

	// src basic
	if allValType(src.elemTpe().Kind(), dst.elemTpe().Kind()) {
		setVal(src, dst)
		return
	}

	// src struct
	stru(src, dst)
}
