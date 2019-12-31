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
	
	func main() {
		happyNewYear := []string{"Happy", "New", "Year"}
		lens := len(happyNewYear)
		for i := 0; i <= lens; i++ {
			log.Println(happyNewYear[i])
		}
	}	
	`
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", rawCode, parser.ParseComments)
	ast.Print(fset, f)
}
