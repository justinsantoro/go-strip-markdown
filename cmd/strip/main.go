package main

import (
	"fmt"
	"os"

	stripmd "github.com/justinsantoro/go-strip-markdown"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	fmt.Print(stripmd.Strip(os.Args[1]))
}
