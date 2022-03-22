package brainfuck

type memory struct {
	buffer []byte
	ptr    int
}

func newMemory(bytes []byte, pointer int) *memory {
	return &memory{
		buffer: bytes,
		ptr:    pointer,
	}
}
