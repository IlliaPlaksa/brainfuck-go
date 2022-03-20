package main

import "brainfuck-go/brainfuck"

func main() {
	program := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	brainfuck.Interpret(program)
}
