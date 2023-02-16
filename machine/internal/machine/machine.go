package machine

const MemorySize = 2 << 15

type Machine struct {
	Registers map[uint8]int16
	Memory   map[int16]uint8
	Halted   bool
}

// Memory Layout
// 0x00 - 0x0 0x0000CFFF

func Boot() *Machine {
	var machine Machine
	// initialise registers
	machine.Registers = make(map[uint8]int16, 4)

	// populate registers addresses
	for i := 0; i < 4; i++ {
		machine.Registers[uint8(i)] = 0
	}

	// Mount memory
	machine.Memory = make(map[uint8]uint8, MemorySize)

	for
}
