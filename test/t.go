package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	rawCode := `package main

	import (
		"log"
	)
	
	func unTrustFunc() int {
		return 0
	}
	
	func main() {
		var intVal int = unTrustFunc()
		intVal = 5 / intVal
		log.Println(intVal)
	}	
	`
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", rawCode, parser.ParseComments)
	ast.Print(fset, f)
}
