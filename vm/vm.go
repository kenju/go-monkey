package vm

import (
	"github.com/kenju/go-monkey/compiler"
	"github.com/kenju/go-monkey/object"
	"github.com/kenju/go-monkey/code"
)

const StackSize = 2048

type VM struct {
	constants []object.Object
	instructions code.Instructions

	stack []object.Object
	sp int // Always points to the next value. Top of stack is stack[sp-1]
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants: bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp: 0,
	}
}

func (v *VM) Run() error {
	return nil
}

func (v *VM) StackTop() {
}