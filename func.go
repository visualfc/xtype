//go:build !llgo

package xtype

import (
	"reflect"
	"unsafe"
)

func ConvertFuncValue(typ Type, fn reflect.Value) reflect.Value {
	(*struct {
		typ Type
		ptr unsafe.Pointer
	})(unsafe.Pointer(&fn)).typ = typ
	return fn
}

func ConvertFunc(typ Type, i interface{}) interface{} {
	p := (*eface)(unsafe.Pointer(&i))
	return *(*interface{})(unsafe.Pointer(&eface{
		typ:  unsafe.Pointer(typ),
		word: p.word,
	}))
}
