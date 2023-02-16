package util

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	haltSyntax         = "^halt$"
	nopSyntax          = "^nop$"
	liSyntax           = "^li (R[1-3]|PC|COND) (0x[0-9A-Fa-f]{1,8}|[a-zA-Z_][a-zA-Z0-9_]*)$"
	lwSyntax           = "^lw (R[1-3]|PC|COND) (PC|COND|R[1-3])$"
	swSyntax           = "^sw (R[1-3]|PC|COND) (PC|COND|R[1-3])$"
	addSyntax          = "^add (R[1-3]|PC|COND) (PC|COND|R[1-3]) (PC|COND|R[1-3])$"
	subSyntax          = "^sub (R[1-3]|PC|COND) (PC|COND|R[1-3]) (PC|COND|R[1-3])$"
	multSyntax         = "^mult (R[1-3]|PC|COND) (PC|COND|R[1-3]) (PC|COND|R[1-3])$"
	divSyntax          = "^div (R[1-3]|PC|COND) (PC|COND|R[1-3]) (PC|COND|R[1-3])$"
	jumpSyntax         = "^j 0x[0-9A-Fa-f]{,8}$"
	jumpRegisterSyntax = "^jr (COND|PC|R[1-3])$"
	beqSyntax          = "^beq (PC|COND|R[1-3]) (PC|COND|R[1-3]) (R[1-3]|PC|COND)$"
	bneSyntax          = "^bne (PC|COND|R[1-3]) (PC|COND|R[1-3]) (R[1-3]|PC|COND)$"
	incSyntax          = "^inc (COND|PC|R[1-3])$"
	decSyntax          = "^dec (COND|PC|R[1-3])$"
	blockSyntax        = "[a-zA-Z_][a-zA-Z0-9_]+:"
)

// ValidateSyntax Goes through a program and halts When an error has been found
func (p *Program) ValidateSyntax() error {
	for _, token := range p.Tokens {
		var re *regexp.Regexp
		switch token.Type {
		case TokenHalt:
			re = regexp.MustCompile(haltSyntax)
			break
		case TokenNop:
			re = regexp.MustCompile(nopSyntax)
			break
		case TokenLi:
			re = regexp.MustCompile(liSyntax)
			break
		case TokenLw:
			re = regexp.MustCompile(lwSyntax)
			break
		case TokenSw:
			re = regexp.MustCompile(swSyntax)
			break
		case TokenAdd:
			re = regexp.MustCompile(addSyntax)
			break
		case TokenSub:
			re = regexp.MustCompile(subSyntax)
			break
		case TokenMult:
			re = regexp.MustCompile(multSyntax)
			break
		case TokenDiv:
			re = regexp.MustCompile(divSyntax)
			break
		case TokenJump:
			re = regexp.MustCompile(jumpSyntax)
			break
		case TokenUJump:
			re = regexp.MustCompile(jumpRegisterSyntax)
			break
		case TokenBeq:
			re = regexp.MustCompile(beqSyntax)
			break
		case TokenBne:
			re = regexp.MustCompile(bneSyntax)
			break
		case TokenInc:
			re = regexp.MustCompile(incSyntax)
			break
		case TokenDec:
			re = regexp.MustCompile(decSyntax)
			break
		case TokenBlock:
			re = regexp.MustCompile(blockSyntax)
		}

		if re == nil {
			return errors.New(fmt.Sprintf("invalid syntax \"%s\" on line %d", token.Raw, token.Line))
		}

		if !re.MatchString(token.Raw) {
			return errors.New(fmt.Sprintf("invalid syntax \"%s\" on line %d", token.Raw, token.Line))
		}
	}

	return nil
}
