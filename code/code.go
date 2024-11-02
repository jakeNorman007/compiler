package code

import (
  "fmt"
)

type Instructions []byte

type Opcode byte

type Definitions struct {
  Name          string
  OperandWidths []int
}

const (
  OpConstant Opcode = iota
)

var definitions = map[Opcode]*Definitions {
  OpConstant: {"Opconstant", []int{2}},
}

func Lookup(op byte) (*Definitions, error) {
  def, ok := definitions[Opcode(op)]
  if !ok {
    return nil, fmt.Errorf("Opcode %d undefined", op)
  }
  return def, nil
}
