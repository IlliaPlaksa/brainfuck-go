package brainfuck

/*Interpret executes BrainFuck code passed as argument.
	Note:
Output of program will be entered to console.
*/
func Interpret(code string) {
	p := compile(code)
	p.execute()
}

func compile(input string) *program {
	p := newParser()
	commands := p.parse(input)
	program := newProgram(commands, newMemory([]byte{0}, 0))
	return program
}
