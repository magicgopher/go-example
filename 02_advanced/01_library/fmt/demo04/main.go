package main

import (
	"fmt"
	"os"
)

func main() {
	// Fprint
	fmt.Fprint(os.Stdout, "Hello", "World") // 输出到标准输出: Hello World

	// Fprintln
	fmt.Fprintln(os.Stdout, "Hello", "World") // 输出到标准输出: Hello World\n

	// Fprintf
	fmt.Fprintf(os.Stdout, "Name: %s, Age: %d\n", "Alice", 25) // 输出到标准输出: Name: Alice, Age: 25
}
