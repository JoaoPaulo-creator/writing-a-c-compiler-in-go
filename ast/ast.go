package ast

import "fmt"

type Node interface {
	String() string
}

// root node
type Program struct {
	Functions []*Function
}

/*

Each of these structs below, represents
a ast node. A binary operation, which basically means a sum or a subtraction
is a node, a functin its a node, return statement, literals, all of these
are the ast nodes.
*/

type Function struct {
	Name   string
	Return string
	Body   []Node
}

type ReturnStatement struct {
	Expr Node
}

type BinaryExpr struct {
	Left     Node
	Operator string
	Right    Node
}

type IntLiteral struct {
	Value int
}

/*
Each of these nodes needs a string implementation. If Im not wrong
everything will become a string at the end. So to  be in compliance with
the rules, every ast node will need to have a string function.

The logic for every function of course varies
*/

func (p *Program) String() string {
	result := ""
	for _, fn := range p.Functions {
		result += fn.String() + "\n"
	}
	return result
}

func (f *Function) String() string {
	return fmt.Sprintf("Function(%s) returns %s: %s", f.Name, f.Return, f.Body[0].String())
}

func (r *ReturnStatement) String() string {
	return fmt.Sprintf("Return(%s)", r.Expr.String())
}

func (b *BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator, b.Right.String())
}

func (i *IntLiteral) String() string {
	return fmt.Sprintf("%d", i.Value)
}
