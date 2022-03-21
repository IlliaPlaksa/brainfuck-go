package brainfuck

/*
	Main function of brainfuck API.
Executes BrainFuck code passed as argument.
	Note:
Output of program will be entered to console.
*/
func Interpret(code string) {
	p := compile(code)
	p.Execute()
}

func compile(input string) *program {
	p := newParser()
	result := newProgram(p.parse(input), newMemory())
	return result
}
