package vm

import "github.com/kenju/go-monkey/compiler"

type Vm struct {
}

func New(bytecode *compiler.Bytecode) *Vm {
	return &Vm{}
}

func (v *Vm) Run() error {
	return nil
}

func (v *Vm) StackTop() {
}