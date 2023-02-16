package util

import "fmt"

type Instruction struct {
	R3 uint8
	R2 uint8
	R1 uint8
	O  uint8
}

func (i *Instruction) String() string {
	return fmt.Sprintf("INSTRUCTION = \033[1;32m0x%00X 0x%00X 0x%00X 0x%00X\u001B[0m", i.O, i.R1, i.R2, i.R3)
}

func ParseByteCode(bytecode int16) *Instruction {
	return &Instruction{
		R3: uint8(bytecode & 0xF),
		R2: uint8((bytecode >> 4) & 0xF),
		R1: uint8((bytecode >> 8) & 0xF),
		O:  uint8((bytecode >> 12) & 0xF),
	}
}
