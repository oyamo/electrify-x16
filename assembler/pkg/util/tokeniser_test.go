package util

import (
	"reflect"
	"testing"
)

func TestTokenise(t *testing.T) {
	type args struct {
		sourcePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Program
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenise(tt.args.sourcePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tokenise() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenise() got = %v, want %v", got, tt.want)
			}
		})
	}
}
