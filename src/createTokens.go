package main

import (
	"fmt"
)

func isMatrix(symbol byte) bool {
	if symbol == '[' {
		return true
	}
	return false
}
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
	if symbol == '+'  || symbol == '-' || symbol == '*' || symbol == '/' || symbol == '%' || symbol == '^'{
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

func isVar(str string) bool {
	if len(str) > 0 {
		i := 0
		if (str[0] == 'i' || str[0] == 'I') && len(str) > 1 {
			i = 1
		}
		if (str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z') {
			return true
		}
		return false
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
			if i + 1 < len(s) && s[i+1] == 'i' {
				tmp = Complex{0, tmp.(float64)}
				i = i + 1
			}
		case isSpace(symbol):
			continue
		case isOperation(symbol):
			tmp, pos, err = GetOperation(s[i:])
			if err != nil {
				return nil, fmt.Errorf("Promlem in GetOperation() with str %s: %v", s, err)
			}
			i = i + pos
		case isBracket(symbol):
			tmp, pos, err = GetBraket(s[i:])
			if err != nil {
				return nil, fmt.Errorf("Promlem in GetBraket() with str %s: %v", s, err)
			}
			i = i + pos
		case isMatrix(symbol):
			tmp, pos, err = GetMatrix(s[i:])
			if err != nil {
				return nil, fmt.Errorf("Promlem in GetBraket() with str %s: %v", s, err)
			}
			i = i + pos
		case isVar(s[i:]):
			var variable string
			variable, pos = checkVar(s[i:])
			tmp, err = getVar(variable)
			if err != nil {
				return nil, fmt.Errorf("Can't find variable: %s",err)
			}
			i = i + pos
		case symbol == 'i':
			tmp = Complex{0, 1}
		default:
			return nil, fmt.Errorf("Can't find symbol %c with code %v", s[i], s[i])
		}
		result = append(result, tmp)
	}
	return result, nil
}
