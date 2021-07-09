package xtype_test

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/visualfc/xtype"
)

const filename = "<src>"

func makePkg(src string) (*types.Package, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.DeclarationErrors)
	if err != nil {
		return nil, err
	}
	// use the package name as package path
	conf := types.Config{Importer: importer.Default()}
	return conf.Check(file.Name.Name, fset, []*ast.File{file}, nil)
}

var src = `package p
var i int
type Point struct {
	X int
	Y int
}
`

type testEntry struct {
	src, str string
}

// dup returns a testEntry where both src and str are the same.
func dup(s string) testEntry {
	return testEntry{s, s}
}

var basicTypes = []testEntry{
	// basic
	dup(`bool`),
	dup(`int`),
	dup(`int8`),
	dup(`int16`),
	dup(`int32`),
	dup(`int64`),
	dup(`uint`),
	dup(`uint8`),
	dup(`uint16`),
	dup(`uint32`),
	dup(`uint64`),
	dup(`uintptr`),
	dup(`float32`),
	dup(`float64`),
	dup(`complex64`),
	dup(`complex128`),
	dup(`string`),
	dup(`unsafe.Pointer`),
}

func TestTypes(t *testing.T) {
	var tests []testEntry
	tests = append(tests, basicTypes...)

	for _, test := range tests {
		src := `package p; import "unsafe"; type _ unsafe.Pointer; type T ` + test.src
		pkg, err := makePkg(src)
		if err != nil {
			t.Errorf("%s: %s", src, err)
			continue
		}
		typ := pkg.Scope().Lookup("T").Type().Underlying()
		rt, err := xtype.ToType(typ)
		if err != nil {
			t.Errorf("%s: ToType error %v", test.src, err)
		}
		if got := rt.String(); got != test.str {
			t.Errorf("%s: got %s, want %s", test.src, got, test.str)
		}
	}
}
