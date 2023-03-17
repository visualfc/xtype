//go:build js && !wasm
// +build js,!wasm

package xtype

import (
	"reflect"
	"unsafe"

	"github.com/gopherjs/gopherjs/js"
)

type eface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

type Type = *js.Object

func TypeOf(i interface{}) Type {
	return js.InternalObject(i).Get("constructor")
}

func TypeOfType(typ reflect.Type) Type {
	return js.InternalObject(typ).Get("jsType")
}

func Bytes(i interface{}) []byte {
	v1 := js.InternalObject(i)
	b := make([]byte, v1.Get("$length").Int(), v1.Get("$capacity").Int())
	v2 := js.InternalObject(b)
	v2.Set("$array", v1.Get("$array"))
	v2.Set("$offset", v1.Get("$offset"))
	return b
}

func Runes(i interface{}) []rune {
	v1 := js.InternalObject(i)
	b := make([]rune, v1.Get("$length").Int(), v1.Get("$capacity").Int())
	v2 := js.InternalObject(b)
	v2.Set("$array", v1.Get("$array"))
	v2.Set("$offset", v1.Get("$offset"))
	return b
}

func Bool(i interface{}) bool {
	v := js.InternalObject(i).Get("$val")
	return v.Bool()
}

func Int(i interface{}) int {
	v := js.InternalObject(i).Get("$val")
	return v.Int()
}

func Int8(i interface{}) int8 {
	v := js.InternalObject(i).Get("$val")
	return int8(v.Int64())
}

func Int16(i interface{}) int16 {
	v := js.InternalObject(i).Get("$val")
	return int16(v.Int64())
}

func Int32(i interface{}) int32 {
	v := js.InternalObject(i).Get("$val")
	return int32(v.Int64())
}

func Int64(i interface{}) int64 {
	v := js.InternalObject(i).Get("$val")
	return v.Int64()
}

func Uint(i interface{}) uint {
	v := js.InternalObject(i).Get("$val")
	return uint(v.Uint64())
}

func Uint8(i interface{}) uint8 {
	v := js.InternalObject(i).Get("$val")
	return uint8(v.Uint64())
}

func Uint16(i interface{}) uint16 {
	v := js.InternalObject(i).Get("$val")
	return uint16(v.Uint64())
}

func Uint32(i interface{}) uint32 {
	v := js.InternalObject(i).Get("$val")
	return uint32(v.Uint64())
}

func Uint64(i interface{}) uint64 {
	v := js.InternalObject(i).Get("$val")
	return v.Uint64()
}

func Uintptr(i interface{}) uintptr {
	v := js.InternalObject(i).Get("$val")
	return uintptr(v.Uint64())
}

func Float32(i interface{}) float32 {
	v := js.InternalObject(i).Get("$val")
	return float32(v.Float())
}

func Float64(i interface{}) float64 {
	v := js.InternalObject(i).Get("$val")
	return v.Float()
}

func Complex64(i interface{}) complex64 {
	v := js.InternalObject(i)
	return complex(float32(v.Get("$real").Float()), float32(v.Get("$imag").Float()))
}

func Complex128(i interface{}) complex128 {
	v := js.InternalObject(i)
	return complex(v.Get("$real").Float(), v.Get("$imag").Float())
}

func String(i interface{}) string {
	v := js.InternalObject(i)
	return v.Get("$val").String()
}

func Pointer(i interface{}) unsafe.Pointer {
	v := js.InternalObject(i)
	return unsafe.Pointer(v.Get("$val").Unsafe())
}

// Make change interface type and return
func Make(typ Type, i interface{}) interface{} {
	v := js.InternalObject(i)
	v.Set("constructor", typ)
	return i
}

func ConvertPtr(typ Type, i interface{}) interface{} {
	rt := toReflectType(typ)
	return reflect.ValueOf(i).Convert(rt).Interface()
}

// //go:linkname typedmemmove reflect.typedmemmove
// func typedmemmove(t Type, dst unsafe.Pointer, src unsafe.Pointer)

// //go:linkname unsafe_New reflect.unsafe_New
// func unsafe_New(t Type) unsafe.Pointer

func ptrTo(typ Type) Type {
	return js.Global.Call("$ptrType", typ)
}

// func unsafe_New(typ Type) unsafe.Pointer {
// 	switch reflect.Kind(typ.Get("kind").Int()) {
// 	case reflect.Struct:
// 		return unsafe.Pointer(typ.Get("ptr").New().Unsafe())
// 	case reflect.Array:
// 		return unsafe.Pointer(typ.Call("zero").Unsafe())
// 	default:
// 		return unsafe.Pointer(js.Global.Call("$newDataPointer", typ.Call("zero"), ptrTo(typ)).Unsafe())
// 	}
// }

// func newObject(typ Type) *js.Object {
// 	switch reflect.Kind(typ.Get("kind").Int()) {
// 	case reflect.Struct:
// 		return typ.Get("ptr").New()
// 	case reflect.Array:
// 		return typ.Call("zero")
// 	default:
// 		return js.Global.Call("$newDataPointer", typ.Call("zero"), ptrTo(typ))
// 	}
// }

// func typedmemmove(t *rtype, dst, src unsafe.Pointer) {
// 	js.InternalObject(dst).Call("$set", js.InternalObject(src).Call("$get"))
// }

// convert copy
func ConvertDirect(typ Type, i interface{}) interface{} {
	rt := toReflectType(typ)
	return reflect.ValueOf(i).Convert(rt).Interface()
}

func toReflectType(typ Type) reflect.Type {
	return toType(reflectType(typ))
}

//go:linkname toType reflect.toType
func toType(t *_type) reflect.Type

//go:linkname reflectType reflect.reflectType
func reflectType(typ *js.Object) *_type

func ConvertBool(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertInt(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertInt8(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertInt16(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertInt32(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertInt64(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertUint(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertUint8(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertUint16(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertUint32(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertUint64(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertUintptr(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertFloat32(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertFloat64(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertComplex64(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertComplex128(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func ConvertString(typ Type, i interface{}) interface{} {
	return Make(typ, i)
}

func Not(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), !v.Get("$val").Bool())
}

func NegInt(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -v.Get("$val").Int())
}

func NegInt8(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -int8(v.Get("$val").Int64()))
}

func NegInt16(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -int16(v.Get("$val").Int64()))
}

func NegInt32(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -int32(v.Get("$val").Int64()))
}

func NegInt64(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -int64(v.Get("$val").Int64()))
}

func NegUint(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -uint(v.Get("$val").Uint64()))
}

func NegUint8(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -uint8(v.Get("$val").Uint64()))
}

func NegUint16(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -uint16(v.Get("$val").Uint64()))
}

func NegUint32(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -uint32(v.Get("$val").Uint64()))
}

func NegUint64(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -uint64(v.Get("$val").Uint64()))
}

func NegUintptr(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -uintptr(v.Get("$val").Uint64()))
}

func NegFloat32(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -float32(v.Get("$val").Float()))
}

func NegFloat64(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -float64(v.Get("$val").Float()))
}

func NegComplex64(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -complex(float32(v.Get("$real").Float()), float32(v.Get("$imag").Float())))
}

func NegComplex128(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), -complex(v.Get("$real").Float(), v.Get("$imag").Float()))
}

func XorInt(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^int(v.Get("$val").Int()))
}

func XorInt8(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^int8(v.Get("$val").Int()))
}

func XorInt16(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^int16(v.Get("$val").Int()))
}

func XorInt32(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^int32(v.Get("$val").Int()))
}

func XorInt64(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^int64(v.Get("$val").Int64()))
}

func XorUint(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^uint(v.Get("$val").Uint64()))
}

func XorUint8(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^uint8(v.Get("$val").Uint64()))
}

func XorUint16(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^uint16(v.Get("$val").Uint64()))
}

func XorUint32(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^uint32(v.Get("$val").Uint64()))
}

func XorUint64(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^uint64(v.Get("$val").Uint64()))
}

func XorUintptr(i interface{}) interface{} {
	v := js.InternalObject(i)
	return Make(v.Get("constructor"), ^uintptr(v.Get("$val").Uint64()))
}

func MakeBool(typ Type, v bool) interface{} {
	return Make(typ, v)
}

func MakeInt(typ Type, v int) interface{} {
	return Make(typ, v)
}

func MakeInt8(typ Type, v int8) interface{} {
	return Make(typ, v)
}

func MakeInt16(typ Type, v int16) interface{} {
	return Make(typ, v)
}

func MakeInt32(typ Type, v int32) interface{} {
	return Make(typ, v)
}

func MakeInt64(typ Type, v int64) interface{} {
	return Make(typ, v)
}

func MakeUint(typ Type, v uint) interface{} {
	return Make(typ, v)
}

func MakeUint8(typ Type, v uint8) interface{} {
	return Make(typ, v)
}

func MakeUint16(typ Type, v uint16) interface{} {
	return Make(typ, v)
}

func MakeUint32(typ Type, v uint32) interface{} {
	return Make(typ, v)
}

func MakeUint64(typ Type, v uint64) interface{} {
	return Make(typ, v)
}

func MakeUintptr(typ Type, v uintptr) interface{} {
	return Make(typ, v)
}

func MakeFloat32(typ Type, v float32) interface{} {
	return Make(typ, v)
}

func MakeFloat64(typ Type, v float64) interface{} {
	return Make(typ, v)
}

func MakeComplex64(typ Type, v complex64) interface{} {
	return Make(typ, v)
}

func MakeComplex128(typ Type, v complex128) interface{} {
	return Make(typ, v)
}

func MakeString(typ Type, v string) interface{} {
	return Make(typ, v)
}

func Alloc(typ Type) interface{} {
	t := toReflectType(typ)
	if t.Kind() == reflect.Ptr {
		return reflect.New(t.Elem()).Interface()
	}
	return reflect.New(t).Elem().Interface()
}

func New(typ, ptrto Type) interface{} {
	t := toReflectType(typ)
	return reflect.New(t).Interface()
}

// func NewPointer(typ Type) unsafe.Pointer {
// 	return unsafe_New(typ)
// }

// func SetPointer(i interface{}, word unsafe.Pointer) interface{} {
// 	p := (*eface)(unsafe.Pointer(&i))
// 	p.word = word
// 	return i
// }

// func SetType(i interface{}, typ Type) interface{} {
// 	p := (*eface)(unsafe.Pointer(&i))
// 	p.typ = unsafe.Pointer(typ)
// 	return i
// }

func ConvertFunc(fn reflect.Value, typ Type) reflect.Value {
	(*struct {
		typ *_type
		ptr unsafe.Pointer
	})(unsafe.Pointer(&fn)).typ = reflectType(typ)
	return fn
}
