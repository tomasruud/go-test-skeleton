package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing file arg")
		os.Exit(1)
	}

	file := os.Args[1]

	src, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load input file: %v\n", err)
		os.Exit(1)
	}

	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, file, src, parser.AllErrors)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parsing file: %v\n", err)
		os.Exit(1)
	}

	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
				// Extract receiver type
				recvType := funcDecl.Recv.List[0].Type
				switch t := recvType.(type) {

				case *ast.Ident:
					fmt.Fprintln(os.Stdout, testDecl(t.Name, funcDecl.Name.Name))

				case *ast.StarExpr:
					if ident, ok := t.X.(*ast.Ident); ok {
						fmt.Fprintln(os.Stdout, testDecl(ident.Name, funcDecl.Name.Name))
					}
				}
			} else {
				fmt.Fprintln(os.Stdout, testDecl(funcDecl.Name.Name))
			}
		}
	}
}

func testDecl(things ...string) string {
	var name string
	for _, thing := range things {
		name += namePart(thing)
	}
	return "func Test" + name + "(t *testing.T) {\n\tt.Skip(\"not implemented\")\n}\n"
}

func namePart(thing string) string {
	if unicode.IsUpper(rune(thing[0])) {
		return thing
	} else {
		return string(unicode.ToUpper(rune(thing[0]))) + thing[1:]
	}
}
