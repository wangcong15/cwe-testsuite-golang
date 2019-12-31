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
	
		"github.com/wangcong15/goassert"
	)
	
	func unTrustFunc() int {
		return 4
	}
	
	func main() {
		arr := [...]int{1, 2, 3, 4}
		idx := unTrustFunc()
		goassert.AssertGte(idx, 0)
		goassert.AssertLt(idx, len(arr))
		arr[idx] = 5
		log.Println(arr)
	}		
	`
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", rawCode, parser.ParseComments)
	ast.Print(fset, f)
}
