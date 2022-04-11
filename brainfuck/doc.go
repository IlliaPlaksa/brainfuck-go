/*
Package brainfuck provides a functionality of interpreter of BrainFuck programming language.

Available instructions:

base: "+", "-", ">", "<", "[", "]", ".".

custom: "$" - increment by five

	Note:
Characters in input BrainFuck program that does not match any available instruction will be ignored.

	Example usage:
		"Hello, world!" program in BrainFuck
		brainfuck.Interpret("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
*/
package brainfuck
