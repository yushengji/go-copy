package gocp

import (
	"reflect"
	"sync"
)

// ReflectEntity object info
type ReflectEntity struct {
	// o if you have t, v it not necessary
	o    interface{}
	t    reflect.Type
	v    reflect.Value
	once sync.Once
}

func NewReflectEntityByObj(o interface{}) *ReflectEntity {
	if o == nil {
		return nil
	}

	return &ReflectEntity{
		o: o,
	}
}

func NewReflectEntityByTpeVal(tpe reflect.Type, val reflect.Value) *ReflectEntity {
	if tpe == nil {
		return nil
	}

	return &ReflectEntity{
		t: tpe,
		v: val,
	}
}

func (e *ReflectEntity) tpe() reflect.Type {
	if e.t == nil {
		e.t = reflect.TypeOf(e.o)
	}
	return e.t
}

func (e *ReflectEntity) val() reflect.Value {
	e.once.Do(func() {
		if e.v.Kind() != reflect.Invalid {
			return
		}

		e.v = reflect.ValueOf(e.o)
	})
	return e.v
}

func (e *ReflectEntity) elemTpe() reflect.Type {
	v := e.tpe()
	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

func (e *ReflectEntity) elemVal() reflect.Value {
	v := e.val()
	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

func (e *ReflectEntity) elem() *ReflectEntity {
	return &ReflectEntity{
		o: e.o,
		t: e.elemTpe(),
		v: e.elemVal(),
	}
}
