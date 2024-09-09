package main

import (
	"compiler/ast"
	"fmt"
	// "os"
	// "strings"
)

func main() {

	// data, err := os.ReadFile("test.lang")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// parsedData := string(data)
	// result := strings.ReplaceAll(parsedData, " ", "")

	// for _, r := range result {
	// 	if string(r) == "+" {

	// 	}
	// 	fmt.Println(string(r))
	// }

	ast := &ast.Program{
		Functions: []*ast.Function{
			{
				Name:   "main",
				Return: "int",
				Body: []ast.Node{
					&ast.ReturnStatement{
						Expr: &ast.BinaryExpr{
							Left:     &ast.IntLiteral{Value: 2},
							Operator: "+",
							Right:    &ast.IntLiteral{Value: 2},
						},
					},
				},
			},
		},
	}

	// Print the AST
	fmt.Println(ast.String())

}
