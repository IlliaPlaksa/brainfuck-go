package main

import (
	"github.com/IlliaPlaksa/brainfuck-go/brainfuck"
)

func main() {
	program := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	brainfuck.Interpret(program)
}
