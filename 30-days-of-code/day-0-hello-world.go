package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')
	fmt.Println("Hello, World.")
	fmt.Println(line)
}
