package brainfuck

/*Interpret executes BrainFuck code passed as argument.
	Note:
Output of program will be entered to console.
*/
func Interpret(code string) {
	p := compile(code)
	p.Execute()
}

func compile(input string) *program {
	p := newParser()
	commands := p.parse(input)
	program := newProgram(commands, newMemory())
	return program
}
