package brainfuck

import "fmt"

type command interface {
	execute(mem *memory)
}

//	Moves pointer in memory struct to right
type moveForward struct{}

func (moveForward) execute(mem *memory) {
	if mem.ptr == len(mem.buffer)-1 {
		mem.buffer = append(mem.buffer, 0)
	}
	mem.ptr++
}

// Moves pointer in memory struct to left
type moveBackward struct{}

func (moveBackward) execute(mem *memory) {
	mem.ptr--
}

// Increments current memory buffer's cell by one
type increment struct{}

func (increment) execute(mem *memory) {
	mem.buffer[mem.ptr]++
}

//	Decrements current memory buffer's cell by one
type decrement struct{}

func (decrement) execute(mem *memory) {
	mem.buffer[mem.ptr]--
}

// Prints value of current memory buffer's cell to console
type output struct{}

func (output) execute(mem *memory) {
	fmt.Print(string(mem.buffer[mem.ptr]))
}

// Executes commands in loop until memory buffer's current cell is non-zero
type loop struct {
	commands []command
}

func (l loop) execute(mem *memory) {
	for mem.buffer[mem.ptr] != 0 {
		for _, cmd := range l.commands {
			cmd.execute(mem)
		}
	}
}
