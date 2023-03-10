package gocp

import "reflect"

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

var typePlugins = map[reflect.Kind]TypePlugin{
	reflect.Ptr:    &ptrCopier{},
	reflect.Struct: &structCopier{},
	reflect.Slice:  &sliceCopier{},
	reflect.Array:  &arrayCopier{},
}

func init() {
	RegisterTypePlugin(&valCopier{})
}
