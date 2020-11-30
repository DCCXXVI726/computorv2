package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			err = fmt.Errorf("problem in ReadString: %s", err)
			panic(err)
		}

		if text == ("exit\n") {
			fmt.Println("See you next time!")
			os.Exit(0)
		}
		fmt.Print("->")
		fmt.Print(text)
	}
}
