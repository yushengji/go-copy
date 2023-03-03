package gocp

import (
	"reflect"
)

// TypePlugin kind value cp
// assign values for different kinds, like ptr, struct, array, slice val(like uint or int and
// more simple type). support for custom kind conversions then add plugin to typePlugin or use
// RegisterTypePlugin.
type TypePlugin interface {
	// Check this object support cp
	Check(src *ReflectEntity) bool

	// Cp copy value
	Cp(src, dst *ReflectEntity)

	// Kd return this plugin support Kind
	Kd() []reflect.Kind
}

var typePlugins = []TypePlugin{
	&ptrCopier{},
	&valCopier{},
	&structCopier{},
	&arrayCopier{},
}

// RegisterTypePlugin register user type plugins
func RegisterTypePlugin(tcs ...TypePlugin) {
	typePlugins = append(typePlugins, tcs...)
}

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
	for _, c := range typePlugins {
		for _, kd := range c.Kd() {
			if src.tpe().Kind() != kd {
				continue
			}
		}

		if c.Check(src) {
			c.Cp(src, dst)
			return
		}
	}
}
