package gocp

import "reflect"

// TypePlugin value cp for different kind
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

// FieldPlugin value cp for struct field
// This plugin is mainly aimed at non-traditional struct field. The main scenarios are:
// different property names, different property types compatible.
// Built-in plugins include name plugins and date plugins.
// FieldPlugin effective after register, you can use RegisterFieldPlugin function register.
// It should be noted that the use of built-in plugins requires registration.
// Once registered, the plugin-processed values are preferred when copying struct properties.
type FieldPlugin interface {
	// Check field should use this plugin.
	Check(src, dst reflect.StructField) bool

	// To convert src value to new value to set.
	To(srcT, dstT reflect.StructField, srcV, dstV reflect.Value) reflect.Value
}

var typePlugins = map[reflect.Kind]TypePlugin{
	reflect.Ptr:    &ptrCopier{},
	reflect.Struct: &structCopier{},
	reflect.Slice:  &sliceCopier{},
	reflect.Array:  &arrayCopier{},
}

var fieldPlugins = make(map[FieldPlugin]struct{})

// RegisterTypePlugin register user type plugins
func RegisterTypePlugin(tps ...TypePlugin) {
	for _, tp := range tps {
		for _, kd := range tp.Kd() {
			typePlugins[kd] = tp
		}
	}
}

// RegisterFieldPlugin register user field plugins
func RegisterFieldPlugin(tps ...FieldPlugin) {
	for i := range tps {
		fieldPlugins[tps[i]] = struct{}{}
	}
}

func init() {
	RegisterTypePlugin(&valCopier{})
}
