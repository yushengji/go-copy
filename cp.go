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
