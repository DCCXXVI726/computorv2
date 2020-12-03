package main

import (
	"fmt"
)

func isAlpha(ch byte) bool {
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}

func Find(symbols string, ch byte) bool {
	for _, v := range symbols {
		if byte(v) == ch {
			return true
		}
	}
	return false
}

func CreateTokens(text string) ([]string, error) {
	tokens := make([]string, 0)
	i := 0
	symbols := "-+*/%^()[]i,?="
	for i < len(text)-1 {
		t := text[i]
		switch {
		case Find(symbols, t):
			tokens = append(tokens, text[i:i+1])
			i++
		case t >= '0' && t <= '9':
			k := i + 1
			for text[k] >= '0' && text[k] <= '9' && k < len(text)-1 {
				k++
			}
			tokens = append(tokens, text[i:k])
			i = k
		case isAlpha(t):
			k := i + 1
			for isAlpha(text[k]) && k < len(text)-1 {
				k++
			}
			tokens = append(tokens, text[i:k])
			i = k
		case t == ' ':
			i++
		default:
			err := fmt.Errorf("неизвестный знак - %c на позиции %d\n", text[i], i)
			return nil, err
		}
		i++
	}
	return tokens, nil
}
