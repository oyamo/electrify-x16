package util

import "fmt"

type Instruction struct {
	R3 uint8
	R2 uint8
	R1 uint8
	O  uint8
}

var ISet = map[uint8]string{
	0x00: "halt",
	0x01: "nop",
	0x02: "li",
	0x03: "lw",
	0x04: "sw",
	0x05: "add",
	0x06: "sub",
	0x07: "mult",
	0x08: "div",
	0x09: "j",
	0x0A: "jr",
	0x0B: "beq",
	0x0C: "bne",
	0x0D: "inc",
	0x0E: "dec",
}

func (i *Instruction) String() string {
	return fmt.Sprintf("INSTRUCTION = \033[1;32m0x%03X \033[1;37m(%s)\033[1;32m 0x%03X 0x%03X 0x%03X\u001B[0m",
		i.O, ISet[i.O], i.R1, i.R2, i.R3)
}

func ParseByteCode(bytecode int16) *Instruction {
	return &Instruction{
		R3: uint8(bytecode & 0xF),
		R2: uint8((bytecode >> 4) & 0xF),
		R1: uint8((bytecode >> 8) & 0xF),
		O:  uint8((bytecode >> 12) & 0xF),
	}
}
