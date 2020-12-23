package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
**	Read line
**	split on numbers/chars
**	main culc
 */

func main() {
	var (
		tokens []string
		vars   map[string]interface{}
	)
	vars = make(map[string]interface{}, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			err = fmt.Errorf("problem in ReadString: %s", err)
			panic(err)
		}
		fmt.Print("->")
		if text == ("exit\n") {
			fmt.Println("See you next time!")
			os.Exit(0)
		} else {
			tokens, err = CreateTokens(text)
			if err != nil {
				err = fmt.Errorf("problem with CreateTokens: %s", err)
				fmt.Println(err)
				continue
			}
		}
		msg, err := CheckTokens(tokens, vars)
		if err != nil {
			err = fmt.Errorf("problem with CheckTokens: %s", err)
			fmt.Println(err)
			continue
		}
		fmt.Println(msg)
	}
}
