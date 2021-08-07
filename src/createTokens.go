package main

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

//Переводит строку в различные токены
func createTokens (s string) ([]interface{}, error) {
	for i := 0; i < len(s); i++ {
		symbol := s[i]
		switch {
		case isNumber(symbol):
			break
		case isSpace(symbol):
			continue
		}
	}
	return nil, nil
}
