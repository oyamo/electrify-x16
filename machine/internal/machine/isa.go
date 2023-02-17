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
