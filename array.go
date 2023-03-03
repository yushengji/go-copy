package gocp

import (
	"reflect"
)

type arrayCopier struct{}

func (a arrayCopier) Check(src *ReflectEntity) bool {
	return src.tpe().Kind() == reflect.Array
}

func (a arrayCopier) Cp(src, dst *ReflectEntity) {
	sL := src.elemVal().Len()
	dL := dst.elemVal().Len()
	loop := sL
	if dL < sL {
		loop = dL
	}

	// cp array items
	for i := 0; i < loop; i++ {
		s := src.elemVal().Index(i)
		d := dst.elemVal().Index(i)

		// cp
		doCp(&ReflectEntity{o: s.Interface(), t: s.Type(), v: s},
			&ReflectEntity{t: d.Addr().Type(), v: d.Addr()})
	}
}

func (a arrayCopier) Kd() []reflect.Kind {
	return []reflect.Kind{
		reflect.Array,
	}
}
