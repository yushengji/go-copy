package gocp

import "reflect"

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

func (n namePlugin) Check(src, dst reflect.StructField) bool {
	return src.Type == dst.Type && dst.Tag.Get("gocp-name") == src.Name
}

func (n namePlugin) To(_, _ reflect.StructField, srcV, _ reflect.Value) reflect.Value {
	return srcV
}
