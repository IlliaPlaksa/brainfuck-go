package brainfuck

type memory struct {
	buffer []byte
	ptr    int
}

func newMemory() *memory {
	return &memory{
		buffer: []byte{0},
		ptr:    0,
	}
}
