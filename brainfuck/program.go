package brainfuck

type program struct {
	commands []command
	memory   *memory
}

func newProgram(c []command, mem *memory) *program {
	return &program{
		commands: c,
		memory:   mem,
	}
}

func (p program) execute() {
	for _, cmd := range p.commands {
		cmd.execute(p.memory)
	}
}
