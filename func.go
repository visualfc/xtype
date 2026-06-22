//go:build !llgo

package xtype

import (
	"reflect"
	"unsafe"
)

func ConvertFunc(fn reflect.Value, typ Type) reflect.Value {
	(*struct {
		typ Type
		ptr unsafe.Pointer
	})(unsafe.Pointer(&fn)).typ = typ
	return fn
}
