package machine

import (
	"fmt"
	"machine/pkg/util"
	"time"
)

const MemorySize = (2 << 15) / 16
const ProgramStartAddress = 0x0000CFFF / 16
const CpuFreq = 8

// Registers
const (
	_ = iota
	R1
	R2
	R3
	PC
	COND
)

// Machine A simulation of a 16Bit computer
type Machine struct {
	Registers      map[uint8]int16 // R1-3, PC, COND
	Memory         map[int16]int16
	Running        bool
	FreeMemory     int16
	CpuFreq        int // CPU Fref in Mhz
	progStopOffset int16
}

// Memory Layout
// 0x00 - 0x0 0x0000CFFF

// Boot Set up the machine by configuring memory and CPU registers
// To maje the machine ready for use
func Boot() *Machine {
	var machine Machine
	// initialise registers
	machine.Registers = make(map[uint8]int16, 4)

	// populate registers addresses
	for i := 1; i < 6; i++ {
		machine.Registers[uint8(i)] = 0
	}

	// Mount memory
	machine.Memory = make(map[int16]int16, MemorySize)

	// Generate free memory space
	for i := int16(0); i < MemorySize; i++ {
		machine.Memory[i] = 0
		// check if there is a nop
	}

	return &machine
}

// ShutDown Destroy the machine
func (m *Machine) ShutDown() error {
	// Clear memory contents
	m.Registers = make(map[uint8]int16, 5)
	m.Memory = make(map[int16]int16, MemorySize)
	//Todo: Determine shutdown challenges
	m.Running = false
	return nil
}

func immediate(instruction *util.Instruction) int16 {
	return int16(uint8(instruction.R2<<4 | instruction.R3))
}

func (m *Machine) LoadProgram(path string) error {
	bytecode, err := util.LoadProgram(path)
	if err != nil {
		return err
	}
	if len(bytecode) > MemorySize-ProgramStartAddress {
		return fmt.Errorf("unlimited memory: %dBytes more needed", (len(bytecode)-MemorySize-ProgramStartAddress)*16)
	}
	for i := ProgramStartAddress; i < len(bytecode)+ProgramStartAddress; i++ {
		m.Memory[int16(i)] = bytecode[i-ProgramStartAddress]
		instruction := util.ParseByteCode(bytecode[i-ProgramStartAddress])
		if instruction.O == Nop && (instruction.R3 != 0 || instruction.R1 != 0 || instruction.R2 != 0) {
			immediate := int16(instruction.R1)<<8 | int16(instruction.R2)<<4 | int16(instruction.R3)
			// Store the instruction address into the memory
			m.Memory[immediate] = int16(i)
		}

	}
	m.progStopOffset = int16(len(bytecode) + ProgramStartAddress - 1)
	// Set the PC Register
	m.Registers[PC] = ProgramStartAddress

	return nil
}

func (m *Machine) Log(i *util.Instruction) {
	r1 := m.Registers[R1]
	pc := m.Registers[PC]
	r2 := m.Registers[R2]
	r3 := m.Registers[R3]
	fmt.Printf("%s AT PC = \033[1;34m0x%X\u001B[0m\tR1 = \033[1;33m0x%X\u001B[0m\tR2 =  \u001B[1;36m0x%X\u001B[0m\tR3 =  \u001B[1;35m0x%X\u001B[0m\n", i.String(), pc, r1, r2, r3)
}

func (m *Machine) RunCpu() error {
	defer m.ShutDown()
	m.Running = true

	clockDelay := time.Nanosecond * (1 / CpuFreq * 1_000_000_000)
	for m.Running {
		nextInstructionAddress := m.Registers[PC]

		if nextInstructionAddress < ProgramStartAddress {
			// Force shutdown
			break
		}

		// Get the instruction from the memory
		instructionByteCode := m.Memory[nextInstructionAddress]

		// parse the instruction
		instruction := util.ParseByteCode(instructionByteCode)
		if instruction == nil {
			// fatal error instruction cannot be parsed
			break
		}

		switch instruction.O {
		case Halt:
			m.Halt()
			break
		case Nop:
			m.Nop()
			break
		case Li:
			imm := immediate(instruction)
			m.Li(instruction.R1, imm)
			break
		case Lw:
			m.Lw(instruction.R1, instruction.R2)
			break
		case Sw:
			m.Sw(instruction.R1, instruction.R2)
			break
		case Add:
			m.Add(instruction.R3, instruction.R1, instruction.R2)
			break
		case Sub:
			m.Sub(instruction.R3, instruction.R1, instruction.R2)
			break
		case Mult:
			m.Mult(instruction.R3, instruction.R1, instruction.R2)
			break
		case Div:

			m.Div(instruction.R3, instruction.R1, instruction.R2)
			break
		case J:
			m.J(immediate(instruction))
			break
		case Jr:
			m.Jr(uint8(immediate(instruction)))
			break
		case Bne:
			m.Bne(instruction.R1, instruction.R2, instruction.R3)
			break
		case Beq:
			m.Beq(instruction.R1, instruction.R2, instruction.R3)
			break
		case Inc:
			m.Inc(instruction.R1)
			break
		case Dec:
			m.Dec(instruction.R1)

		}

		m.Log(instruction)

		// Check if there hasn't been any branching and increment PC
		if nextInstructionAddress == m.Registers[PC] {
			// check if we havent reached the end of memory
			if nextInstructionAddress < m.progStopOffset {
				m.Registers[PC]++
			} else {
				break
			}
		}

		time.Sleep(clockDelay)
	}
	return nil
}
