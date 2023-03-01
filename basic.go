package gocp

import "reflect"

// setVal value cp
func setVal(src, dst *entity) {
	dst.elemVal().Set(src.elemVal())
}

func allValType(ks ...reflect.Kind) bool {
	for i := range ks {
		if ks[i] == reflect.Struct {
			return false
		}
	}

	return true
}
