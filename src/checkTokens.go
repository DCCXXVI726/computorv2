package main

import (
	"fmt"
)

// CheckTokens convert cycle tokens and do main culc
func CheckTokens(tokens []string, vars map[string]interface{}) (string, error) {
	newTokens := make([]interface{}, 0)
	for i := 0; i < len(tokens); i++ {
		token, err := ConvToken(tokens[i:], vars)
		if err != nil {
			err = fmt.Errorf("problem with ConvToken: %s", err)
			return "", err
		}
		newTokens = append(newTokens, token)
	}

	number, err := Culc(newTokens)
	if err != nil {
		err = fmt.Errorf("problem with Culc: %s", err)
		return "", err
	}

	return number.(Complex).toString(), nil
}
