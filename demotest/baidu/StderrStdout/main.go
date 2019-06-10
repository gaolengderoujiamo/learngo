package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "Hello ")
	fmt.Fprintf(os.Stderr, "World!")
}
