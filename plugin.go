package gocp

import (
	"reflect"
	"sort"
)

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
	// Check if the target struct field can be processed.
	// Field can be processed, return true, can't return false.
	Check(dstF reflect.StructField) bool

	// Match source struct field against target struct field.
	// If match success, return the field and true.
	// On the contrary, return zero value and false.
	Match(srcT reflect.Type, dstF reflect.StructField) (reflect.StructField, bool)

	// Verify that the source struct fields match the target struct fields.
	Verify(srcF, dstF reflect.StructField) bool

	// Transform the source value by the target value.
	// Return transformed value.
	Transform(srcV, dstV reflect.Value) reflect.Value

	// Order is plugin order, the smaller the front.
	// Order will matter when processing struct fields.
	Order() int
}

var typePlugins = map[reflect.Kind]TypePlugin{
	reflect.Ptr:    &ptrCopier{},
	reflect.Struct: &structCopier{},
	reflect.Slice:  &sliceCopier{},
	reflect.Array:  &arrayCopier{},
}

var fieldPlugins []FieldPlugin

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
loop:
	for _, tp := range tps {
		for _, exist := range fieldPlugins {
			if exist == tp {
				continue loop
			}
		}

		fieldPlugins = append(fieldPlugins, tp)
	}

	sort.Slice(fieldPlugins, func(i, j int) bool {
		return fieldPlugins[i].Order() > fieldPlugins[j].Order()
	})
}

func init() {
	RegisterTypePlugin(&valCopier{})
}
