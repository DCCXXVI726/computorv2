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
	if len(newTokens) > 2 {
		switch newTokens[0].(type) {
		case string:
			// добавить проверку на все знаки
			if newTokens[0].(string) != "=" {
				switch newTokens[1].(type) {
				case string:
					if newTokens[1].(string) == "=" {
						value, err := Culc(newTokens[2:])
						if err != nil {
							err = fmt.Errorf("problem with culc vars: %s", err)
							return "", err
						}
						vars[newTokens[0].(string)] = value
						return value.(Complex).toString(), nil

					}
				}
			}
		}
	}
	number, err := Culc(newTokens)
	if err != nil {
		err = fmt.Errorf("problem with Culc: %s", err)
		return "", err
	}

	return number.(Complex).toString(), nil
}
