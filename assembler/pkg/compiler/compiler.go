package compiler

import (
	"assembler/pkg/util"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	registerSyntax = "^R[1-3]$"
	memLocSyntax   = "^[a-zA-Z_][a-zA-Z_0-9]*"
	hexSyntax      = "^0x[a-fA-F0-9]*"
	labelSyntax    = "[a-zA-Z_][a-zA-Z0-9_]+:"
)

var (
	OpCodes = map[string]int16{
		"halt": 0x00,
		"nop":  0x01,
		"li":   0x02,
		"lw":   0x03,
		"sw":   0x04,
		"add":  0x05,
		"sub":  0x06,
		"mult": 0x07,
		"div":  0x08,
		"j":    0x09,
		"jr":   0x0A,
		"beq":  0x0B,
		"bne":  0x0C,
		"inc":  0x0D,
		"dec":  0x0E,
	}
	Registers = map[string]int16{
		"R1":   0x01,
		"R2":   0x02,
		"R3":   0x03,
		"PC":   0x04,
		"COND": 0x05,
	}

	symbolTable = SymbolTable{}
)

func genMemOpLiCode(tk util.Token, baseOp *int16) error {
	reg := Registers[tk.Operands[0]]
	// check if imediate is an hexadecimal value
	re := regexp.MustCompile(hexSyntax)
	if re.Match([]byte(tk.Operands[1])) {
		hexValue, err := strconv.ParseInt(tk.Operands[1], 0, 0xF)
		if err != nil {
			return fmt.Errorf("failed to parse %s at \"%s\" in line %d")
		}
		*baseOp = *baseOp | (reg << 8) | int16(hexValue)
	} else {
		address, err := symbolTable.GetAddress(tk.Operands[1])
		if err != nil {
			return err
		}
		*baseOp = *baseOp | (reg << 8) | int16(address)
	}

	return nil
}

func genMemOpLRCode(tk util.Token, baseOp *int16) error {
	reg1 := Registers[tk.Operands[0]]
	reg2 := Registers[tk.Operands[1]]

	*baseOp = *baseOp | (reg1 << 8) | reg2
	return nil
}

func genArithCode(tk util.Token, baseOp *int16) error {

	if len(tk.Operands) < 3 {
		return errors.New(fmt.Sprintf("invalid syntax \"%s\" on line %d", tk.Raw, tk.Line))
	}
	reg1 := Registers[tk.Operands[0]]
	reg2 := Registers[tk.Operands[1]]
	reg3 := Registers[tk.Operands[2]]

	*baseOp = *baseOp | (reg1 << 8) | (reg2 << 4) | reg3
	return nil
}

func genBranchCode(tk util.Token, baseOp *int16) error {

	if len(tk.Operands) < 3 {
		return errors.New(fmt.Sprintf("invalid syntax \"%s\" on line %d", tk.Raw, tk.Line))
	}
	reg1 := Registers[tk.Operands[0]]
	reg2 := Registers[tk.Operands[1]]
	reg3 := Registers[tk.Operands[2]]

	*baseOp = *baseOp | (reg1 << 8) | (reg2 << 4) | reg3
	return nil
}
func genIncCode(tk util.Token, baseOp *int16) error {

	if len(tk.Operands) < 1 {
		return errors.New(fmt.Sprintf("invalid syntax \"%s\" on line %d", tk.Raw, tk.Line))
	}

	reg1 := Registers[tk.Operands[0]]
	*baseOp = *baseOp | reg1<<8
	return nil
}

func genJumpCode(tk util.Token, baseOp *int16) error {

	if len(tk.Operands) < 1 {
		return errors.New(fmt.Sprintf("invalid syntax \"%s\" on line %d", tk.Raw, tk.Line))
	}

	// j 0x00000000
	hex := strings.TrimPrefix(tk.Operands[0], "0x")
	hexValue, _ := strconv.ParseInt(hex, 0xF, 0xF)
	*baseOp = *baseOp | int16(hexValue)
	return nil
}

func genJumpRegisterCode(tk util.Token, baseOp *int16) error {

	if len(tk.Operands) < 1 {
		return errors.New(fmt.Sprintf("invalid syntax \"%s\" on line %d", tk.Raw, tk.Line))
	}
	// j R2
	reg1 := Registers[tk.Operands[0]]
	*baseOp = *baseOp | reg1
	return nil
}

func genLabelCode(tk util.Token, baseOp *int16) error {
	label := strings.TrimSuffix(tk.Raw, ":")
	address, err := symbolTable.GetAddress(label)
	if err != nil {
		return err
	}
	*baseOp = *baseOp | int16(address)
	return nil
}

func genTargetCode(in string) ([]int16, error) {
	byteCode := make([]int16, 0)

	programSource, err := util.Tokenise(in)
	if err != nil {
		return nil, err
	}

	for _, token := range programSource.Tokens {
		var opcode int16
		baseOpCode, found := OpCodes[token.Type.String()]
		if found {
			opcode = baseOpCode << 12
			if baseOpCode == OpCodes["li"] {
				err = genMemOpLiCode(token, &opcode)
			} else if baseOpCode == OpCodes["lw"] ||
				baseOpCode == OpCodes["sw"] {
				err = genMemOpLRCode(token, &opcode)
			} else if baseOpCode == OpCodes["add"] ||
				baseOpCode == OpCodes["sub"] ||
				baseOpCode == OpCodes["div"] ||
				baseOpCode == OpCodes["mult"] {
				err = genArithCode(token, &opcode)
			} else if baseOpCode == OpCodes["bne"] ||
				baseOpCode == OpCodes["beq"] {
				err = genArithCode(token, &opcode)
			} else if baseOpCode == OpCodes["jr"] {
				err = genJumpRegisterCode(token, &opcode)
			} else if baseOpCode == OpCodes["j"] {
				err = genJumpCode(token, &opcode)
			} else if baseOpCode == OpCodes["bne"] ||
				baseOpCode == OpCodes["beq"] {
				err = genBranchCode(token, &opcode)
			} else if baseOpCode == OpCodes["inc"] ||
				baseOpCode == OpCodes["dec"] {
				err = genIncCode(token, &opcode)
			}
		} else if token.Type.String() == "-" {
			opcode = OpCodes["nop"] << 12
			err = genLabelCode(token, &opcode)
		}

		if err != nil {
			return nil, err
		}

		byteCode = append(byteCode, opcode)
	}

	return byteCode, nil
}

func writeFile(bytecode []int16, out string) error {
	// open outfile
	outfile, err := os.Create(out)
	if err != nil {
		return err
	}

	defer outfile.Close()

	enc := gob.NewEncoder(outfile)
	err = enc.Encode(bytecode)
	if err != nil {
		return err
	}
	return nil
}

func Assemble(in, out string) error {
	// Generate targetCode
	byteCode, err := genTargetCode(in)
	if err != nil {
		return err
	}

	err = writeFile(byteCode, out)
	if err != nil {
		return err
	}
	return nil
}
