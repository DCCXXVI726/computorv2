package main

import (
	"fmt"
)

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
		case t >= '0' && t <= '9':
			k := i + 1
			for text[k] >= '0' && text[k] <= '9' && k < len(text)-1 {
				k++
			}
			tokens = append(tokens, text[i:k])
			i = k - 1
		default:
			fmt.Printf("неизвестный знак - %c на позиции %d\n", text[i], i)
		}
		i++
	}
	return tokens, nil
}
