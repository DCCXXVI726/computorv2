package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		tokens []string
	)
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
		} else {
			tokens, err = CreateTokens(text)
			if err != nil {
				err = fmt.Errorf("problem with CreateTokens: %s", err)
				panic(err)
			}
		}
		fmt.Print("->")
		fmt.Println(tokens)
	}
}
