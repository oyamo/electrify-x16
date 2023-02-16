package util

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

type TokenType string

const (
	// TokenHalt Terminate program
	TokenHalt = TokenType("halt")
	// TokenNop Do nothing
	TokenNop = TokenType("nop")
	// TokenLi Load immediate
	TokenLi = TokenType("li")
	// TokenLw Load word
	TokenLw = TokenType("lw")
	// TokenSw Store Word
	TokenSw    = TokenType("sw")
	TokenAdd   = TokenType("add")
	TokenSub   = TokenType("sub")
	TokenMult  = TokenType("mult")
	TokenDiv   = TokenType("div")
	TokenJump  = TokenType("j")
	TokenUJump = TokenType("jr")
	TokenBeq   = TokenType("beq")
	TokenBne   = TokenType("bne")
	TokenInc   = TokenType("inc")
	TokenDec   = TokenType("dec")
	TokenBlock = TokenType("-")
)

type Token struct {
	Type     TokenType
	Operands []string
	Raw      string
	Line     int
}

type Program struct {
	Tokens []Token
}

func (t TokenType) String() string {
	return string(t)
}

var tokenTypes = map[string]TokenType{
	"halt": TokenHalt,
	"nop":  TokenNop,
	"li":   TokenLi,
	"lw":   TokenLw,
	"sw":   TokenSw,
	"add":  TokenAdd,
	"sub":  TokenSub,
	"mult": TokenMult,
	"div":  TokenDiv,
	"j":    TokenJump,
	"jr":   TokenUJump,
	"beq":  TokenBeq,
	"bne":  TokenBne,
	"inc":  TokenInc,
	"dec":  TokenDec,
}

// readSource
func readSource(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner, err
}

// tokeniseFrom Returns a list of tokens
func tokeniseFrom(scanner *bufio.Scanner) (*Program, error) {
	var program Program

	if scanner == nil {
		return nil, errors.New("scanner nil")
	}
	var lineCount int
	for scanner.Scan() {
		var token Token
		txt := scanner.Text()
		txt = strings.TrimSpace(txt)
		// Increment linecount
		lineCount++

		// Skip  blank lines
		if len(txt) == 0 {
			continue
		}
		parts := strings.Fields(txt)
		if len(parts) > 0 {
			// Skip comments
			if txt[0] == ';' {
				continue
			}
			token.Type = TokenType(parts[0])
		}
		if len(parts) > 1 {
			token.Operands = append(token.Operands, parts[1:]...)
		}

		if _, ok := tokenTypes[token.Type.String()]; !ok {
			re := regexp.MustCompile("[a-zA-Z_][a-zA-Z0-9_]+:")
			if re.Match([]byte(txt)) {
				token.Type = TokenBlock
			}
		}

		token.Raw = txt
		token.Line = lineCount

		program.Tokens = append(program.Tokens, token)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return &program, nil
}

func Tokenise(sourcePath string) (*Program, error) {
	scanner, err := readSource(sourcePath)
	if err != nil {
		return nil, err
	}

	program, err := tokeniseFrom(scanner)

	err = program.ValidateSyntax()
	if err != nil {
		return nil, err
	}
	return program, nil
}
