package machine

import "testing"

func TestMachine_LoadProgram(t *testing.T) {
	type fields struct {
		Registers map[uint8]int16
		Memory    map[int16]int16
		Halted    bool
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Machine{
				Registers: tt.fields.Registers,
				Memory:    tt.fields.Memory,
				Halted:    tt.fields.Halted,
			}
			if err := m.LoadProgram(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("LoadProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
