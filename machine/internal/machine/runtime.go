package machine

type Runtime interface {
	Halt()
	Nop()
	Li(reg uint8, value int16)
	Lw(regDest uint8, regAddr uint8)
	Sw(regSrc uint8, regAddr uint8)
	Add(regDest uint8, regSrc1 uint8, regSrc2 uint8)
	Sub(regDest uint8, regSrc1 uint8, regSrc2 uint8)
	Mult(regDest uint8, regSrc1 uint8, regSrc2 uint8)
	Div(regDest uint8, regSrc1 uint8, regSrc2 uint8)
	J(addr int16)
	Jr(reg uint8)
	Beq(reg1 uint8, reg2 uint8, addr uint8)
	Bne(reg1 uint8, reg2 uint8, addr uint8)
	Inc(reg uint8)
	Dec(reg uint8)
}

func (m *Machine) Halt() {
	m.Running = false
}

func (m *Machine) Nop() {
	// Do nothing
}

func (m *Machine) Li(reg uint8, value int16) {
	m.Registers[reg] = value
}

func (m *Machine) Lw(regDest uint8, regAddr uint8) {
	m.Registers[regDest] = m.Registers[regAddr]
}

func (m *Machine) Sw(regSrc uint8, regAddr uint8) {
	// R1 stores a memory location
	// Therefore Calculate the address of the memory location
	addr := m.Registers[regAddr]
	// Store the contents of the source register in the memory location
	m.Memory[addr] = m.Registers[regSrc]
}

func (m *Machine) Add(regDest uint8, regSrc1 uint8, regSrc2 uint8) {
	v1 := m.Registers[regSrc1]
	v2 := m.Registers[regSrc2]
	m.Registers[regDest] = v2 + v1
}

func (m *Machine) Sub(regDest uint8, regSrc1 uint8, regSrc2 uint8) {
	v1 := m.Registers[regSrc1]
	v2 := m.Registers[regSrc2]
	m.Registers[regDest] = v1 - v2
}

func (m *Machine) Mult(regDest uint8, regSrc1 uint8, regSrc2 uint8) {
	v1 := m.Registers[regSrc1]
	v2 := m.Registers[regSrc2]
	m.Registers[regDest] = v1 * v2
}

func (m *Machine) Div(regDest uint8, regSrc1 uint8, regSrc2 uint8) {
	v1 := m.Registers[regSrc1]
	v2 := m.Registers[regSrc2]
	m.Registers[regDest] = v1 / v2
}

func (m *Machine) J(addr int16) {
	// Jump direct to the memory address
	// TODO Check if the memory address specified is an instruction address
	m.Registers[PC] = addr
}

func (m *Machine) Jr(reg uint8) {
	// Set the program counter to the memory address
	//stored in the specified register
	m.Registers[PC] = m.Registers[reg]
}

func (m *Machine) Beq(reg1 uint8, reg2 uint8, addr uint8) {
	v1 := m.Registers[reg1]
	v2 := m.Registers[reg2]
	v3 := m.Registers[addr]
	if v1 == v2 {
		instructionAddress := m.Memory[v3] + 1
		m.Registers[PC] = instructionAddress
	}
}
func (m *Machine) Bne(reg1 uint8, reg2 uint8, addr uint8) {
	v1 := m.Registers[reg1]
	v2 := m.Registers[reg2]
	v3 := m.Registers[addr]
	if v1 != v2 {
		instructionAddress := m.Memory[v3] + 1
		m.Registers[PC] = instructionAddress
	}
}

func (m *Machine) Inc(reg uint8) {
	m.Registers[reg]++
}

func (m *Machine) Dec(reg uint8) {
	m.Registers[reg]--
}
