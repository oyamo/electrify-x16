package util

import "testing"

func TestProgram_ValidateSyntax(t *testing.T) {
	type fields struct {
		Tokens []Token
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				Tokens: tt.fields.Tokens,
			}
			if err := p.ValidateSyntax(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSyntax() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
