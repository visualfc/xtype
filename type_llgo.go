//go:build llgo

package xtype

import "unsafe"

type tflag uint8

type _type struct {
	size_       uintptr
	ptrBytes    uintptr // number of (prefix) bytes in the type that can contain pointers
	hash        uint32  // hash of type; avoids computation in hash tables
	tflag       tflag   // extra type information flags
	align_      uint8   // alignment of variable with this type
	fieldAlign_ uint8   // alignment of struct field with this type
	kind        uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// GCData stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, GCData is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata     *byte
	str_       string // string form
	ptrToThis_ *_type // type for pointer to this type, may be nil
}

const (
	kindDirectIface = 1 << 5
	kindMask        = kindDirectIface - 1
	kindPointer     = 22
)

func isDirectIface(t *_type) bool {
	return t.kind&kindDirectIface != 0
}

type ptrType struct {
	_type
	elem *_type
}
