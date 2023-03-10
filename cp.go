package gocp

import (
	"reflect"
)

// Cp copy value from src to dst
func Cp(src, dst interface{}) {
	if src == nil || dst == nil {
		return
	}

	srcEntity := &ReflectEntity{o: src, t: reflect.TypeOf(src)}
	dstEntity := &ReflectEntity{o: dst, t: reflect.TypeOf(dst)}

	// dst must ptr
	if dstEntity.tpe().Kind() != reflect.Ptr {
		return
	}

	// match kind
	if srcEntity.elemTpe().Kind() != dstEntity.elemTpe().Kind() {
		return
	}

	doCp(srcEntity, dstEntity)
}

func doCp(src, dst *ReflectEntity) {
	// type plugin Cp
	plugin, ok := typePlugins[src.tpe().Kind()]
	if ok && plugin.Check(src) {
		plugin.Cp(src, dst)
		return
	}
}
