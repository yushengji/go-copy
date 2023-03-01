package gocp

import (
	"reflect"
	"sync"
)

// entity object info
type entity struct {
	o    interface{}
	t    reflect.Type
	v    reflect.Value
	once sync.Once
}

func (e *entity) tpe() reflect.Type {
	if e.t == nil {
		e.t = reflect.TypeOf(e.o)
	}
	return e.t
}

func (e *entity) val() reflect.Value {
	e.once.Do(func() {
		e.v = reflect.ValueOf(e.o)
	})
	return e.v
}

func (e *entity) elemTpe() reflect.Type {
	v := e.tpe()
	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

func (e *entity) elemVal() reflect.Value {
	v := e.val()
	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}
