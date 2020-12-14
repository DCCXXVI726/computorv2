package main

import (
	"fmt"
)

// check symbol in string
func find(symbols string, ch byte) bool {
	for _, v := range symbols {
		if byte(v) == ch {
			return true
		}
	}
	return false
}

//CreateTokens split text to the tokens
func CreateTokens(text string) ([]string, error) {
	tokens := make([]string, 0)
	i := 0
	symbols := "-+*/%^()[]i,?="
	for i < len(text)-1 {
		t := text[i]
		switch {
		case find(symbols, t):
			tokens = append(tokens, text[i:i+1])
			i++
		case t >= '0' && t <= '9':
			k := i + 1
			for text[k] >= '0' && text[k] <= '9' && k < len(text)-1 {
				k++
			}
			tokens = append(tokens, text[i:k])
			i = k
		case IsAlpha(t):
			k := i + 1
			for IsAlpha(text[k]) && k < len(text)-1 {
				k++
			}
			tokens = append(tokens, text[i:k])
			i = k
		case t == ' ':
			i++
		default:
			err := fmt.Errorf("неизвестный знак - %c на позиции %d", text[i], i)
			return nil, err
		}
	}
	return tokens, nil
}
