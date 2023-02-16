package machine

import "testing"

func TestMachine_LoadProgram(t *testing.T) {
	machine := Boot()
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		machine *Machine
		args    args
		wantErr bool
	}{
		{
			name:    "Load simple Program",
			machine: machine,
			args: args{
				path: "/home/oyamo/GolandProjects/pp-assembler/machine/example/loop.o",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.machine.LoadProgram(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("LoadProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMachine_RunCpu(t *testing.T) {
	machine := Boot()
	err := machine.LoadProgram("/home/oyamo/GolandProjects/pp-assembler/machine/example/loop.o")
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name    string
		fields  *Machine
		wantErr bool
	}{
		{
			name:    "Execute simple program",
			fields:  machine,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.RunCpu(); (err != nil) != tt.wantErr {
				t.Errorf("RunCpu() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
