package brainfuck

import (
	"reflect"
	"testing"
)

func TestLoop_execute(t *testing.T) {
	input := newMemory()
	input.buffer[0] = 1

	expected := newMemory()
	expected.buffer = append(expected.buffer, 2)

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
			fields{[]command{MoveForward{}, Increment{}}},
			args{memory: newMemory()},
			*newMemory(),
		},
		{
			"executes all commands if current cell value is non-zero",
			fields{[]command{
				MoveForward{}, Increment{}, Increment{}, MoveBackward{}, Decrement{},
			}},
			args{memory: input},
			*expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Loop{commands: tt.fields.commands}
			l.Execute(tt.args.memory)
			if !reflect.DeepEqual(*tt.args.memory, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, *tt.args.memory)
			}
		})
	}
}
