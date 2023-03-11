package gocp

import "reflect"

type structCopier struct{}

func (s structCopier) Check(src *ReflectEntity) bool {
	// kind dst == src
	return src.tpe().Kind() == reflect.Struct
}

func (s structCopier) Cp(src, dst *ReflectEntity) {
	s.doCp(src.tpe(), dst.elemTpe(), src.val(), dst.elemVal())
}

func (s structCopier) Kd() []reflect.Kind {
	return []reflect.Kind{
		reflect.Struct,
	}
}

// doCp fill type fields
func (s structCopier) doCp(srcT, dstT reflect.Type, srcV, dstV reflect.Value) {
	num := srcT.NumField()
loop:
	for i := 0; i < num; i++ {
		srcFieldT := srcT.Field(i)
		dstFieldT := dstT.Field(i)

		if !srcFieldT.IsExported() || !dstFieldT.IsExported() {
			continue
		}

		set := false
		dstFieldV := dstV.Field(i)
		srcFieldV := srcV.Field(i)
		resultV := srcFieldV
		for plugin := range fieldPlugins {
			if plugin.Check(srcFieldT, dstFieldT) {
				resultV = plugin.To(srcFieldT, dstFieldT, resultV, dstFieldV)
				set = true
			}
		}

		if set {
			dstFieldV.Set(resultV)
			continue loop
		}

		s.defCp(srcFieldT, dstFieldT, srcFieldV, dstFieldV)
	}
}

func (s structCopier) defCp(srcFieldT, dstFieldT reflect.StructField, srcFieldV, dstFieldV reflect.Value) {
	if srcFieldT.Name != dstFieldT.Name {
		return
	}

	if srcFieldT.Type.Kind() != dstFieldT.Type.Kind() {
		return
	}

	if srcFieldT.Type != dstFieldT.Type {
		if srcFieldT.Type.Kind() == reflect.Struct {
			s.doCp(srcFieldT.Type, dstFieldT.Type, srcFieldV, dstFieldV)
			return
		}
		return
	}

	dstFieldV.Set(srcFieldV)
}
