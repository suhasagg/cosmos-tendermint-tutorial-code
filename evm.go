package main

import (
	"fmt"
	"math/big"
)

type Opcode byte

const (
	OP_ADD Opcode = iota + 1
	OP_SUB
	OP_MUL
	OP_DIV
	OP_PUSH
)

type Instruction struct {
	Opcode Opcode
	Arg    []byte
}

type Stack struct {
	data []*big.Int
}

func (s *Stack) Push(val *big.Int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() *big.Int {
	if len(s.data) == 0 {
		return nil
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) Peek() *big.Int {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

type EVM struct {
	Program    []Instruction
	Stack      Stack
	Memory     []byte
	Storage    map[common.Hash]common.Hash
	Gas        *big.Int
	CallValue  *big.Int
	Caller     common.Address
	Origin     common.Address
	GasPrice   *big.Int
	BlockCoinbase common.Address
	BlockNumber *big.Int
	BlockTime  *big.Int
	Difficulty *big.Int
}

func (e *EVM) Execute() error {
	for pc := 0; pc < len(e.Program); pc++ {
		instruction := e.Program[pc]
		switch instruction.Opcode {
		case OP_ADD:
			val1 := e.Stack.Pop()
			val2 := e.Stack.Pop()
			sum := new(big.Int).Add(val1, val2)
			e.Stack.Push(sum)
		case OP_PUSH:
			val := new(big.Int).SetBytes(instruction.Arg)
			e.Stack.Push(val)
		}
	}
	return nil
}

func main() {
	program := []Instruction{
		{Opcode: OP_PUSH, Arg: []byte{0x01}},
		{Opcode: OP_PUSH, Arg: []byte{0x02}},
		{Opcode: OP_ADD},
	}
	evm := EVM{
		Program: program,
		Stack:   Stack{},
		Memory:  []byte{},
		Storage: make(map[common.Hash]common.Hash),
		Gas:     big.NewInt(1000000),
	}
	err := evm.Execute()
	if err != nil {
		fmt.Printf("Error executing program: %v\n", err)
		return
	}
	result := evm.Stack.Pop()
	fmt.Printf("Result: %v\n", result)
}
