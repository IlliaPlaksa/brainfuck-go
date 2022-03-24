package brainfuck

// Interpret executes BrainFuck code passed as argument.
//	Note:
// Output of program will be entered to console.
func Interpret(input string) {
	p := compile(input)
	p.execute()
}

func compile(input string) *program {
	p := newParser()
	commands := p.parse(input)
	mem := newMemory([]byte{0}, 0)
	return newProgram(commands, mem)
}
