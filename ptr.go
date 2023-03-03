package gocp

import "reflect"

type ptrCopier struct{}

func (p ptrCopier) Check(src *ReflectEntity) bool {
	return src.tpe().Kind() == reflect.Ptr
}

func (p ptrCopier) Cp(src, dst *ReflectEntity) {
	srcE := src.elem()

	for _, c := range typePlugins[1:] {
		if c.Check(srcE) {
			c.Cp(srcE, dst)
			return
		}
	}
}

func (p ptrCopier) Kd() []reflect.Kind {
	return []reflect.Kind{reflect.Ptr}
}
