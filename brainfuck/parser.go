package brainfuck

type parser struct {
	stack [][]command
	ptr   int
}

func newParser() *parser {
	return &parser{
		stack: [][]command{{}},
		ptr:   0,
	}
}

func (p *parser) parse(input string) []command {
	for _, char := range input {
		p.parseInstruction(char)
	}
	return p.stack[0]
}

func (p *parser) parseInstruction(char rune) {
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
	}
}
