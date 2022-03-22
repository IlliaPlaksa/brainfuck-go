package brainfuck

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestInterpret(t *testing.T) {
	type args struct {
		program string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name:     "Hello world!",
			args:     args{"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."},
			expected: "Hello World!\n",
		},

		{
			name:     "Dirty hello world",
			args:     args{"+++vsdfv+++++[>++++s[>++>++a+>+v++>+<a<<<-]>+>d+>->>+[w!<]<-]>>.>--\t-.++++vdf+++..+++.>>.<-.<.++efb +.------.-----\n---.>>+.>++."},
			expected: "Hello World!\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := captureStdout(func() {
				Interpret(tt.args.program)
			})

			if out != tt.expected {
				t.Errorf("Expected %v but have %v", tt.expected, out)
			}
		})
	}
}
