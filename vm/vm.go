package vm

import (
	"github.com/kenju/go-monkey/compiler"
	"github.com/kenju/go-monkey/object"
	"github.com/kenju/go-monkey/code"
	"fmt"
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

func (vm *VM) Run() error {
	// iterate through the whole instructions
	for ip := 0; ip < len(vm.instructions); ip++ {
		// convert one of instruction to an Opcode
		op := code.Opcode(vm.instructions[ip])

		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1] // -1 because sp(stack pointer) always points to the next value
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp ++

	return nil
}