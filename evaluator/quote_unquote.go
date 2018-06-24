package evaluator

import (
	"github.com/kenju/go-monkey/object"
	"github.com/kenju/go-monkey/ast"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{
		Node: node,
	}
}