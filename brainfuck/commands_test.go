package brainfuck

import (
	"reflect"
	"testing"
)

func TestSimpleCommands_Execute(t *testing.T) {
	type args struct {
		memory *memory
	}
	tests := []struct {
		name     string
		command  command
		args     args
		expected memory
	}{
		{
			name:     "should increment buffer's cell with index [0] by one",
			command:  increment{},
			args:     args{memory: newMemory([]byte{0}, 0)},
			expected: *newMemory([]byte{1}, 0),
		},
		{
			name:     "should decrement buffer's cell with index [0] by one",
			command:  decrement{},
			args:     args{memory: newMemory([]byte{3}, 0)},
			expected: *newMemory([]byte{2}, 0),
		},
		{
			name:     "should move memory pointer to the right",
			command:  moveForward{},
			args:     args{memory: newMemory([]byte{0}, 0)},
			expected: *newMemory([]byte{0, 0}, 1),
		},
		{
			name:     "should move memory pointer to the left",
			command:  moveBackward{},
			args:     args{memory: newMemory([]byte{0, 0, 0}, 2)},
			expected: *newMemory([]byte{0, 0, 0}, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.command.Execute(tt.args.memory)
			if !reflect.DeepEqual(*tt.args.memory, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, *tt.args.memory)
			}
		})
	}
}

func TestOutput_Execute(t *testing.T) {
	type args struct {
		memory *memory
	}
	tests := []struct {
		name     string
		command  command
		args     args
		expected string
	}{
		{
			name:     "should output letter \"A\" to console",
			command:  output{},
			args:     args{memory: newMemory([]byte{65}, 0)},
			expected: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := captureStdout(func() {
				tt.command.Execute(tt.args.memory)
			})

			if !reflect.DeepEqual(out, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, out)
			}
		})
	}
}

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
			name:     "does not execute commands if current cell equals to zero",
			fields:   fields{[]command{moveForward{}, increment{}}},
			args:     args{memory: newMemory([]byte{0}, 0)},
			expected: *newMemory([]byte{0}, 0),
		},
		{
			name:     "executes all commands if current cell value is non-zero",
			fields:   fields{[]command{moveForward{}, increment{}, increment{}, moveBackward{}, decrement{}}},
			args:     args{memory: newMemory([]byte{1}, 0)},
			expected: *newMemory([]byte{0, 2}, 0),
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
