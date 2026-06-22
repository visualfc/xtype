//go:build llgo

package xtype

import (
	"reflect"
	"unsafe"
)

//go:linkname closureOf reflect.closureOf
func closureOf(ftyp Type) Type

func ConvertFunc(fn reflect.Value, typ Type) reflect.Value {
	(*struct {
		typ Type
		ptr unsafe.Pointer
	})(unsafe.Pointer(&fn)).typ = closureOf(typ)
	return fn
}
