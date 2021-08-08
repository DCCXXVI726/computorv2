package main

import(
	"fmt"
)

func atof(str string) (float64, int, error) {
	dot := false
	result := 0.0
	k := 1
	i := 0
	for ;i < len(str); i++ {
		symbol := str[i]
		switch {
		case symbol == '.':
			if dot {
				return 0.0, 0, fmt.Errorf("double dot in %s", str)
			}
			dot = true

		case symbol >= '0' && symbol <= '9':
			if dot {
				tmp := float64(symbol - '0')
				for l := k; l > 0; l-- {
					tmp = tmp / 10
				}
				result = result + tmp
				k++
			} else {
				result = result * 10 + float64(symbol - '0')
			}
		default:
			return 0.0, 0, fmt.Errorf("can't find symbol %v in %s", symbol, str)
		}
	}
	return result, i, nil
}
