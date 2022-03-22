package brainfuck

import (
	"reflect"
	"testing"
)

func TestLoop_execute(t *testing.T) {
	type fields struct {
		commands []command
	}
	type args struct {
		memory *memory
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected memory
	}{
		{
			"does not execute commands if current cell equals to zero",
			fields{[]command{moveForward{}, increment{}}},
			args{memory: newMemory([]byte{0}, 0)},
			*newMemory([]byte{0}, 0),
		},
		{
			"executes all commands if current cell value is non-zero",
			fields{[]command{
				moveForward{}, increment{}, increment{}, moveBackward{}, decrement{},
			}},
			args{memory: newMemory([]byte{1}, 0)},
			*newMemory([]byte{0, 2}, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := loop{commands: tt.fields.commands}
			l.Execute(tt.args.memory)
			if !reflect.DeepEqual(*tt.args.memory, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, *tt.args.memory)
			}
		})
	}
}
