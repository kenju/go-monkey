package evaluator

import (
	"github.com/kenju/go-monkey/ast"
	"github.com/kenju/go-monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{
		Node: node,
	}
}
