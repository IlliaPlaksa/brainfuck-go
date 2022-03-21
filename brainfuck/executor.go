package brainfuck

/*
	Main function of brainfuck API.
Executes BrainFuck program passed as argument.
	Note:
Output of program will be entered to console.
*/
func Interpret(program string) {
	exec := compile(program)
	exec.Execute()
}

func compile(input string) *program {
	p := newParser()
	result := newProgram(p.parse(input), newMemory())
	return result
}

type Parser struct {
	stack [][]command
	ptr   int
}

func newParser() *Parser {
	return &Parser{
		stack: [][]command{{}},
		ptr:   0,
	}
}

func (p *Parser) parse(input string) []command {
	for _, char := range input {
		p.parseInstruction(char)
	}
	return p.stack[0]
}

func (p *Parser) parseInstruction(char rune) {
	switch char {
	case '+':
		p.stack[p.ptr] = append(p.stack[p.ptr], Increment{})
	case '-':
		p.stack[p.ptr] = append(p.stack[p.ptr], Decrement{})
	case '>':
		p.stack[p.ptr] = append(p.stack[p.ptr], MoveForward{})
	case '<':
		p.stack[p.ptr] = append(p.stack[p.ptr], MoveBackward{})
	case '.':
		p.stack[p.ptr] = append(p.stack[p.ptr], Print{})
	case '[':
		p.stack = append(p.stack, []command{})
		p.ptr++
	case ']':
		p.stack[p.ptr-1] = append(p.stack[p.ptr-1], Loop{p.stack[p.ptr]})
		p.stack = p.stack[:p.ptr]
		p.ptr--
	default:
	}
}
