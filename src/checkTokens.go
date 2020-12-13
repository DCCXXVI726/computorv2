package main

import (
	"fmt"
)

func CheckTokens(tokens []string, vars map[string]interface{}) (string, error) {
	newTokens := make([]interface{}, 0)
	for i := 0; i < len(tokens); i++ {
		token, err := ConvToken(tokens[i:], vars)
		if err != nil {
			err = fmt.Errorf("can't conv token %s: %s", token, err)
			return "", err
		}
		newTokens = append(newTokens, token)
	}
	number, err := Culc(newTokens)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(number)
	return "", nil
}
