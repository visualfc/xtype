package xtype_test

import (
	"reflect"
	"testing"

	"github.com/visualfc/xtype"
)

var (
	benchmarkFunc   TFunc
	benchmarkResult int
	benchmarkAny    interface{}
	benchmarkType   interface{}
)

type benchmarkIntType int

type benchmarkStruct struct {
	N int
	S string
}

type benchmarkNamedStruct benchmarkStruct

func BenchmarkTypeOf(b *testing.B) {
	value := interface{}(benchmarkStruct{N: 100, S: "xtype"})

	b.Run("XType", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkType = xtype.TypeOf(value)
		}
	})
	b.Run("Reflect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkType = reflect.TypeOf(value)
		}
	})
}

func BenchmarkReadInt(b *testing.B) {
	value := interface{}(benchmarkIntType(100))
	reflected := reflect.ValueOf(value)

	b.Run("XType", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkResult = xtype.Int(value)
		}
	})
	b.Run("TypeAssertion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkResult = int(value.(benchmarkIntType))
		}
	})
	b.Run("Reflect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkResult = int(reflected.Int())
		}
	})
}

func BenchmarkMakeInt(b *testing.B) {
	typ := xtype.TypeOf(benchmarkIntType(0))
	reflectType := reflect.TypeOf(benchmarkIntType(0))
	value := 100

	b.Run("XType", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = xtype.MakeInt(typ, value)
		}
	})
	b.Run("GoConversion", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = benchmarkIntType(value)
		}
	})
	b.Run("Reflect", func(b *testing.B) {
		reflected := reflect.ValueOf(value)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = reflected.Convert(reflectType).Interface()
		}
	})
}

func BenchmarkConvertInt(b *testing.B) {
	typ := xtype.TypeOf(benchmarkIntType(0))
	reflectType := reflect.TypeOf(benchmarkIntType(0))
	value := interface{}(int(100))

	b.Run("XType", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = xtype.ConvertInt(typ, value)
		}
	})
	b.Run("Reflect", func(b *testing.B) {
		reflected := reflect.ValueOf(value)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = reflected.Convert(reflectType).Interface()
		}
	})
}

func BenchmarkConvertDirect(b *testing.B) {
	typ := xtype.TypeOf(benchmarkNamedStruct{})
	reflectType := reflect.TypeOf(benchmarkNamedStruct{})
	value := interface{}(benchmarkStruct{N: 100, S: "xtype"})

	b.Run("XType", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = xtype.ConvertDirect(typ, value)
		}
	})
	b.Run("Reflect", func(b *testing.B) {
		reflected := reflect.ValueOf(value)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = reflected.Convert(reflectType).Interface()
		}
	})
}

func BenchmarkNegInt(b *testing.B) {
	value := interface{}(benchmarkIntType(100))

	b.Run("XType", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = xtype.NegInt(value)
		}
	})
	b.Run("Go", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = -value.(benchmarkIntType)
		}
	})
}

func BenchmarkAlloc(b *testing.B) {
	typ := xtype.TypeOf(benchmarkStruct{})
	reflectType := reflect.TypeOf(benchmarkStruct{})

	b.Run("XType", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = xtype.Alloc(typ)
		}
	})
	b.Run("Reflect", func(b *testing.B) {
		zero := reflect.Zero(reflectType)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = zero.Interface()
		}
	})
}

func BenchmarkNew(b *testing.B) {
	value := benchmarkStruct{}
	typ := xtype.TypeOf(value)
	ptrType := xtype.TypeOf(&value)
	reflectType := reflect.TypeOf(value)

	b.Run("XType", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = xtype.New(typ, ptrType)
		}
	})
	b.Run("Reflect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			benchmarkAny = reflect.New(reflectType).Interface()
		}
	})
}

func BenchmarkConvertFunc(b *testing.B) {
	base := N(100)
	fn := func(n N) N {
		return n + base
	}
	typ := xtype.TypeOfType(reflect.TypeOf(TFunc(nil)))

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkFunc = xtype.ConvertFunc(typ, fn).(TFunc)
	}
}

func BenchmarkConvertFuncValue(b *testing.B) {
	base := N(100)
	fn := reflect.ValueOf(func(n N) N {
		return n + base
	})
	typ := xtype.TypeOfType(reflect.TypeOf(TFunc(nil)))

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkFunc = xtype.ConvertFuncValue(typ, fn).Interface().(TFunc)
	}
}

func BenchmarkConvertedFuncCall(b *testing.B) {
	base := N(100)
	fn := func(n N) N {
		return n + base
	}
	typ := xtype.TypeOfType(reflect.TypeOf(TFunc(nil)))
	converted := xtype.ConvertFunc(typ, fn).(TFunc)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkResult = converted(i)
	}
}
