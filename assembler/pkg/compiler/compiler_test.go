package compiler

import (
	"reflect"
	"testing"
)

func TestAssemble(t *testing.T) {
	type args struct {
		in  string
		out string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Simple Assembly file",
			args: args{
				in:  "/home/oyamo/GolandProjects/pp-assembler/assembler/example/simple.s",
				out: "buffer.o",
			},
			wantErr: false,
		}, {
			name: "Assembly file with loop",
			args: args{
				in:  "/home/oyamo/GolandProjects/pp-assembler/assembler/example/loop.s",
				out: "buffer.o",
			},
			wantErr: false,
		}, {
			name: "Assembly file Syntax error 1",
			args: args{
				in:  "/home/oyamo/GolandProjects/pp-assembler/assembler/example/syntax-err-1.s",
				out: "buffer.o",
			},
			wantErr: true,
		}, {
			name: "Assembly file Syntax error 2",
			args: args{
				in:  "/home/oyamo/GolandProjects/pp-assembler/assembler/example/syntax-err-3.s",
				out: "buffer.o",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Assemble(tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Assemble() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_genTargetCode(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    []int16
		wantErr bool
	}{
		{
			name: "Generate object code 1",
			args: args{
				in: "/home/oyamo/GolandProjects/pp-assembler/assembler/example/simple.s",
			},
			wantErr: false,
			want: []int16{
				0x2100, // 8448
				0x2100, // 8448
				0x2211, // 8721
				0x2212, // 8722
				0x2218, // 8728
				0x21A3, // 8611
				0x2400, // 9216
				0x2500, // 9472
				0x0000, // 0
			},
		}, {
			name: "Generate object code 2",
			args: args{
				in: "/home/oyamo/GolandProjects/pp-assembler/assembler/example/simple-2.s",
			},
			wantErr: false,
			want:    []int16{8465, 8696, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genTargetCode(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("genTargetCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genTargetCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
