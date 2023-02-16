package util

import "testing"

func TestProgram_ValidateSemantics(t *testing.T) {
	type fields struct {
		Tokens []Token
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				Tokens: tt.fields.Tokens,
			}
			if err := p.ValidateSemantics(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSemantics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
