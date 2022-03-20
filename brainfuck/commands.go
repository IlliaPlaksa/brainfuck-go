package brainfuck

import "fmt"

type command interface {
	Execute(mem *memory)
}

type MoveForward struct{}

func (MoveForward) Execute(mem *memory) {
	if mem.ptr == len(mem.buffer)-1 {
		mem.buffer = append(mem.buffer, 0)
	}
	mem.ptr++
}

type MoveBackward struct{}

func (MoveBackward) Execute(mem *memory) {
	mem.ptr--
}

type Increment struct{}

func (Increment) Execute(mem *memory) {
	mem.buffer[mem.ptr]++
}

type Decrement struct{}

func (Decrement) Execute(mem *memory) {
	mem.buffer[mem.ptr]--
}

type Print struct{}

func (Print) Execute(mem *memory) {
	fmt.Print(string(mem.buffer[mem.ptr]))
}

type Loop struct {
	commands []command
}

func (l Loop) Execute(mem *memory) {
	for mem.buffer[mem.ptr] != 0 {
		for _, cmd := range l.commands {
			cmd.Execute(mem)
		}
	}
}
