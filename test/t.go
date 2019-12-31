package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	rawCode := `package main

	import "log"
	
	func main() {
		int_val := 4
		arr_val := [...]int{1, 2, 3}
		log.Println(arr_val[0:int_val])
	}		
	`
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", rawCode, parser.ParseComments)
	ast.Print(fset, f)
}
