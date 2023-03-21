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
	return finallyElemTpe(e.tpe())
}

func (e *ReflectEntity) elemVal() reflect.Value {
	return finallyElemVal(e.val())
}

func (e *ReflectEntity) elem() *ReflectEntity {
	return &ReflectEntity{
		o: e.o,
		t: e.elemTpe(),
		v: e.elemVal(),
	}
}

func (e *ReflectEntity) isNil() bool {
	return isNil(e.val())
}

func (e *ReflectEntity) setPtrVal(val reflect.Value) {
	setPtrVal(e.val(), val)
}

func setPtrVal(dst reflect.Value, val reflect.Value) {
	if dst.Kind() == reflect.Ptr && dst.Elem().Kind() == reflect.Ptr {
		setPtrVal(dst.Elem(), val)
		return
	}
	dst.Set(val)
}

func isNil(val reflect.Value) bool {
	if val.Kind() == reflect.Ptr {
		if val.Elem().Kind() == reflect.Ptr {
			return isNil(val.Elem())
		}
		return val.IsNil()
	}
	return false
}

func finallyElemTpe(tpe reflect.Type) reflect.Type {
	if tpe.Kind() == reflect.Ptr {
		if tpe.Elem().Kind() == reflect.Ptr {
			return finallyElemTpe(tpe.Elem())
		}

		return tpe.Elem()
	}

	return tpe
}

func finallyElemVal(val reflect.Value) reflect.Value {
	if val.Kind() == reflect.Ptr {
		if val.Elem().Kind() == reflect.Ptr {
			return finallyElemVal(val.Elem())
		}

		return val.Elem()
	}

	return val
}
