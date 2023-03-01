package gocp

import "reflect"

var matchCache = make(map[*entity]map[*entity]bool)

func cp(src, dst *entity) {
	vMap := matchCache[src]
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
