package gocp

import (
	"reflect"
)

// NamePlugin Suitable for different name struct field value copy.
// For example, we have struct A:
//
//	type A struct {
//			A int
//	}
//
// and we have struct B:
//
//	type B struct {
//			B int
//	}
//
// We copied struct A to struct B, which is obviously invalid. The reason is that the A
// attribute is A and the B attribute is named B. You can add tag like this:
//
//	type B struct {
//			B int `gocp-name:"A"`
//	}
//
// Then the framework can auto find the field named A.
var NamePlugin = &namePlugin{}

type namePlugin struct{}

func (n namePlugin) Check(dstF reflect.StructField) bool {
	return dstF.Tag.Get("gocp-name") != ""
}

func (n namePlugin) Match(src reflect.Type, dstF reflect.StructField) (reflect.StructField, bool) {
	return src.FieldByName(dstF.Tag.Get("gocp-name"))
}

func (n namePlugin) Verify(srcF, dstF reflect.StructField) bool {
	return srcF.Type == dstF.Type
}

func (n namePlugin) Transform(srcV, _ reflect.Value) reflect.Value {
	return srcV
}

func (n namePlugin) Order() int {
	return 1
}
