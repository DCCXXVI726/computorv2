package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		fmt.Print(text)
	}
}
