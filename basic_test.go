package gocp

import (
	"reflect"
	"testing"
)

func TestSetVal(t *testing.T) {
	src := 100
	dst := 0

	setVal(&entity{o: src}, &entity{o: &dst})
	if dst != 100 {
		t.Fatalf("dst should be 100")
		return
	}
}

func TestAllValType(t *testing.T) {
	if allValType(reflect.Struct, reflect.Struct) {
		t.Fatalf("2 struct shouldent val type")
		return
	}

	if allValType(reflect.Array, reflect.Struct) {
		t.Fatalf("array and struct should val type")
		return
	}

	if !allValType(reflect.Array, reflect.Int) {
		t.Fatalf("array and int should val type")
		return
	}
}
