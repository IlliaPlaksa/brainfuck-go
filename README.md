# brainfuck-go

[`BrainFuck`](https://en.wikipedia.org/wiki/Brainfuck) programming language interpreter implementation on Golang
## How to install

### Before installation
You should have installed [`go1.17`](https://go.dev/dl/) or higher

### Installation from terminal
```bash
go get github.com/IlliaPlaksa/brainfuck-go
```

## How to use
Package `brainfuck` can be used by calling `brainfuck.Interpret` method with BrainFuck program as input parameter

```go
package main

import (
	"github.com/IlliaPlaksa/brainfuck-go/brainfuck"
)

func main() {
	program := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	brainfuck.Interpret(program)
}
```

This program will print "Hello, world!" to console

## Running tests
To run simple tests from terminal
```bash
cd brainfuck
go test
```

To run tests with code coverage
```bash
go test -cover
```