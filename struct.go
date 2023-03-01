package gocp

import "reflect"

func stru(src, dst *entity) {
	setStructField(src.elemTpe(), dst.elemTpe(), src.elemVal(), dst.elemVal())
}

func setStructField(srcT, dstT reflect.Type, srcV, dstV reflect.Value) {
	num := srcV.NumField()
	for i := 0; i < num; i++ {
		srcField := srcT.Field(i)
		dstField := dstT.Field(i)

		if !srcField.IsExported() || !dstField.IsExported() {
			continue
		}

		if srcField.Name != dstField.Name {
			continue
		}

		if srcField.Type != dstField.Type {
			continue
		}

		dstV.Field(i).Set(srcV.Field(i))
	}
}
