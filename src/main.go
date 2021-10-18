package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello there.")
	printUsage()
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
		result, err, flag, _ := culc(tokens)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if flag {
			fmt.Println("end with close bracket")
			continue
		}
		fmt.Println(result)
	}
}
