package main

import "fmt"

type Braket struct {
	symbol byte
}

func GetBraket(str string) (interface{}, int , error){
	if isBracket(str[0]) {
		return Braket{str[0]}, 0, nil
	}
	return nil, 0, fmt.Errorf("can't make braket from %c", str[0])
}

func (br Braket) isOpen() bool {
	if br.symbol == '(' {
		return true
	}
	return false
}

func (br Braket) isClose() bool {
	if br.symbol == ')' {
		return true
	}
	return false
}