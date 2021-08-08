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

func isPlus(symbol byte) bool {
	if symbol == '+' {
		return true
	}
	return false
}

//Переводит строку в различные токены
func createTokens (s string) ([]interface{}, error) {
	result := make([]interface{}, 0)
	for i := 0; i < len(s); i++ {
		var tmp interface{}
		symbol := s[i]
		switch {
		case isNumber(symbol):
			tmp1, pos, err := atof(s[i:])
			if err != nil {
				return nil, fmt.Errorf("Can't atof %v: %v", s[i:], err)
			}
			tmp = tmp1
			i = i + pos - 1
		case isSpace(symbol):
			continue
		case isPlus(symbol):
			tmp = "+"
		default:
			return nil, fmt.Errorf("Can't find symbol %v", s[i])
		}
		result = append(result, tmp)
	}
	return nil, nil
}
