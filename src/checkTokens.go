package main

import (
	"fmt"
)

func CheckTokens(tokens []string, vars map[string]interface{}) int {
	for i := 0; i < len(tokens); i++ {
		token, err := ConvToken(tokens[i:], vars)
		if err == nil {
			fmt.Println(token)
		}
	}
	return 0
}
