//go:build llgo

package xtype

import (
	"reflect"
	"unsafe"
)

//go:linkname closureOf reflect.closureOf
func closureOf(ftyp Type) Type

func ConvertFuncValue(typ Type, fn reflect.Value) reflect.Value {
	(*struct {
		typ Type
		ptr unsafe.Pointer
	})(unsafe.Pointer(&fn)).typ = closureOf(typ)
	return fn
}

func ConvertFunc(typ Type, i interface{}) interface{} {
	p := (*eface)(unsafe.Pointer(&i))
	return *(*interface{})(unsafe.Pointer(&eface{
		typ:  unsafe.Pointer(closureOf(typ)),
		word: p.word,
	}))
}
