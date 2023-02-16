package util

import (
	"testing"
)

func Test_readFile(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantErr bool
	}{
		{
			name: "Ok file",
			args: args{
				in: "/home/oyamo/GolandProjects/pp-assembler/machine/example/simple.o",
			},
			wantNil: false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) != tt.wantNil {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
