package gocp

import "reflect"

func stru(src, dst *entity) {
	setStructField(src.elemTpe(), dst.elemTpe(), src.elemVal(), dst.elemVal())
}

func setStructField(srcT, dstT reflect.Type, srcV, dstV reflect.Value) {
	for _, j := range filterFields(srcT, dstT) {
		dstV.Field(j).Set(srcV.Field(j))
	}
}
