package compiler

import (
	"github.com/kenju/go-monkey/code"
	"github.com/kenju/go-monkey/object"
	"github.com/kenju/go-monkey/ast"
)

type Compiler struct {
	instructions code.Instructions
	constants []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants: []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants: c.constants,
	}
}

// Bytecode is what we'll pass to the VM
type Bytecode struct {
	Instructions code.Instructions
	Constants []object.Object
}