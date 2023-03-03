package gocp

import "reflect"

type structCopier struct{}

func (s structCopier) Check(src *ReflectEntity) bool {
	// kind dst == src
	return src.tpe().Kind() == reflect.Struct
}

func (s structCopier) Cp(src, dst *ReflectEntity) {
	dstElemV := dst.elemVal()
	s.doCp(src.tpe(), dst.elemTpe(), src.val(), dstElemV)
}

func (s structCopier) Kd() []reflect.Kind {
	return []reflect.Kind{
		reflect.Struct,
	}
}

// doCp fill type fields
func (s structCopier) doCp(srcT, dstT reflect.Type, srcV, dstV reflect.Value) {
	num := srcT.NumField()
	for i := 0; i < num; i++ {
		srcFieldT := srcT.Field(i)
		dstFieldT := dstT.Field(i)

		if !srcFieldT.IsExported() || !dstFieldT.IsExported() {
			continue
		}

		if srcFieldT.Name != dstFieldT.Name {
			continue
		}

		dstFieldV := dstV.Field(i)
		if srcFieldT.Type != dstFieldT.Type {
			if srcFieldT.Type.Kind() == reflect.Struct {
				s.doCp(srcFieldT.Type, dstFieldT.Type, srcV.Field(i), dstFieldV)
				continue
			}
			continue
		}

		dstFieldV.Set(srcV.Field(i))
	}
}
