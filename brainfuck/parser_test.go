package brainfuck

import (
	"reflect"
	"testing"
)

func TestParser_parse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		expected []command
	}{
		{
			name: "should parse proper input",
			args: args{"[++++++-><.$]"},
			expected: []command{
				loop{
					[]command{
						incrementByFive{}, increment{}, decrement{}, moveForward{}, moveBackward{}, output{}, incrementByFive{},
					},
				},
			},
		},
		{
			name:     "should skip third party symbols in input",
			args:     args{"+$#-@><)."},
			expected: []command{increment{}, incrementByFive{}, decrement{}, moveForward{}, moveBackward{}, output{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newParser()
			commands := p.parse(tt.args.input)
			if !reflect.DeepEqual(commands, tt.expected) {
				t.Errorf("Expected %v \n but have %v", tt.expected, commands)
			}
		})
	}
}
