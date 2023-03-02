package gocp

import "reflect"

// setStruct set struct fields
func setStruct(src, dst *entity) {
	setStructField(src.elemTpe(), dst.elemTpe(), src.elemVal(), dst.elemVal())
}

// setStructField fill type fields
func setStructField(srcT, dstT reflect.Type, srcV, dstV reflect.Value) {
	num := srcT.NumField()
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
			if srcField.Type.Kind() == reflect.Struct {
				setStructField(srcField.Type, dstField.Type, srcV.Field(i), dstV.Field(i))
				continue
			}
			continue
		}

		dstV.Field(i).Set(srcV.Field(i))
	}
}
