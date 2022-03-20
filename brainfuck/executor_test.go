package brainfuck

import "testing"

func TestInterpret(t *testing.T) {
	type args struct {
		program string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Hello, world!", args{"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."}},
		{"Dirty hello world", args{"+++vsdfv+++++[>++++s[>++>++a+>+v++>+<a<<<-]>+>d+>->>+[w!<]<-]>>.>--\t-.++++vdf+++..+++.>>.<-.<.++efb +.------.-----\n---.>>+.>++."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interpret(tt.args.program)
		})
	}
}
