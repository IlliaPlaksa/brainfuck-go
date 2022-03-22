package brainfuck

import "strings"

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

var repeatedInstructions = map[string]rune{
	"+++++": '$',
}

func replaceRepeatedInstructions(program string) string {
	optimized := program

	for seq, symbol := range repeatedInstructions {
		optimized = strings.Replace(optimized, seq, string(symbol), -1)
	}

	return optimized
}

func (p *parser) parseInstruction(char rune) {
	switch char {
	case '+':
		p.stack[p.ptr] = append(p.stack[p.ptr], increment{})
	case '-':
		p.stack[p.ptr] = append(p.stack[p.ptr], decrement{})
	case '>':
		p.stack[p.ptr] = append(p.stack[p.ptr], moveForward{})
	case '<':
		p.stack[p.ptr] = append(p.stack[p.ptr], moveBackward{})
	case '.':
		p.stack[p.ptr] = append(p.stack[p.ptr], output{})
	case '[':
		p.stack = append(p.stack, []command{})
		p.ptr++
	case ']':
		p.stack[p.ptr-1] = append(p.stack[p.ptr-1], loop{p.stack[p.ptr]})
		p.stack = p.stack[:p.ptr]
		p.ptr--
	}
}
