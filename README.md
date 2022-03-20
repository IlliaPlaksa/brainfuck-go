# brainfuck-go

`BrainFuck` programming language interpreter implementation on Golang

##How to use
Package `brainfuck` can be used by calling `brainfuck.Interpret` method with BrainFuck program as input parameter
```bash
import "brainfuck-go/brainfuck"

func main() {
	program := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	brainfuck.Interpret(program)
}
```
This program will print "Hello, world!" to console

