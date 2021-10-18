package main

import (
	"fmt"
)

func isNumber(symbol byte) bool {
	if symbol >= '0' && symbol <= '9' {
		return true
	}
	return false
}

func isSpace(symbol byte) bool {
	if symbol == ' ' {
		return true
	}
	return false
}

func isOperation(symbol byte) bool {
	if symbol == '+'  || symbol == '-' || symbol == '*' || symbol == '/'{
		return true
	}
	return false
}

func isBracket(symbol byte) bool {
	if symbol == '(' || symbol == ')' {
		return true
	}
	return false
}

//Переводит строку в различные токены
func createTokens (s string) ([]interface{}, error) {
	result := make([]interface{}, 0)
	for i := 0; i < len(s); i++ {
		var (
			tmp interface{}
			pos	int
			err	error
		)

		symbol := s[i]
		switch {
		case isNumber(symbol):
			tmp, pos, err = atof(s[i:])
			if err != nil {
				return nil, fmt.Errorf("Can't atof %v: %v", s[i:], err)
			}
			i = i + pos
		case isSpace(symbol):
			continue
		case isOperation(symbol):
			tmp, pos, err = GetOperation(s[i:])
			if err != nil {
				return nil, fmt.Errorf("Promlem in GetOperation() with str %s: %v", s, tmp)
			}
			i = i + pos
		case isBracket(symbol):
			tmp, pos, err = GetBraket(s[i:])
			if err != nil {
				return nil, fmt.Errorf("Promlem in GetBraket() with str %s: %v", s, tmp)
			}
			i = i + pos
		default:
			return nil, fmt.Errorf("Can't find symbol %v", s[i])
		}
		result = append(result, tmp)
	}
	return result, nil
}
