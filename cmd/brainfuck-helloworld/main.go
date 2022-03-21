package main

import "brainfuck/brainfuck"

func main() {
	program := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	brainfuck.Interpret(program)
}
