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
			fmt.Print("bay bay")
			break
		}
		i := 0
		find_equal := false
		for i = 0; i < len(str); i++ {
			if str[i] == '=' {
				if i != len(str) -1 {
					find_equal = true
					break
				}
			}
		}
		var str1, str2 string
		if find_equal {
			 str1, _ = checkVar(str[:i])
			 str2 = str[i+1:]
		} else {
			str2 = str
		}
		tokens, err := createTokens(str2)
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
		if find_equal {
			setVar(str1, result)
		}
		printResult(result)
		fmt.Println()
	}
}
