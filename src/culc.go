package main

import (
	"fmt"
)

func Culc(tokens []interface{}) (interface{}, error) {
	var numbers []interface{}
	var opers []Operation
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].(type) {
		case Complex:
			numbers = append(numbers, tokens[i])
		case Operation:
			if len(opers) != 0 && Priority(tokens[i].(Operation), opers[len(opers)-1]) {
				if len(numbers) > 1 {
					oper := opers[len(opers)-1].getFunc()
					opers = opers[:len(opers)-1]
					newNumber, _ := oper(numbers[len(numbers)-2], numbers[len(numbers)-1])
					numbers = numbers[:len(numbers)-2]
					numbers = append(numbers, newNumber)
				}
			}
			opers = append(opers, tokens[i].(Operation))
		}
	}
	for k := len(opers) - 1; k >= 0; k-- {
		if len(numbers) > 1 {
			oper := opers[k].getFunc()
			opers = opers[:k]
			newNumber, _ := oper(numbers[len(numbers)-2], numbers[len(numbers)-1])
			numbers = numbers[:len(numbers)-2]
			numbers = append(numbers, newNumber)
		}
	}
	if len(numbers) == 1 {
		return numbers[0], nil
	}
	err := fmt.Errorf("Что-то пошло не так")
	return nil, err
}
