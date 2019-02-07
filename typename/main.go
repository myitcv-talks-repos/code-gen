package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	typeName := os.Args[1]
	// 2
	found := false
	// 25
	pkgName := os.Getenv("GOPACKAGE")
	declFile := os.Getenv("GOFILE")
	// 3
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, declFile, nil, 0)
	if err != nil {
		log.Fatalf("failed to parse %v: %v", declFile, err)
	}
	// 4
Decls:
	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue Decls
			}
			for _, s := range d.Specs {
				s := s.(*ast.TypeSpec)
				if s.Name.Name != typeName {
					continue Decls
				}
				if id, ok := s.Type.(*ast.Ident); ok && id.Name == "int" {
					found = true
					break Decls
				}
			}
		}
	}
	// 5
	if !found {
		log.Fatalf("failed to find type declaration named %v of type int", typeName)
	}
	// 6
	output := fmt.Sprintf(`
	package %[1]v
	func (v %[2]v) TypeName() string {
		return "%[2]v"
	}
	`, pkgName, typeName)
	outfile := fmt.Sprintf("gen_%v_typename.go", typeName)
	// 7
	if err := ioutil.WriteFile(outfile, []byte(output), 0666); err != nil {
		log.Fatalf("failed to write output to %v: %v", outfile, err)
	}
	// 8
}
