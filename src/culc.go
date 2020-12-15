package main

import (
	"fmt"
)

// culc all operation that has more priorirt than token
func culcHighPriority(numbers []interface{},
	opers []Operation,
	token Operation) (
	[]interface{},
	[]Operation,
	error) {
	for len(opers) != 0 && Priority(token, opers[len(opers)-1]) {
		if len(numbers) > 1 {
			oper := opers[len(opers)-1].getFunc()
			opers = opers[:len(opers)-1]
			newNumber, err := oper(numbers[len(numbers)-2],
				numbers[len(numbers)-1])
			if err != nil {
				err = fmt.Errorf("problem with operations: %s", err)
				return nil, nil, err
			}
			numbers = numbers[:len(numbers)-2]
			numbers = append(numbers, newNumber)
		} else {
			err := fmt.Errorf("don't have numbers for operation: %s",
				opers[len(opers)-1].text)
			return nil, nil, err
		}
	}
	return numbers, opers, nil
}

// Culc culculate numbers and matrix
func Culc(tokens []interface{}) (interface{}, error) {
	var numbers []interface{}
	var opers []Operation
	var err error
	for i := 0; i < len(tokens); i++ {
		switch token := tokens[i].(type) {
		case Complex:
			numbers = append(numbers, token)
		case Operation:
			numbers, opers, err = culcHighPriority(numbers,
				opers, token)
			if err != nil {
				err := fmt.Errorf("problem with culcHighPriopity: %s", err)
				return nil, err
			}
			opers = append(opers, token)
		}
	}
	for k := len(opers) - 1; k >= 0; k-- {
		if len(numbers) > 1 {
			oper := opers[k].getFunc()
			opers = opers[:k]
			newNumber, err := oper(numbers[len(numbers)-2],
				numbers[len(numbers)-1])
			if err != nil {
				err = fmt.Errorf("problem with operation: %s", err)
				return nil, err
			}
			numbers = numbers[:len(numbers)-2]
			numbers = append(numbers, newNumber)
		} else {
			err = fmt.Errorf("Нет чисел для операций")
			return nil, err
		}
	}
	if len(numbers) == 1 {
		return numbers[0], nil
	}
	err = fmt.Errorf("Осталось слишком много чисел")
	return nil, err
}
