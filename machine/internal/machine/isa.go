package machine

const (
	Halt = uint8(0x00)
	Nop  = uint8(0x01)
	Li   = uint8(0x02)
	Lw   = uint8(0x03)
	Sw   = uint8(0x04)
	Add  = uint8(0x05)
	Sub  = uint8(0x06)
	Mult = uint8(0x07)
	Div  = uint8(0x08)
	J    = uint8(0x09)
	Jr   = uint8(0x0A)
	Beq  = uint8(0x0B)
	Bne  = uint8(0x0C)
	Inc  = uint8(0x0D)
	Dec  = uint8(0x0E)
)

var ISet = map[int16]string{
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
