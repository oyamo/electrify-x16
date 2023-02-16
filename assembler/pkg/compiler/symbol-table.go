package compiler

import "errors"

const (
	AddrSpace uint16 = 0xFFFF
)

// SymbolTable Maps labels to addresses
type SymbolTable struct {
	symbols     map[string]uint16
	lastAddress uint16
}

func (s *SymbolTable) GetAddress(sym string) (uint16, error) {
	if s.symbols == nil {
		s.symbols = make(map[string]uint16)
	}
	if s.lastAddress == AddrSpace-1 {
		return 0xFFFF, errors.New("address space exceeded")
	}
	addr, found := s.symbols[sym]
	if found {
		return addr, nil
	}

	s.lastAddress++
	s.symbols[sym] = s.lastAddress

	return s.lastAddress, nil
}
