package gocp

import (
	"reflect"
)

type structCopier struct{}

func (s structCopier) Check(src *ReflectEntity) bool {
	// kind dst == src
	return src.tpe().Kind() == reflect.Struct
}

func (s structCopier) Cp(src, dst *ReflectEntity) {
	// set ** struct type
	if dst.isNil() {
		dst.setPtrVal(reflect.New(dst.elemTpe()))
	}

	s.doCp(src.tpe(), dst.elemTpe(), src.val(), dst.elemVal())
}

func (s structCopier) Kd() []reflect.Kind {
	return []reflect.Kind{
		reflect.Struct,
	}
}

// doCp fill type fields
func (s structCopier) doCp(srcT, dstT reflect.Type, srcV, dstV reflect.Value) {
	num := dstT.NumField()
	srcFieldNum := srcT.NumField()
	srcFields := make([]reflect.StructField, 0, srcFieldNum)
	for i := 0; i < srcFieldNum; i++ {
		srcFields = append(srcFields, srcT.Field(i))
	}

	for i := 0; i < num; i++ {
		dstF := dstT.Field(i)
		if !dstF.IsExported() {
			continue
		}

		dstFV := dstV.Field(i)
		if s.tryPlugin(srcT, dstF, srcV, dstFV) {
			continue
		}

		srcF, ok := srcT.FieldByName(dstF.Name)
		if !ok {
			continue
		}

		s.defCp(srcF, dstF, srcV.FieldByName(dstF.Name), dstFV)
	}
}

func (s structCopier) tryPlugin(srcT reflect.Type, dstF reflect.StructField, srcV, dstFV reflect.Value) bool {
	set := false
	var srcF reflect.StructField
	for _, plugin := range fieldPlugins {
		if !plugin.Check(dstF) {
			continue
		}

		matched, ok := plugin.Match(srcT, dstF)
		if !ok {
			continue
		}

		if !matched.IsExported() {
			continue
		}

		srcF = matched
	}

	resultV := srcV.FieldByName(srcF.Name)
	for _, otherPlugin := range fieldPlugins {
		if !otherPlugin.Check(dstF) || !otherPlugin.Verify(srcF, dstF) {
			continue
		}

		set = true
		resultV = otherPlugin.Transform(resultV, dstFV)
	}

	if set {
		dstFV.Set(resultV)
	}

	return set
}

func (s structCopier) defCp(srcFieldT, dstFieldT reflect.StructField, srcFieldV, dstFieldV reflect.Value) {
	if srcFieldT.Name != dstFieldT.Name {
		return
	}

	fieldKd := srcFieldT.Type.Kind()
	if fieldKd != dstFieldT.Type.Kind() {
		return
	}

	if srcFieldT.Type != dstFieldT.Type {
		if fieldKd == reflect.Struct {
			s.doCp(srcFieldT.Type, dstFieldT.Type, srcFieldV, dstFieldV)
			return
		}

		if fieldKd == reflect.Array || fieldKd == reflect.Slice {
			doCp(&ReflectEntity{
				t: srcFieldT.Type,
				v: srcFieldV,
			}, &ReflectEntity{
				t: dstFieldT.Type,
				v: dstFieldV,
			})
			return
		}
		return
	}

	dstFieldV.Set(srcFieldV)
}
