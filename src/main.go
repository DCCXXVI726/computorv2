package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		str, _ := reader.ReadString('\n')

		str = strings.Replace(str,"\n", "", -1)
		if str == "exit" {
			fmt.Errorf("bay bay")
			break
		}
		tokens, err := createTokens(str)
		if err != nil {
			err = fmt.Errorf("Promlem in createTokens: %v", err)
			fmt.Println(err)
			continue
		}
		result, err := culc(tokens)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
